package configx_test

import (
	"github.com/bearchit/gox/configx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestYAMLLoader(t *testing.T) {
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
			configx.NewYAMLLoader("testdata/test.yml"),
			true,
		),
	)
	err := cfgx.Load(&config)
	require.NoError(t, err)
	assert.Equal(t, "simple", config.Simple)
}
