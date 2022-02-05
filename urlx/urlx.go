package urlx

import "net/url"

func ParseStringURLs(urls []string) ([]*url.URL, error) {
	result := make([]*url.URL, len(urls))
	for ui, u := range urls {
		if purl, err := url.Parse(u); err != nil {
			return nil, err
		} else {
			result[ui] = purl
		}
	}

	return result, nil
}
