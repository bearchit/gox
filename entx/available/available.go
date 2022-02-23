package available

import "time"

type QueryOption struct {
	Preview bool
	At      time.Time
}

type QueryOptionFunc func(*QueryOption)

func WithPreview(preview bool) QueryOptionFunc {
	return func(qo *QueryOption) {
		qo.Preview = preview
	}
}

func WithAt(at time.Time) QueryOptionFunc {
	return func(qo *QueryOption) {
		qo.At = at
	}
}
