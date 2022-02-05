package image

import (
	"bytes"
	"context"
	stdimage "image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"io/ioutil"
	"mime"
	"net/http"
	"net/url"
	"path/filepath"
	"time"

	"github.com/pkg/errors"

	_ "golang.org/x/image/webp"
)

func parseStringURLs(urls []string) ([]*url.URL, error) {
	result := make([]*url.URL, len(urls))
	for ui, u := range urls {
	}
}

func FetchFromURLBulk(ctx context.Context, rawURLs []string) ([]*Image, error) {
	hc := http.Client{
		Timeout: time.Second * 3,
	}

	urls := make([]*url.URL, len(rawURLs))
	for ui, u := range rawURLs {
		if v, err := url.Parse(u); err != nil {
			return nil, err
		} else {
			urls[ui] = v
		}
	}

	result := make([]*Image, len(urls))

	for i, u := range urls {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
		if err != nil {
			return nil, err
		}

		resp, err := hc.Do(req)
		if err != nil {
			return nil, err
		}

		content, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		if err := resp.Body.Close(); err != nil {
			return nil, err
		}

		md, err := parseConfig(bytes.NewBuffer(content), u.Path)
		if err != nil {
			return nil, err
		}

		result[i] = &Image{
			Content:  content,
			Metadata: *md,
		}
	}

	return result, nil
}

func FetchFromURL(ctx context.Context, rawURL string) (*Image, error) {
	images, err := FetchFromURLBulk(ctx, []string{rawURL})
	if err != nil {
		return nil, errors.Wrapf(err, "failed fetch image from an URL [%v]", rawURL)
	}

	return images[0], nil
}

func FromReader(r io.Reader, path string) (*Image, error) {
	content, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	metadata, err := parseConfig(bytes.NewBuffer(content), path)
	if err != nil {
		return nil, err
	}

	return &Image{
		Content:  content,
		Metadata: *metadata,
	}, nil
}

func parseConfig(r io.Reader, path string) (*Metadata, error) {
	cfg, format, err := stdimage.DecodeConfig(r)
	if err != nil {
		return nil, err
	}
	_, filename := filepath.Split(path)

	metadata := &Metadata{
		FileName: filename,
		Format:   format,
		Width:    cfg.Width,
		Height:   cfg.Height,
	}

	if v := mime.TypeByExtension("." + format); v != "" {
		metadata.Mime = v
	}

	return metadata, nil
}
