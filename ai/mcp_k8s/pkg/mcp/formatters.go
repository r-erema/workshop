package mcp

import (
	"fmt"
	"strings"
	"time"

	"k8s.io/apimachinery/pkg/util/json"
)

type ResourceFormatter struct{}

func NewResourceFormatter() *ResourceFormatter {
	return &ResourceFormatter{}
}

func (f ResourceFormatter) FormatPodForAI(podData []byte) (string, error) {
	var pod map[string]any

	err := json.Unmarshal(podData, &pod)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal pod data: %w", err)
	}

	summary := &strings.Builder{}
	summary.WriteString("# Pod Summary\n\n")

	_, _ = fmt.Fprintf(summary, "**Name**: %s\n", pod["name"])
	_, _ = fmt.Fprintf(summary, "**Namespace**: %s\n", pod["namespace"])
	_, _ = fmt.Fprintf(summary, "**Status**: %s\n", pod["status"])
	_, _ = fmt.Fprintf(summary, "**Node**: %s\n", pod["node"])

	if restarts, ok := pod["restarts"].(float64); ok && restarts > 0 {
		_, _ = fmt.Fprintf(summary, "**⚠️ Restarts**: %.0f\n", restarts)
	}

	if createdAt, ok := pod["createdAt"].(string); ok {
		t, err := time.Parse(time.RFC3339, createdAt)
		if err == nil {
			age := time.Since(t)

			_, _ = fmt.Fprintf(summary, "**Age**: %s\n", age)
		}
	}

	f.writeContainers(summary, pod)
	f.writeConditions(summary, pod)
	f.writeLabels(summary, pod)

	summary.WriteString("\n---\n")
	summary.WriteString("*Use this information to understand the pod's current state and troubleshoot any issues.*")

	return summary.String(), nil
}

func (f ResourceFormatter) writeContainers(summary *strings.Builder, pod map[string]any) {
	summary.WriteString("\n## Containers\n\n")

	containers, ok := pod["containers"].([]any)
	if !ok {
		return
	}

	for i := range containers {
		container, ok := containers[i].(map[string]any)
		if !ok {
			continue
		}

		name, _ := container["name"].(string)
		image, _ := container["image"].(string)
		ready, _ := container["ready"].(bool)
		state, _ := container["state"].(string)

		status := "🟢 Ready"
		if !ready {
			status = "🔴 Not Ready"
		}

		_, _ = fmt.Fprintf(summary, "- **%s**: %s\n", name, status)
		_, _ = fmt.Fprintf(summary, "- Image: `%s`\n", image)
		_, _ = fmt.Fprintf(summary, "- State: %s\n", state)

		if restarts, ok := container["restarts"].(float64); ok && restarts > 0 {
			_, _ = fmt.Fprintf(summary, "- Restarts**: %.0f\n", restarts)
		}
	}
}

func (f ResourceFormatter) writeConditions(summary *strings.Builder, pod map[string]any) {
	conditions, ok := pod["conditions"].([]any)
	if !ok || len(conditions) == 0 {
		return
	}

	summary.WriteString("\n## Conditions\n\n")

	for i := range conditions {
		_, _ = fmt.Fprintf(summary, "- %s\n", conditions[i])
	}
}

func (f ResourceFormatter) writeLabels(summary *strings.Builder, pod map[string]any) {
	labels, ok := pod["labels"].(map[string]any)
	if !ok || len(labels) == 0 {
		return
	}

	summary.WriteString("\n## Labels\n\n")

	for k, v := range labels {
		_, _ = fmt.Fprintf(summary, "- `%s`: `%s`\n", k, v)
	}
}
