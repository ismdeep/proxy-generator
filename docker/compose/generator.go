package compose

type Service struct {
	Image       string            `yaml:"image"`
	Command     []string          `yaml:"command"`
	Ports       []string          `yaml:"ports"`
	Environment map[string]string `yaml:"environment"`
	Volumes     []string          `yaml:"volumes"`
	Restart     string            `yaml:"restart"`
}

type DockerComposeYAML struct {
	Version  string             `yaml:"version"`
	Services map[string]Service `yaml:"services"`
}
