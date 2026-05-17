package config

type Config struct {
	Server ServerConfig `yaml:"server"`
	K8s    K8sConfig    `yaml:"k8s"`
}

type ServerConfig struct {
	Name        string `yaml:"name"`
	Version     string `yaml:"version"`
	Description string `yaml:"description"`
}

type K8sConfig struct {
	ConfigPath string   `yaml:"configPath"`
	Context    string   `yaml:"context"`
	Namespaces []string `yaml:"namespaces"`
}

func Load() (*Config, error) {
	cfg := &Config{
		Server: ServerConfig{
			Name:        "k8s-mcp-server",
			Version:     "1.0.0",
			Description: "Kubernetes MCP Server for AI-powered cluster management",
		},
	}

	return cfg, nil
}
