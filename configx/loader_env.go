package configx

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type EnvLoader struct {
	prefix string
}

func NewEnvLoader(prefix string) EnvLoader {
	return EnvLoader{prefix: prefix}
}

func (loader EnvLoader) Unmarshal(v interface{}) error {
	if err := envconfig.Process(loader.prefix, v); err != nil {
		return fmt.Errorf("unmarshal env :%w", err)
	}
	return nil
}
