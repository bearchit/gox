package image

import (
	"bytes"
	"io"
	"io/ioutil"
)

func Read(r io.Reader, path string) (Image, error) {
	content, err := ioutil.ReadAll(r)
	if err != nil {
		return Image{}, err
	}

	metadata, err := parseMetadata(bytes.NewBuffer(content), path)
	if err != nil {
		return Image{}, err
	}

	return Image{
		Content:  content,
		Metadata: *metadata,
	}, nil
}
