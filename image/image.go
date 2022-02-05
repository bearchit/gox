package image

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"strconv"
)

type Image struct {
	Metadata Metadata
	Content  []byte
}

func (img Image) Ratio() float64 {
	return float64(img.Metadata.Width) / float64(img.Metadata.Height)
}

func (img Image) Checksum() string {
	h := md5.New()
	h.Write(img.Content)
	return hex.EncodeToString(h.Sum(nil))
}

type Metadata struct {
	FileName string `json:"file_name"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	Format   string `json:"format"`
	Mime     string `json:"mime"`
}

type Size string

const (
	SizeXs Size = "XS"
	SizeSm Size = "SM"
	SizeMd Size = "MD"
	SizeLg Size = "LG"
	SizeXl Size = "XL"
)

var sizes = []Size{
	SizeXs,
	SizeSm,
	SizeMd,
	SizeLg,
	SizeXl,
}

func (e Size) IsValid() bool {
	for _, x := range sizes {
		if e == x {
			return true
		}
	}
	return false
}

func (e Size) String() string {
	return string(e)
}

func (e *Size) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Size(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Size", str)
	}
	return nil
}

func (e Size) MarshalGQL(w io.Writer) {
	if _, err := fmt.Fprint(w, strconv.Quote(e.String())); err != nil {
		panic(err)
	}
}
