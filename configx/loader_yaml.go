package configx

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type YAMLLoader struct {
	path string
}

func NewYAMLLoader(path string) YAMLLoader {
	return YAMLLoader{path: path}
}

func (loader YAMLLoader) Unmarshal(v interface{}) error {
	content, err := os.ReadFile(loader.path)
	if err != nil {
		return fmt.Errorf("failed to read YAML file: %w", err)
	}
	if err = yaml.Unmarshal(content, v); err != nil {
		return fmt.Errorf("unmarshal yaml: %w", err)
	}
	return nil
}
