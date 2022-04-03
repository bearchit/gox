package configx_test

import (
	"testing"

	"github.com/bearchit/gox/configx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEnvLoader(t *testing.T) {
	var config struct {
		Simple string
		Nested struct {
			Name string
		}
		CamelCase string
		SnakeCase string
	}

	cfgx := configx.New(
		configx.NewLoader(
			configx.NewEnvLoader(""),
			true,
		),
	)
	t.Setenv("SIMPLE", "simple")
	err := cfgx.Load(&config)
	require.NoError(t, err)
	assert.Equal(t, "simple", config.Simple)
}
