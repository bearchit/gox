package imagex_test

import (
	"context"
	"github.com/bearchit/gox/imagex"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	_ "golang.org/x/image/webp"
)

var checksums = map[string]string{
	"gae8b9ffe1cceaa155812f29ac70059fbabb2c23eeac7f3fdbc8bdfe7b0d92954bd85d810c2e9dd3c3c3f4f3c3265b6eae10afc2401594f543f0e68f44c183f6acc06f15b824c25ae33c572a16f87c4a7_640.jpg": "e5fc6af4a9f71be87badd703c6304795",
	"g670fa3045388cff5904056109444a6ac991aaa695812ab70491c4be9022537d27c525c89dcfa434a88b20fe9e91e33683a0a10193e5a367a493742fee8e55cf5998eaf12257f2af44c9d70447c6f3058_640.jpg": "dc00817e1855114e26bbdcc96d3c9154",
	"gc3644bb7563bc67658940d8712d89d5107f3f289f7ee555694313500d4301c2cc85426e3084d93468e7c47f31ead70ea3933e9e4abdf4f695dceba35dc3a1963466312e663a1d8bfe29a1426388ce043_640.jpg": "8738c73b3436bc8c6df2212c0451339a",
}

func TestFetchFromURLBulk(t *testing.T) {
	tests := []struct {
		urls    []string
		wantErr bool
	}{
		{
			[]string{
				"https://pixabay.com/get/gae8b9ffe1cceaa155812f29ac70059fbabb2c23eeac7f3fdbc8bdfe7b0d92954bd85d810c2e9dd3c3c3f4f3c3265b6eae10afc2401594f543f0e68f44c183f6acc06f15b824c25ae33c572a16f87c4a7_640.jpg",
				"https://pixabay.com/get/g670fa3045388cff5904056109444a6ac991aaa695812ab70491c4be9022537d27c525c89dcfa434a88b20fe9e91e33683a0a10193e5a367a493742fee8e55cf5998eaf12257f2af44c9d70447c6f3058_640.jpg",
				"https://pixabay.com/get/gc3644bb7563bc67658940d8712d89d5107f3f289f7ee555694313500d4301c2cc85426e3084d93468e7c47f31ead70ea3933e9e4abdf4f695dceba35dc3a1963466312e663a1d8bfe29a1426388ce043_640.jpg",
			},
			false,
		},
		{
			[]string{
				"https://",
				"https://pixabay.com/get/g670fa3045388cff5904056109444a6ac991aaa695812ab70491c4be9022537d27c525c89dcfa434a88b20fe9e91e33683a0a10193e5a367a493742fee8e55cf5998eaf12257f2af44c9d70447c6f3058_640.jpg",
				"https://pixabay.com/get/gc3644bb7563bc67658940d8712d89d5107f3f289f7ee555694313500d4301c2cc85426e3084d93468e7c47f31ead70ea3933e9e4abdf4f695dceba35dc3a1963466312e663a1d8bfe29a1426388ce043_640.jpg",
			},
			true,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run("", func(t *testing.T) {
			t.Parallel()

			images, err := imagex.FetchBulk(context.Background(), tt.urls)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				require.NoError(t, err)

				assert.Equal(t, len(tt.urls), len(images))
				for _, img := range images {
					checksum, ok := checksums[img.Metadata.FileName]
					assert.True(t, ok)
					assert.Equal(t, checksum, img.Checksum())
				}
			}
		})
	}
}
