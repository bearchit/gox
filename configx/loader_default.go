package configx

import (
	"fmt"

	"github.com/creasty/defaults"
)

type DefaultLoader struct{}

func NewDefaultLoader() *DefaultLoader {
	return &DefaultLoader{}
}

func (loader DefaultLoader) Unmarshal(v interface{}) error {
	if err := defaults.Set(v); err != nil {
		return fmt.Errorf("unmarshal default value: %w", err)
	}
	return nil
}
