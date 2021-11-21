package storage

import (
	"fmt"
	"io"
	"os"
)

var (
	storageErrPrefix = "error storage: %s\n"
)

type imageStorage struct {
	storagePath string
}

func NewLocalImageStorage(storagePath string) *imageStorage {
	return &imageStorage{
		storagePath: storagePath,
	}
}

func (ls *imageStorage) GetFileByHashname(filename string) (io.ReadCloser, error) {
	imageFile, err := os.Open(fmt.Sprintf("%s/%s", ls.storagePath, filename)) //TODO make it clear
	if err != nil {
		return nil, fmt.Errorf(storageErrPrefix, err.Error())
	}
	return imageFile, nil
}

func (ls *imageStorage) SaveNewImage(filenameHash string, imageDataStream io.Reader) error {
	imageFile, err := os.Create(fmt.Sprintf("%s/%s", ls.storagePath, filenameHash)) //TODO make it clear
	if err != nil {
		return fmt.Errorf(storageErrPrefix, err.Error())
	}
	defer imageFile.Close()
	_, err = io.Copy(imageFile, imageDataStream)
	if err != nil {
		return fmt.Errorf(storageErrPrefix, err.Error())
	}
	return nil
}
