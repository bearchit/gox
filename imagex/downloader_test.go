package imagex_test

import (
	"github.com/bearchit/gox/imagex"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

func TestNewHttpDownloader(t *testing.T) {
	client := http.Client{Timeout: time.Second * 10}
	downloader := imagex.NewHttpDownloader(client)
	assert.Equal(t, client, downloader.Client())
}
