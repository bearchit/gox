package available

import (
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

type (
	Extension struct {
		entc.DefaultExtension
		templates []*gen.Template
	}

	ExtensionOption func(*Extension) error
)

func NewExtension(opts ...ExtensionOption) (*Extension, error) {
	ex := &Extension{templates: AllTemplates}
	for _, opt := range opts {
		if err := opt(ex); err != nil {
			return nil, err
		}
	}
	return ex, nil
}

func (e *Extension) Templates() []*gen.Template {
	return e.templates
}
