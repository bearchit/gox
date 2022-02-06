package urlx

import "net/url"

func ParseStringURLs(urls []string) ([]*url.URL, error) {
	result := make([]*url.URL, len(urls))
	for ui, u := range urls {
		purl, err := url.Parse(u)
		if err != nil {
			return nil, err
		}

		result[ui] = purl
	}

	return result, nil
}
