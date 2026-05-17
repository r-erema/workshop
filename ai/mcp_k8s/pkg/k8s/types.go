package k8s

import "time"

type PodInfo struct {
	Name      string            `json:"name"      yaml:"name"`
	Namespace string            `json:"namespace" yaml:"namespace"`
	Status    string            `json:"status"    yaml:"status"`
	Phase     string            `json:"phase"     yaml:"phase"`
	Node      string            `json:"node"      yaml:"node"`
	Labels    map[string]string `json:"labels"    yaml:"labels"`
	CreatedAt time.Time         `json:"createdAt" yaml:"createdAt"`
	Restarts  int32             `json:"restarts"  yaml:"restarts"`
}

type NamespaceInfo struct {
	Name      string            `yaml:"name"`
	Status    string            `yaml:"status"`
	Labels    map[string]string `yaml:"labels"`
	CreatedAt time.Time         `yaml:"createdAt"`
}

type ContainerInfo struct {
	Name     string `json:"name"`
	Image    string `json:"image"`
	Ready    bool   `json:"ready"`
	Restarts int32  `json:"restarts"`
	State    string `json:"state"`
}
