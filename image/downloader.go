package image

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type downloader interface {
	Get(ctx context.Context, u *url.URL) (io.Reader, error)
}

type httpDownloader struct {
	client http.Client
}

var _ downloader = (*httpDownloader)(nil)

func NewHttpDownloader() *httpDownloader {
	return &httpDownloader{
		client: http.Client{
			Timeout: time.Second * 3,
		},
	}
}

func (dn httpDownloader) Get(ctx context.Context, u *url.URL) (io.Reader, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := dn.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return bytes.NewBuffer(content), nil
}
