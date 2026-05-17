package k8s

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

const (
	logBufSize = 2048
)

type Client struct {
	clientset kubernetes.Interface
}

func NewClient(clientset kubernetes.Interface) *Client {
	return &Client{clientset: clientset}
}

func (c *Client) HealthCheck() error {
	_, err := c.clientset.Discovery().ServerVersion()
	if err != nil {
		return fmt.Errorf("k8s cluster is not reachable: %w", err)
	}

	return nil
}

func (c *Client) ListPods(ctx context.Context, namespace string) ([]PodInfo, error) {
	pods, err := c.clientset.CoreV1().Pods(namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("list pods err: %w", err)
	}

	podInfos := make([]PodInfo, len(pods.Items))
	for i := range pods.Items {
		podInfos[i] = PodInfo{
			Name:      pods.Items[i].GetName(),
			Namespace: pods.Items[i].GetNamespace(),
			Status:    string(pods.Items[i].Status.Phase),
			Phase:     string(pods.Items[i].Status.Phase),
			Node:      pods.Items[i].Spec.NodeName,
			Labels:    pods.Items[i].GetLabels(),
			CreatedAt: pods.Items[i].GetCreationTimestamp().Time,
			Restarts:  getTotalRestarts(&pods.Items[i]),
		}
	}

	return podInfos, nil
}

func (c *Client) GetPod(ctx context.Context, namespace, name string) ([]byte, error) {
	pod, err := c.clientset.CoreV1().Pods(namespace).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to get pod %s/%s: %w", namespace, name, err)
	}

	type PodDetail struct {
		*PodInfo

		Containers []ContainerInfo `json:"containers"`
		Events     []string        `json:"events"`
		Conditions []string        `json:"conditions"`
	}

	podDetail := PodDetail{
		PodInfo: &PodInfo{
			Name:      pod.Name,
			Namespace: pod.Namespace,
			Status:    string(pod.Status.Phase),
			Phase:     string(pod.Status.Phase),
			Node:      pod.Spec.NodeName,
			Labels:    pod.Labels,
			CreatedAt: pod.CreationTimestamp.Time,
			Restarts:  getTotalRestarts(pod),
		},
		Containers: getContainerInfo(pod),
		Conditions: getPodConditions(pod),
	}

	data, err := json.MarshalIndent(podDetail, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal pod details: %w", err)
	}

	return data, nil
}

func (c *Client) GetPodContainers(ctx context.Context, namespace, name string) ([]string, error) {
	pod, err := c.clientset.CoreV1().Pods(namespace).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to get pod %s/%s: %w", namespace, name, err)
	}

	containers := make([]string, 0, len(pod.Spec.Containers))
	for _, container := range pod.Spec.Containers {
		containers = append(containers, container.Name)
	}

	return containers, nil
}

func (c *Client) GetPodLogs(
	ctx context.Context,
	namespace, podName, containerName string,
	tailLines *int64,
	sinceSeconds *int64,
) (string, error) {
	logOptions := &v1.PodLogOptions{}

	if containerName != "" {
		logOptions.Container = containerName
	}

	if tailLines != nil {
		logOptions.TailLines = tailLines
	}

	if sinceSeconds != nil {
		logOptions.SinceSeconds = sinceSeconds
	}

	// Get log stream
	req := c.clientset.CoreV1().Pods(namespace).GetLogs(podName, logOptions)

	podLogs, err := req.Stream(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to get logs for pod %s/%s: %w", namespace, podName, err)
	}

	defer func() {
		err = podLogs.Close()
		if err != nil {
			log.Printf("failed to close logs for pod %s/%s: %v", namespace, podName, err)
		}
	}()

	// Read all log content
	buf := make([]byte, logBufSize)

	var logs []byte

	for {
		numBytes, err := podLogs.Read(buf)
		if numBytes == 0 {
			break
		}

		if err != nil {
			break
		}

		logs = append(logs, buf[:numBytes]...)
	}

	return string(logs), nil
}

func getContainerInfo(pod *v1.Pod) []ContainerInfo {
	containers := make([]ContainerInfo, 0, len(pod.Spec.Containers))

	for i, container := range pod.Spec.Containers {
		info := ContainerInfo{
			Name:  container.Name,
			Image: container.Image,
		}

		if i < len(pod.Status.ContainerStatuses) {
			status := pod.Status.ContainerStatuses[i]
			info.Ready = status.Ready
			info.Restarts = status.RestartCount
			info.State = getContainerState(status)
		}

		containers = append(containers, info)
	}

	return containers
}

func getContainerState(status v1.ContainerStatus) string {
	switch {
	case status.State.Running != nil:
		return "Running"
	case status.State.Waiting != nil:
		return "Waiting: " + status.State.Waiting.Reason
	case status.State.Terminated != nil:
		return "Terminated: " + status.State.Terminated.Reason
	default:
		return "Unknown"
	}
}

func getPodConditions(pod *v1.Pod) []string {
	var conditions []string

	for _, condition := range pod.Status.Conditions {
		if condition.Status == v1.ConditionTrue {
			conditions = append(conditions, string(condition.Type))
		}
	}

	return conditions
}

func getTotalRestarts(pod *v1.Pod) int32 {
	var total int32
	for _, status := range pod.Status.ContainerStatuses {
		total += status.RestartCount
	}

	return total
}
