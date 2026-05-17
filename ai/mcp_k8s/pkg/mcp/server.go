package mcp

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/mark3labs/mcp-go/mcp"
	mcpserver "github.com/mark3labs/mcp-go/server"
	"github.com/r-erema/workshop/ai/mcp_k8s/internal/config"
	"github.com/r-erema/workshop/ai/mcp_k8s/pkg/k8s"
	"github.com/r-erema/workshop/ai/mcp_k8s/pkg/tools"
)

var (
	ErrInvalidURIF           = errors.New("invalid URI format. Expected k8s://<resource-type>/<namespace>/<name>")
	ErrInvalidURIPartsCountF = errors.New("invalid URI format: wrong number of parts")
	ErrInvalidResourceType   = errors.New("unsupported resource type")
	ErrToolExecutionFailed   = errors.New("tool execution failed")
)

const (
	uriPartsCount         = 3
	maxLogTruncateLength  = 5000
	logTruncateSafeMargin = 20
	resourceLimit         = 5
)

type Server struct {
	config       *config.Config
	k8sClient    *k8s.Client
	mcpServer    *mcpserver.MCPServer
	toolExecutor *tools.ToolExecutor
	formatter    *ResourceFormatter
}

func InitServer(config *config.Config, k8sClient *k8s.Client) *Server {
	server := &Server{
		config:    config,
		k8sClient: k8sClient,
		mcpServer: mcpserver.NewMCPServer(
			"k8s-mcp-server",
			"1.0.0",
			mcpserver.WithResourceCapabilities(true, true),
			mcpserver.WithToolCapabilities(true),
		),
		toolExecutor: tools.NewToolExecutor(k8sClient),
		formatter:    NewResourceFormatter(),
	}

	server.registerResources()
	server.registerTools()

	return server
}

// GetMCPServer returns the underlying MCP server for testing purposes.
func (s *Server) GetMCPServer() *mcpserver.MCPServer {
	return s.mcpServer
}

func (s *Server) registerResources() {
	ctx := context.Background()

	pods, err := s.k8sClient.ListPods(ctx, "")
	if err != nil {
		log.Printf("list pods err: %v", err)

		return
	}

	limit := min(len(pods), resourceLimit)

	for i := range pods[:limit] {
		resource := mcp.Resource{
			URI:         fmt.Sprintf("k8s://pod/%s/%s", pods[i].Namespace, pods[i].Name),
			Name:        fmt.Sprintf("Pod: %s/%s", pods[i].Namespace, pods[i].Name),
			Description: fmt.Sprintf("Kubernetes Pod in namespace %s (Status: %s)", pods[i].Namespace, pods[i].Status),
			MIMEType:    "application/json",
		}
		s.mcpServer.AddResource(resource, s.handleResourceRead)
	}
}

func (s *Server) handleResourceRead(
	ctx context.Context,
	request mcp.ReadResourceRequest,
) ([]mcp.ResourceContents, error) {
	uri := request.Params.URI
	log.Printf("Handling read_resource request for URI: %s", uri)

	if !strings.HasPrefix(uri, "k8s://") {
		return nil, fmt.Errorf("%w: got %s", ErrInvalidURIF, uri)
	}

	parts := strings.Split(strings.TrimPrefix(uri, "k8s://"), "/")
	if len(parts) != uriPartsCount {
		return nil, fmt.Errorf("%w: got %d parts", ErrInvalidURIPartsCountF, len(parts))
	}

	resourceType, namespace, name := parts[0], parts[1], parts[2]

	if resourceType != string(ResourceTypePod) {
		return nil, fmt.Errorf("%w: %s", ErrInvalidResourceType, resourceType)
	}

	content, err := s.k8sClient.GetPod(ctx, namespace, name)
	if err != nil {
		return nil, fmt.Errorf("failed to get resource %s: %w", uri, err)
	}

	formattedContent, err := s.formatter.FormatPodForAI(content)
	if err != nil {
		return nil, fmt.Errorf("failed to format pod data: %w", err)
	}

	return []mcp.ResourceContents{
		&mcp.TextResourceContents{
			URI:      uri,
			MIMEType: "application/json",
			Text:     formattedContent,
		},
	}, nil
}

func (s *Server) registerTools() {
	toolDefinitions := tools.GetToolDefinitions()

	for _, toolDef := range toolDefinitions {
		s.mcpServer.AddTool(toolDef, s.handleToolCall)
		log.Printf("Registered tool: %s", toolDef.Name)
	}

	log.Printf("Registered %d tools", len(toolDefinitions))
}

func (s *Server) handleToolCall(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	toolName := request.Params.Name
	arguments := request.Params.Arguments

	log.Printf("Handling tool call: %s with arguments: %v", toolName, arguments)

	args, ok := arguments.(map[string]any)
	if !ok {
		return nil, fmt.Errorf("%w: arguments must be a map[string]any", ErrToolExecutionFailed)
	}

	result := s.toolExecutor.ExecuteTool(ctx, toolName, args)

	if result.Success {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{
					Type: "text",
					Text: formatToolResult(result),
				},
			},
		}, nil
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{
				Type: "text",
				Text: formatToolError(result),
			},
		},
	}, fmt.Errorf("%w: %s", ErrToolExecutionFailed, result.Error)
}

func formatToolResult(result *tools.ExecuteResult) string {
	var output strings.Builder

	_, _ = fmt.Fprintf(&output, "# ✅ %s\n\n", result.Message)
	_, _ = fmt.Fprintf(&output, "**Executed at**: %s\n\n", result.Timestamp.Format(time.RFC3339))

	if len(result.Data) > 0 {
		formatResultData(&output, result.Data)
	}

	output.WriteString("\n---\n*Operation completed successfully*")

	return output.String()
}

func formatResultData(output *strings.Builder, data map[string]any) {
	output.WriteString("## Result Details\n\n")

	for key, value := range data {
		switch typedValue := value.(type) {
		case string:
			formatStringValue(output, key, typedValue)
		case int, int32, int64, float64:
			_, _ = fmt.Fprintf(output, "- **%s**: %v\n", key, typedValue)
		case time.Time:
			_, _ = fmt.Fprintf(output, "- **%s**: %s\n", key, typedValue.Format(time.RFC3339))
		case map[string]any:
			_, _ = fmt.Fprintf(output, "- **%s**: %v\n", key, typedValue)
		default:
			_, _ = fmt.Fprintf(output, "- **%s**: %v\n", key, typedValue)
		}
	}
}

func formatStringValue(output *strings.Builder, key string, value string) {
	if key == "logs" {
		if len(value) > maxLogTruncateLength {
			truncated := value[:maxLogTruncateLength-logTruncateSafeMargin]
			_, _ = fmt.Fprintf(output, "**%s**: (truncated to %d chars)\n```\n%s\n...\n```\n\n",
				key, maxLogTruncateLength, truncated)
		} else {
			_, _ = fmt.Fprintf(output, "**%s**:\n```\n%s\n```\n\n", key, value)
		}
	} else {
		_, _ = fmt.Fprintf(output, "- **%s**: %s\n", key, value)
	}
}

func formatToolError(result *tools.ExecuteResult) string {
	output := fmt.Sprintf("# ❌ %s\n\n", result.Message)
	output += fmt.Sprintf("**Error**: %s\n\n", result.Error)
	output += fmt.Sprintf("**Timestamp**: %s\n\n", result.Timestamp.Format(time.RFC3339))

	output += "## Troubleshooting\n\n"
	output += "- Check that the resource exists and you have permission to access it\n"
	output += "- Verify that the namespace and resource names are correct\n"
	output += "- Ensure the Kubernetes cluster is accessible\n"
	output += "- Review the error message above for specific details\n\n"

	output += "---\n*Operation failed - review the error details above*"

	return output
}
