package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/zdos/image_previewer/internal/handler"
	service "github.com/zdos/image_previewer/internal/services/image_service"
	lruCache "github.com/zdos/image_previewer/internal/services/lru_cache"
	"github.com/zdos/image_previewer/internal/storage"
	"go.uber.org/zap"
)

func main() {
	//TODO Add Graceful Shutdown
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{"stdout"}
	zapLogger, err := config.Build()
	if err != nil {
		fmt.Printf("error was occured while configuring logger: %s", err.Error())
		os.Exit(1)
	}
	zap.ReplaceGlobals(zapLogger)
	newCacheInstance := lruCache.NewCache(5)                                                //TODO configure cache size
	filePathForStoreImages := "/home/koryan/go-public-projects/image-previewer/image_cache" //TODO configure stope path
	storage := storage.NewLocalImageStorage(filePathForStoreImages)                         //TODO configure storage type
	imageService := service.NewImageProcessService(newCacheInstance, storage)
	handler := handler.NewRouter(imageService)

	if err := http.ListenAndServe(":8888", handler); err != nil { //TODO configure HTTP adress:port
		zap.S().Fatalf("%s\n", err.Error())
	}
}
