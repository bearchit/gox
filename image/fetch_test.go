package image_test

import (
	"bearchit/gox/image"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	_ "golang.org/x/image/webp"
)

func TestFetchFromURLBulk(t *testing.T) {
	tests := []struct {
		urls      []string
		checksums []string
	}{
		{
			[]string{
				"https://i.pinimg.com/564x/b2/6b/76/b26b762548cdbba397f9b0079264fb4b.jpg",
				"https://i.pinimg.com/564x/e3/c7/ac/e3c7ac97d1a893e5e6a3bd48ce3fcd02.jpg",
				"https://i.pinimg.com/564x/b7/ff/5c/b7ff5c2511373ebd7d0c1fff02978c7b.jpg",
			},
			[]string{
				"a6ca16a40cfcf28ebfb420821f2d0597",
				"8ecd6bf1683327b69d9007b07113ec0e",
				"4998b7015fded86aef6ad12c9e1ab360",
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run("", func(t *testing.T) {
			t.Parallel()

			images, err := image.FetchFromURLBulk(context.Background(), tt.urls)
			require.NoError(t, err)

			for i, x := range images {
				assert.Equal(t, tt.checksums[i], x.Checksum())
			}
		})
	}
}
