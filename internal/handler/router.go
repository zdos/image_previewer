package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/zdos/image_previewer/internal/services/image_service"
	"go.uber.org/zap"
)

type imageProcessService interface {
	FindAndResize(ctx context.Context, width, height int, url string) ([]byte, http.Header, error)
}

type Router struct {
	imageService imageProcessService
}

func NewRouter(imageService imageProcessService) http.Handler {
	return Router{
		imageService: imageService,
	}
}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/favicon.ico" {
		http.NotFound(w, r)
		return
	}
	if r.URL.Path == "/" {
		http.NotFound(w, r)
		return
	}
	router.ImageProcessHandler(w, r)
}

func (router Router) ImageProcessHandler(w http.ResponseWriter, r *http.Request) {
	jobDetails, err := ParseLinkDetails(r.URL.Path)
	if err != nil {
		httpBadRequest(w, err.Error(), err)
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	imageData, originHeader, err := router.imageService.FindAndResize(ctx, jobDetails.width, jobDetails.height, jobDetails.originImageURL)
	switch err.(type) {
	case nil:
	case *image_service.ImageDownloadError:
		{
			httpBadGatewayError(w, err.Error(), err)
			return
		}
	default:
		{
			httpInternalServerError(w, err.Error(), err)
			return
		}
	}
	sendImage(w, originHeader, imageData)
	zap.S().Infof("image sended")
}
