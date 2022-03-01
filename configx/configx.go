package configx

import (
	"fmt"
)

// Package load configurations from several sources.

type Configx struct {
	loaders []Loader
}

func New(loaders ...Loader) Configx {
	loaders = append(loaders, NewLoader(NewDefaultLoader(), false))
	return Configx{loaders: loaders}
}

func (cfgx Configx) Load(v interface{}) error {
	for _, loader := range cfgx.loaders {
		err := loader.unmarshaller.Unmarshal(v)
		if err != nil {
			if loader.breakOnError {
				return fmt.Errorf("failed to load configuration: %w", err)
			}
		}
	}
	return nil
}

type Loader struct {
	unmarshaller Unmarshaller
	breakOnError bool
}

func NewLoader(
	unmarshaller Unmarshaller,
	breakOnError bool,
) Loader {
	return Loader{
		unmarshaller: unmarshaller,
		breakOnError: breakOnError,
	}
}

type Unmarshaller interface {
	Unmarshal(v interface{}) error
}
