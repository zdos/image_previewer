package image_service

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"io"

	"github.com/nfnt/resize"
	"go.uber.org/zap"
)

var (
	resizeErrPrefix = "error resize: %s\n"
)

func resizeImage(width int, height int, imageFile io.Reader) ([]byte, error) {
	originImageData, _, err := image.Decode(imageFile)
	if err != nil {
		return nil, fmt.Errorf(resizeErrPrefix, err.Error())
	}
	resizedImageBuf := new(bytes.Buffer)
	resizedImage := resize.Resize(uint(width), uint(height), originImageData, resize.Lanczos3)
	err = jpeg.Encode(resizedImageBuf, resizedImage, nil)
	if err != nil {
		return nil, fmt.Errorf(resizeErrPrefix, err.Error())
	}
	zap.S().Infof("image resized")
	return resizedImageBuf.Bytes(), nil
}
