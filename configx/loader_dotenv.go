package configx

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/mitchellh/mapstructure"
)

type DotEnvLoader struct {
	paths []string
}

func NewDotEnvLoader(paths ...string) DotEnvLoader {
	loader := DotEnvLoader{paths: paths}
	if len(paths) == 0 {
		loader.paths = []string{
			".env",
			".env.local",
		}
	}
	return loader
}

func (loader DotEnvLoader) Unmarshal(v interface{}) error {
	reads, err := godotenv.Read(loader.paths...)
	if err != nil {
		return fmt.Errorf("failed to read dotenv files: %w", err)
	}
	if err = mapstructure.Decode(reads, v); err != nil {
		return fmt.Errorf("unmarshal dotenv: %w", err)
	}
	return nil
}
