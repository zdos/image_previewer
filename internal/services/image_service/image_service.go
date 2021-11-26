package image_service

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"

	"go.uber.org/zap"
)

type imageStorage interface {
	GetFileByHashname(string) (io.ReadCloser, error)
	SaveNewImage(string, io.Reader) error
}

type imageCacheData struct {
	imageHashname             string
	imageOriginResponseHeader http.Header
}

type imageCache interface {
	Set(key string, value interface{}) bool
	Get(key string) (interface{}, bool)
}

type imageService struct {
	imageCache   imageCache
	imageStorage imageStorage
}

func NewImageProcessService(cache imageCache, storage imageStorage) *imageService {
	return &imageService{
		imageCache:   cache,
		imageStorage: storage,
	}
}

func (s *imageService) FindAndResize(ctx context.Context, width, height int, url string) ([]byte, http.Header, error) {
	cacheValue, isExist := s.imageCache.Get(url)

	if isExist {
		zap.S().Infof("image by url %s was found in cache\n", url)
		cachedImageData, ok := cacheValue.(imageCacheData)
		if !ok {
			return nil, nil, fmt.Errorf("error cache: cannot get image from cache")
		}
		imageFile, err := s.imageStorage.GetFileByHashname(cachedImageData.imageHashname)
		if err != nil {
			return nil, nil, err
		}
		defer imageFile.Close()
		imageData, err := resizeImage(width, height, imageFile)
		if err != nil {
			return nil, nil, err
		}
		return imageData, nil, nil
	}
	imageFile, originResponseHeader, err := downloadImageByURL(ctx, url)
	if err != nil {
		return nil, nil, err
	}
	defer imageFile.Close()
	hashedFilename := hashByURL(url)
	if err = s.imageStorage.SaveNewImage(hashedFilename, imageFile); err != nil {
		return nil, nil, err
	}
	newImageCacheItem := imageCacheData{
		imageHashname:             hashedFilename,
		imageOriginResponseHeader: originResponseHeader,
	}
	s.imageCache.Set(url, newImageCacheItem)
	zap.S().Infof("new image by url %s added to cache\n", url)
	localImageFile, err := s.imageStorage.GetFileByHashname(hashedFilename)
	if err != nil {
		return nil, nil, err
	}
	defer localImageFile.Close()
	resizedImage, err := resizeImage(width, height, localImageFile)
	if err != nil {
		return nil, nil, err
	}
	return resizedImage, originResponseHeader, nil
}

func hashByURL(url string) string {
	hashInstance := sha1.New()
	hashInstance.Write([]byte(url))
	hash := hex.EncodeToString(hashInstance.Sum(nil))
	return hash
}
