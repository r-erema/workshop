package mcp

import "encoding/json"

type Resource struct {
	URI         string            `json:"uri"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	MIMEType    string            `json:"mimeType"`
	Metadata    map[string]string `json:"metadata,omitempty"`
}

type ResourceContent struct {
	URI      string          `json:"uri"`
	MIMEType string          `json:"mimeType"`
	Text     string          `json:"text"`
	Blob     string          `json:"blob"`
	Metadata json.RawMessage `json:"metadata,omitempty"`
}

type K8sResourceType string

const (
	ResourceTypePod K8sResourceType = "pod"
)

type ResourceIdentifier struct {
	Type      K8sResourceType `json:"type"`
	Namespace string          `json:"namespace"`
	Name      string          `json:"name"`
}

func (r ResourceIdentifier) ToURI() string {
	if r.Namespace != "" {
		return "k8s://" + string(r.Type) + "/" + r.Name
	}

	return "k8s://" + string(r.Type) + "/" + r.Namespace + "/" + r.Name
}
