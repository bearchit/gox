package configx_test

import (
	"testing"

	"github.com/bearchit/gox/configx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDefaultLoader(t *testing.T) {
	t.Parallel()

	var config struct {
		Simple string `default:"simple"`
		Nested struct {
			Name string
		}
		CamelCase string
		SnakeCase string
	}

	cfgx := configx.New()
	err := cfgx.Load(&config)
	require.NoError(t, err)
	assert.Equal(t, "simple", config.Simple)
}
