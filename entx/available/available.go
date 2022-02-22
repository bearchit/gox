package available

type QueryOption struct {
	Preview bool
}

type QueryOptionFunc func(*QueryOption)

func WithPreview(preview bool) QueryOptionFunc {
	return func(qo *QueryOption) {
		qo.Preview = preview
	}
}
