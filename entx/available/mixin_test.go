package available_test

import (
	"entgo.io/ent"
	"github.com/bearchit/gox/entx/available"
	"github.com/stretchr/testify/assert"
	"testing"
)

func hasFields(fields []ent.Field, fieldNames ...string) bool {
	matches := make([]bool, 0)
	for _, f := range fields {
		for _, fn := range fieldNames {
			if f.Descriptor().Name == fn {
				matches = append(matches, true)
			}
		}
	}
	return len(matches) == len(fieldNames)
}

func TestWithActivationOption(t *testing.T) {
	mixin := available.NewMixin(
		available.WithActivationOption(func(option *available.ActivationOption) {
			option.SetFieldName("state")
		}),
	)
	assert.True(t, hasFields(mixin.Fields(), "state"))
}
