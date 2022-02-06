package imagex

import (
	"context"
	"net/url"
	"sync"

	urlx "github.com/bearchit/gox/urlx"
)

type Fetcher struct {
	downloader downloader
}

func NewFetcher(
	downloader downloader,
) *Fetcher {
	return &Fetcher{
		downloader: downloader,
	}
}

func (fetcher Fetcher) fetch(ctx context.Context, u *url.URL) (Image, error) {
	reader, err := fetcher.downloader.Get(ctx, u)
	if err != nil {
		return Image{}, err
	}

	return Read(reader, u.Path)
}

const (
	maxFetchWorkers = 10
)

func (fetcher Fetcher) Fetch(ctx context.Context, rawURL string) (Image, error) {
	var imageURL *url.URL
	if v, err := urlx.ParseStringURLs([]string{rawURL}); err != nil {
		return ZeroImage, err
	} else {
		imageURL = v[0]
	}

	return fetcher.fetch(ctx, imageURL)
}

func (fetcher Fetcher) FetchBulk(ctx context.Context, rawURLs []string) ([]Image, error) {
	urls, err := urlx.ParseStringURLs(rawURLs)
	if err != nil {
		return nil, err
	}

	jobs := make(chan *url.URL, 100)
	images := make(chan Image, 100)
	errs := make(chan error, 100)

	var wg sync.WaitGroup
	for w := 1; w <= maxFetchWorkers; w++ {
		wg.Add(1)

		go func(jobs <-chan *url.URL, results chan<- Image, errs chan<- error) {
			defer wg.Done()

			for j := range jobs {
				if image, err := fetcher.fetch(ctx, j); err != nil {
					errs <- err
				} else {
					images <- image
				}
			}
		}(jobs, images, errs)
	}

	for _, u := range urls {
		jobs <- u
	}
	close(jobs)
	wg.Wait()

	close(images)
	close(errs)

	for e := range errs {
		return nil, e
	}

	result := make([]Image, 0, len(urls))
	for img := range images {
		result = append(result, img)
	}

	return result, nil
}

var DefaultFetcher = NewFetcher(DefaultHttpDownloader)

func Fetch(ctx context.Context, rawURL string) (Image, error) {
	return DefaultFetcher.Fetch(ctx, rawURL)
}

func FetchBulk(ctx context.Context, rawURLs []string) ([]Image, error) {
	return DefaultFetcher.FetchBulk(ctx, rawURLs)
}
