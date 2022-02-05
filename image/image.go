package image

import (
	"crypto/md5"
	"encoding/hex"
)

type Image struct {
	Metadata Metadata
	Content  []byte

	checksum string
}

var ZeroImage = Image{}

func (img Image) IsZero() bool {
	return img.Content == nil
}

func (img Image) Ratio() float64 {
	return float64(img.Metadata.Width) / float64(img.Metadata.Height)
}

func (img *Image) Checksum() string {
	if img.checksum == "" {
		h := md5.New()
		h.Write(img.Content)
		img.checksum = hex.EncodeToString(h.Sum(nil))
	}

	return img.checksum
}
