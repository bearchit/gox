package image

import "context"

type uploader interface {
	Upload(ctx context.Context, image Image) (string, error)
}
