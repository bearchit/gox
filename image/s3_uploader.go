package image

import (
	"bytes"
	"context"
	"path/filepath"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type s3Uploader struct {
	client *s3.Client
	bucket string
	path   string
}

func NewS3Uploader(
	client *s3.Client,
	bucket string,
	path string,
) *s3Uploader {
	return &s3Uploader{
		client: client,
		bucket: bucket,
		path:   path,
	}
}

var _ uploader = (*s3Uploader)(nil)

func (u s3Uploader) Upload(ctx context.Context, image Image) (string, error) {
	uploader := manager.NewUploader(u.client)

	key := image.Checksum()
	buf := bytes.NewBuffer(image.Content)
	result, err := uploader.Upload(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(u.bucket),
		Key:         aws.String(filepath.Join(u.path, key)),
		Body:        buf,
		ContentType: aws.String(image.Metadata.Mime),
	})
	if err != nil {
		return "", err
	}

	return result.Location, nil
}
