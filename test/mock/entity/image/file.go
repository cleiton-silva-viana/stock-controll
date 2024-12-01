package imagemock

import (
	"bytes"
	"image"
	"image/gif"
	"image/jpeg"
	"io"
	"mime/multipart"
)

type ByteFile struct {
	*bytes.Reader
}

func (b *ByteFile) Close() error {
	return nil
}

type Encode func(w io.Writer, m image.Image) error

func File(encode Encode) (multipart.File, error) {
	var buffer bytes.Buffer
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
	err := encode(&buffer, img)
	if err != nil {
		return nil, err
	}
	return &ByteFile{bytes.NewReader(buffer.Bytes())}, nil
}

func FileJPEG() (multipart.File, error) {
	return File(func(w io.Writer, m image.Image) error {
		return jpeg.Encode(w, m, nil)
	})
}

func FileGIF() (multipart.File, error) {
	return File(func(w io.Writer, m image.Image) error {
		return gif.Encode(w, m, nil)
	})
}
