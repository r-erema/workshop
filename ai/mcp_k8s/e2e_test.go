package mcp_k8s_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"testing"

	"github.com/mark3labs/mcp-go/client"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/ollama/ollama/api"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/r-erema/workshop/ai/mcp_k8s/internal/config"
	k8sclient "github.com/r-erema/workshop/ai/mcp_k8s/pkg/k8s"
	mcpk8s "github.com/r-erema/workshop/ai/mcp_k8s/pkg/mcp"
	"github.com/r-erema/workshop/utils/test"
	"github.com/testcontainers/testcontainers-go"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	ollamaModel       = "tinyllama:latest" // Smallest available model (~637MB)
	cachedOllamaImage = "ollama-with-model-tinyllama:latest"

	testPodName   = "test-pod"
	testNamespace = "default"
)

var (
	errNoJSONObject      = errors.New("no JSON object found in response")
	errNoValidJSONObject = errors.New("no valid JSON object found in response")
)

func TestMCPK8SIntegration(t *testing.T) {
	if os.Getenv("GITHUB_ACTIONS") == "true" {
		Skip("Skip in Github Actions")
	}

	t.Parallel()

	RegisterFailHandler(Fail)
	RunSpecs(t, "MCP K8S Integration Suite")
}

// ToolCall represents a tool call from the LLM.
type ToolCall struct {
	Name      string         `json:"name"`
	Arguments map[string]any `json:"arguments"`
}

var _ = Describe("MCP K8S Interaction with LLM", func() {
	var (
		ollamaContainer testcontainers.Container
		mcpServer       *mcpk8s.Server
		mcpClient       *client.Client
	)

	BeforeEach(func() {
		test.SkipInGitHubActions()

		ctx := GinkgoT().Context()

		k8sClientset := test.RunCluster()

		By("Creating a test pod in the default namespace")

		testPod := &corev1.Pod{
			ObjectMeta: metav1.ObjectMeta{
				Name:      testPodName,
				Namespace: testNamespace,
				Labels: map[string]string{
					"app": "test",
				},
			},
			Spec: corev1.PodSpec{
				Containers: []corev1.Container{
					{
						Name:  "nginx",
						Image: "nginx:latest",
					},
				},
			},
		}
		_, err := k8sClientset.CoreV1().Pods(testNamespace).Create(ctx, testPod, metav1.CreateOptions{})
		Expect(err).NotTo(HaveOccurred())

		By("Run Ollama container")

		ollamaContainer = test.StartOllamaContainer(ctx)
		Expect(ollamaContainer).NotTo(BeNil())
		DeferCleanup(func() {
			By("Stopping Ollama container")

			err := ollamaContainer.Terminate(ctx)
			Expect(err).NotTo(HaveOccurred())
		})

		By("Initializing MCP Server")

		cfg := &config.Config{
			Server: config.ServerConfig{
				Name:        "k8s-mcp-server",
				Version:     "1.0.0",
				Description: "Kubernetes MCP Server for testing",
			},
			K8s: config.K8sConfig{
				ConfigPath: "",
			},
		}
		mcpServer = mcpk8s.InitServer(cfg, k8sclient.NewClient(k8sClientset))
		Expect(mcpServer).NotTo(BeNil())

		By("Initializing MCP Client")

		mcpClient, err = client.NewInProcessClient(mcpServer.GetMCPServer())
		Expect(err).NotTo(HaveOccurred())
		DeferCleanup(func() {
			err := mcpClient.Close()
			Expect(err).NotTo(HaveOccurred())
		})

		initRequest := mcp.InitializeRequest{
			Request: mcp.Request{
				Method: "initialize",
			},
			Params: mcp.InitializeParams{
				ProtocolVersion: mcp.LATEST_PROTOCOL_VERSION,
				ClientInfo: mcp.Implementation{
					Name:    "test-client",
					Version: "1.0.0",
				},
			},
		}
		_, err = mcpClient.Initialize(ctx, initRequest)
		Expect(err).NotTo(HaveOccurred())
	})

	It("should use LLM to list pods via MCP tool calling", func() {
		test.SkipInGitHubActions()

		ctx := GinkgoT().Context()

		By("Sending prompt to LLM to trigger k8s_list_pods tool")

		toolsList, err := mcpClient.ListTools(ctx, mcp.ListToolsRequest{})
		Expect(err).NotTo(HaveOccurred())

		toolDescriptions := buildToolDescriptions(toolsList)

		prompt := fmt.Sprintf(`You are a Kubernetes assistant with access to tools.
Available tools:
%s

User request: list all pods in default namespace

You MUST respond ONLY with a JSON object containing the tool call in this exact format:
{"name": "k8s_list_pods", "arguments": {"namespace": "default"}}

You MUST not include any other text, only the JSON object.`, toolDescriptions)

		ollamaHost, err := ollamaContainer.Host(ctx)
		Expect(err).NotTo(HaveOccurred())

		ollamaPort, err := ollamaContainer.MappedPort(ctx, "11434")
		Expect(err).NotTo(HaveOccurred())

		oLLamaEndpoint := (&url.URL{
			Scheme: "http",
			Host:   net.JoinHostPort(ollamaHost, ollamaPort.Port()),
		}).String()

		ollamaUrl, err := url.Parse(oLLamaEndpoint)

		client := api.NewClient(ollamaUrl, http.DefaultClient)

		Expect(err).NotTo(HaveOccurred())

		req := &api.GenerateRequest{
			Model:  ollamaModel,
			Prompt: prompt,
			Stream: new(false),
		}

		var llmResponse string

		err = client.Generate(ctx, req, func(response api.GenerateResponse) error {
			llmResponse = response.Response

			return nil
		})
		Expect(err).NotTo(HaveOccurred())
		Expect(llmResponse).NotTo(BeEmpty())

		By("Parsing LLM response and invoking MCP tool")

		toolCall, err := parseToolCall(llmResponse)
		Expect(err).NotTo(HaveOccurred())
		Expect(toolCall.Name).To(Equal("k8s_list_pods"))

		toolRequest := mcp.CallToolRequest{
			Request: mcp.Request{
				Method: "tools/call",
			},
			Params: mcp.CallToolParams{
				Name:      toolCall.Name,
				Arguments: toolCall.Arguments,
			},
		}

		toolResult, err := mcpClient.CallTool(ctx, toolRequest)
		Expect(err).NotTo(HaveOccurred())
		Expect(toolResult.Content).NotTo(BeEmpty())

		var resultText string

		for _, content := range toolResult.Content {
			if textContent, ok := content.(mcp.TextContent); ok {
				resultText = textContent.Text

				break
			}
		}

		Expect(resultText).NotTo(BeEmpty())
		Expect(resultText).To(ContainSubstring(testPodName))
		Expect(resultText).To(ContainSubstring(testNamespace))

		Expect(resultText).To(ContainSubstring("Successfully listed"))
	})
})

func buildToolDescriptions(toolsList *mcp.ListToolsResult) string {
	var buf bytes.Buffer
	for _, tool := range toolsList.Tools {
		buf.WriteString(fmt.Sprintf("- %s: %s\n", tool.Name, tool.Description))
		buf.WriteString("  Parameters:\n")

		for paramName, paramSchema := range tool.InputSchema.Properties {
			if paramMap, ok := paramSchema.(map[string]any); ok {
				desc := ""
				if d, ok := paramMap["description"].(string); ok {
					desc = d
				}

				buf.WriteString(fmt.Sprintf("    - %s: %s\n", paramName, desc))
			}
		}
	}

	return buf.String()
}

func parseToolCall(response string) (*ToolCall, error) {
	startIdx := strings.Index(response, "{")
	if startIdx == -1 {
		return nil, errNoJSONObject
	}

	endIdx := strings.LastIndex(response, "}")
	if endIdx == -1 || endIdx <= startIdx {
		return nil, errNoValidJSONObject
	}

	jsonStr := response[startIdx : endIdx+1]

	var toolCall ToolCall

	err := json.Unmarshal([]byte(jsonStr), &toolCall)
	if err != nil {
		return nil, fmt.Errorf("parsing tool call JSON: %w", err)
	}

	return &toolCall, nil
}
