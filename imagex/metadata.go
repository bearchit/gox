package imagex

import (
	"io"
	"mime"
	"path/filepath"

	stdimage "image"

	_ "golang.org/x/image/webp"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

type Metadata struct {
	FileName string `json:"file_name"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	Format   string `json:"format"`
	Mime     string `json:"mime"`
}

func parseMetadata(r io.Reader, path string) (*Metadata, error) {
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
