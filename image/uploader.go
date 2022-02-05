package image

import "context"

type Uploader interface {
	Upload(ctx context.Context, image Image) (string, error)
}
