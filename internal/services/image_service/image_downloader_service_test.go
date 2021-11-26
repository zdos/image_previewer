package image_service

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestErrorDownloadImageByURL(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	url := "someurl"
	_, _, err := downloadImageByURL(ctx, url)
	require.Error(t, err)
}

func TestSuccessDownloadImageByURL(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	url := "https://cdn.maikoapp.com/3d4b/4quqa/150.jpg"
	body, _, err := downloadImageByURL(ctx, url)
	require.NoError(t, err)
	require.NotNil(t, body)
}
