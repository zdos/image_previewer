package image_service

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

type ImageDownloadError struct {
	Err error
}

func (e *ImageDownloadError) Error() string {
	return "image download error:" + e.Err.Error() + "\n"
}

func downloadImageByURL(ctx context.Context, url string) (io.ReadCloser, http.Header, error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, nil, fmt.Errorf("downloading request error: %w\n", err)
	}
	client := http.DefaultClient
	response, err := client.Do(request)
	if err != nil {
		return nil, nil, fmt.Errorf("downloading request error: %w\n", err)
	}
	if response.Header.Get("content-type") != "image/jpeg" {
		return nil, nil, &ImageDownloadError{Err: fmt.Errorf("%s\n", "URL not contain JPEG image")}
	}
	originResponseHeader := response.Header.Clone()
	return response.Body, originResponseHeader, nil
}
