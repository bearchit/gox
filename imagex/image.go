package imagex

import (
	"crypto/md5"
	"encoding/hex"
)

type Image struct {
	Metadata Metadata
	Content  []byte
}

func (img Image) IsZero() bool {
	return img.Content == nil
}

func (img Image) Ratio() float64 {
	return float64(img.Metadata.Width) / float64(img.Metadata.Height)
}

func (img Image) Checksum() string {
	h := md5.New()
	h.Write(img.Content)
	return hex.EncodeToString(h.Sum(nil))
}
