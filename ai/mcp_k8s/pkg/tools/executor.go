package tools

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/r-erema/workshop/ai/mcp_k8s/pkg/k8s"
)

var ErrUnknownTool = errors.New("unknown tool")

type ExecuteResult struct {
	Success   bool           `json:"success"`
	Message   string         `json:"message"`
	Data      map[string]any `json:"data,omitempty"`
	Error     string         `json:"error,omitempty"`
	Timestamp time.Time      `json:"timestamp"`
}

type ToolExecutor struct {
	k8sClient *k8s.Client
}

func NewToolExecutor(k8sClient *k8s.Client) *ToolExecutor {
	return &ToolExecutor{
		k8sClient: k8sClient,
	}
}

func (e *ToolExecutor) ExecuteTool(ctx context.Context, toolName string, inputs map[string]any) *ExecuteResult {
	start := time.Now()

	log.Printf("tool_call(tool name: %s, inputs: %v)", toolName, inputs)

	var result *ExecuteResult

	switch toolName {
	case "k8s_get_pod_logs":
		result = e.executeGetPodLogs(ctx, inputs)
	case "k8s_list_pods":
		result = e.executeListPods(ctx, inputs)
	default:
		result = &ExecuteResult{
			Success:   false,
			Message:   "Unknown tool",
			Error:     fmt.Sprintf("Tool '%s' is not supported", toolName),
			Timestamp: start,
		}
		log.Printf(
			"tool_call(tool name: %s, inputs: %s)",
			time.Since(start),
			fmt.Errorf("%w: %s", ErrUnknownTool, toolName),
		)
	}

	return result
}

func (e *ToolExecutor) executeGetPodLogs(ctx context.Context, inputs map[string]any) *ExecuteResult {
	namespace, ok := inputs[keyNamespace].(string)
	if !ok {
		return errorResult("Invalid namespace input", "namespace must be a string")
	}

	name, ok := inputs[keyName].(string)
	if !ok {
		return errorResult("Invalid name input", "name must be a string")
	}

	containerName := extractContainerName(inputs)
	lines := extractTailLines(inputs)
	sinceSeconds := extractSinceSeconds(inputs)

	containerName = e.resolveContainerName(ctx, namespace, name, containerName)
	if containerName == "" {
		return errorResult("Failed to resolve container name", "could not determine container name")
	}

	logs, err := e.k8sClient.GetPodLogs(ctx, namespace, name, containerName, &lines, sinceSeconds)
	if err != nil {
		return &ExecuteResult{
			Success:   false,
			Message:   "Failed to retrieve pod logs",
			Error:     err.Error(),
			Timestamp: time.Now(),
		}
	}

	return &ExecuteResult{
		Success: true,
		Message: fmt.Sprintf(
			"Successfully retrieved logs from pod %s/%s (container: %s)",
			namespace,
			name,
			containerName,
		),
		Data: map[string]any{
			keyNamespace: namespace,
			keyPod:       name,
			keyContainer: containerName,
			keyTailLines: lines,
			keyLogs:      logs,
			keyLogLength: len(logs),
		},
		Timestamp: time.Now(),
	}
}

func extractContainerName(inputs map[string]any) string {
	if container, exists := inputs[keyContainer]; exists {
		if containerStr, ok := container.(string); ok {
			return containerStr
		}
	}

	return ""
}

func extractTailLines(inputs map[string]any) int64 {
	lines := int64(defaultTailLines)

	if tl, exists := inputs["tailLines"]; exists {
		if tlFloat, ok := tl.(float64); ok {
			lines = int64(tlFloat)
		}
	}

	return lines
}

func extractSinceSeconds(inputs map[string]any) *int64 {
	if ss, exists := inputs["sinceSeconds"]; exists {
		if ssFloat, ok := ss.(float64); ok {
			seconds := int64(ssFloat)

			return &seconds
		}
	}

	return nil
}

func (e *ToolExecutor) resolveContainerName(ctx context.Context, namespace, name, containerName string) string {
	if containerName != "" {
		return containerName
	}

	containers, err := e.k8sClient.GetPodContainers(ctx, namespace, name)
	if err != nil {
		return ""
	}

	if len(containers) == 0 {
		return ""
	}

	return containers[0]
}

func (e *ToolExecutor) executeListPods(ctx context.Context, inputs map[string]any) *ExecuteResult {
	namespace, ok := inputs[keyNamespace].(string)
	if !ok {
		return errorResult("Invalid namespace input", "namespace must be a string")
	}

	pods, err := e.k8sClient.ListPods(ctx, namespace)
	if err != nil {
		return &ExecuteResult{
			Success:   false,
			Message:   "Failed to list pods",
			Error:     err.Error(),
			Timestamp: time.Now(),
		}
	}

	podList := make([]map[string]any, len(pods))
	for i, pod := range pods {
		podList[i] = map[string]any{
			keyName:      pod.Name,
			keyNamespace: pod.Namespace,
			keyStatus:    pod.Status,
			"phase":      pod.Phase,
			"node":       pod.Node,
			"labels":     pod.Labels,
			"createdAt":  pod.CreatedAt.Format(time.RFC3339),
			"restarts":   pod.Restarts,
		}
	}

	return &ExecuteResult{
		Success: true,
		Message: fmt.Sprintf("Successfully listed %d pods in namespace %s", len(pods), namespace),
		Data: map[string]any{
			keyNamespace: namespace,
			"podCount":   len(pods),
			"pods":       podList,
		},
		Timestamp: time.Now(),
	}
}

func errorResult(message, errorMsg string) *ExecuteResult {
	return &ExecuteResult{
		Success:   false,
		Message:   message,
		Error:     errorMsg,
		Timestamp: time.Now(),
	}
}
