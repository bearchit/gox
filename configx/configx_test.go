package configx_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mitchellh/mapstructure"

	"github.com/bearchit/gox/configx"
)

func TestConfigx(t *testing.T) {
	t.Parallel()

	t.Run("", func(t *testing.T) {
		t.Parallel()

		type config struct {
			Name string
		}

		cfgx := configx.New(
			configx.NewLoader(mockMarshaller{wantErr: true}, true),
			configx.NewLoader(mockMarshaller{}, true),
		)
		var got config
		err := cfgx.Load(&got)
		assert.Error(t, errors.Unwrap(err), "error")

		// require.NoError(t, err)
		// assert.Equal(t, "simple", got.Name)
	})
}

type mockMarshaller struct {
	wantErr bool
}

func (m mockMarshaller) Unmarshal(v interface{}) error {
	if m.wantErr {
		return fmt.Errorf("error")
	}
	return mapstructure.Decode(map[string]interface{}{
		"Name": "simple",
	}, v)
}
