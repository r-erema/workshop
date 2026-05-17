package tools

import "github.com/mark3labs/mcp-go/mcp"

// Constants for repeated strings.
const (
	keyNamespace   = "namespace"
	keyName        = "name"
	keyContainer   = "container"
	keyPod         = "pod"
	keyStatus      = "status"
	keyTailLines   = "tailLines"
	keyLogLength   = "logLength"
	keyLogs        = "logs"
	keyDescription = "description"
	keyType        = "type"
	keyPattern     = "pattern"

	maxTailLines     = 10000
	defaultTailLines = 100
	maxSinceSeconds  = 86400 // 24 hours max

	namespacePattern = `^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`
	stringType       = "string"
	integerType      = "integer"
)

func GetToolDefinitions() []mcp.Tool {
	return []mcp.Tool{
		{
			Name:        "k8s_get_pod_logs",
			Description: "Retrieve logs from a Kubernetes pod with filtering options",
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]any{
					keyNamespace: map[string]any{
						keyType:        stringType,
						keyDescription: "Kubernetes namespace containing the pod",
						keyPattern:     namespacePattern,
					},
					keyName: map[string]any{
						keyType:        stringType,
						keyDescription: "Name of the pod to get logs from",
						keyPattern:     namespacePattern,
					},
					keyContainer: map[string]any{
						keyType:        stringType,
						keyDescription: "Container name (optional, defaults to first container)",
						keyPattern:     namespacePattern,
					},
					"tailLines": map[string]any{
						keyType:        integerType,
						keyDescription: "Number of lines to tail (optional, defaults to 100)",
						"minimum":      1,
						"maximum":      maxTailLines,
						"default":      defaultTailLines,
					},
					"sinceSeconds": map[string]any{
						keyType:        integerType,
						keyDescription: "Show logs from this many seconds ago (optional)",
						"minimum":      1,
						"maximum":      maxSinceSeconds,
					},
				},
				Required: []string{keyNamespace, keyName},
			},
		},
		{
			Name:        "k8s_list_pods",
			Description: "List all pods in a Kubernetes namespace with their status and details",
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]any{
					keyNamespace: map[string]any{
						keyType:        stringType,
						keyDescription: "Kubernetes namespace to list pods from",
						keyPattern:     namespacePattern,
					},
				},
				Required: []string{keyNamespace},
			},
		},
	}
}
