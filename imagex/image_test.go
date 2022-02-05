package imagex_test

import (
	"bearchit/gox/imagex"
	"bearchit/gox/imagex/testdata"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestImage_IsZero(t *testing.T) {
	t.Run("zero", func(t *testing.T) {
		assert.True(t, imagex.Image{}.IsZero())
	})

	t.Run("non zero", func(t *testing.T) {
		content, err := testdata.Images.ReadFile("images/g670fa3045388cff5904056109444a6ac991aaa695812ab70491c4be9022537d27c525c89dcfa434a88b20fe9e91e33683a0a10193e5a367a493742fee8e55cf5998eaf12257f2af44c9d70447c6f3058_640.jpg")
		require.NoError(t, err)

		image := imagex.Image{
			Content: content,
		}
		assert.False(t, image.IsZero())
	})
}

func TestImage_Ratio(t *testing.T) {
	tests := []struct {
		width  int
		height int
		want   float64
	}{
		{640, 480, 1.3333333333333333},
		{200, 100, 2},
		{100, 200, 0.5},
	}

	for _, tt := range tests {
		tt := tt

		t.Run("", func(t *testing.T) {
			t.Parallel()

			image := imagex.Image{Metadata: imagex.Metadata{Width: tt.width, Height: tt.height}}
			assert.Equal(t, tt.want, image.Ratio())
		})
	}
}

func TestImage_Checksum(t *testing.T) {
	content, err := testdata.Images.ReadFile("images/g670fa3045388cff5904056109444a6ac991aaa695812ab70491c4be9022537d27c525c89dcfa434a88b20fe9e91e33683a0a10193e5a367a493742fee8e55cf5998eaf12257f2af44c9d70447c6f3058_640.jpg")
	require.NoError(t, err)

	image := imagex.Image{
		Content: content,
	}

	assert.Equal(t, "dc00817e1855114e26bbdcc96d3c9154", image.Checksum())
}
