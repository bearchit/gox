package configx

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type DotEnvLoader struct {
	prefix string
	paths  []string
}

func NewDotEnvLoader(prefix string, paths ...string) DotEnvLoader {
	loader := DotEnvLoader{prefix: prefix, paths: paths}
	if len(paths) == 0 {
		loader.paths = []string{
			".env",
			".env.local",
		}
	}
	return loader
}

func (loader DotEnvLoader) Unmarshal(v interface{}) error {
	if err := godotenv.Load(loader.paths...); err != nil {
		return fmt.Errorf("failed to load dotenv files: %w", err)
	}
	if err := envconfig.Process(loader.prefix, v); err != nil {
		return fmt.Errorf("unmarshal dotenv: %w", err)
	}
	return nil
}
