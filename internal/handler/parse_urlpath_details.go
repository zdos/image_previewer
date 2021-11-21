package handler

import (
	"errors"
	"strconv"
	"strings"
)

var (
	errWidthOrHeightNotSpecified = errors.New("url parse error: width or height not correctly specified")
	errOriginImageURLEmpty       = errors.New("url parse error: origin image url empty")
	errNotJpgFormat              = errors.New("url parse error: origin image is not in jpeg/jpg format")
)

type imageDetails struct {
	width           int
	height          int
	originImageURL  string
	originImageName string
}

//TODO Refactor code (make clear)
// Example : /200/300/https://vk.com/profile/avatar.jpg
// Output should be:
// {
// 	width: 200
// 	height: 300
// 	originImageURL: https://vk.com/profile/avatar.jpg
// 	originImageName: avatar.jpg
// }
func ParseLinkDetails(p string) (*imageDetails, error) {
	newImageDetails := new(imageDetails)

	width, restPart := splitBySlashes(p[1:])
	widthInt, err := strconv.Atoi(width)
	if err != nil {
		return newImageDetails, errWidthOrHeightNotSpecified
	}
	newImageDetails.width = widthInt
	heightString, restPart := splitBySlashes(restPart)
	height, err := strconv.Atoi(heightString)
	if err != nil {
		return newImageDetails, errWidthOrHeightNotSpecified
	}
	newImageDetails.height = height
	if len(restPart) == 0 {
		return newImageDetails, errOriginImageURLEmpty
	}
	newImageDetails.originImageURL = restPart
	fileName, _ := splitBySlashes(reverse(restPart))
	newImageDetails.originImageName = reverse(fileName)
	if !checkFileFormatToJPEG(newImageDetails.originImageName) {
		return newImageDetails, errNotJpgFormat
	}
	return newImageDetails, nil
}

func splitBySlashes(p string) (string, string) {
	newPartIndex := strings.Index(p, "/")
	if newPartIndex > 0 {
		newPartString, p := p[:newPartIndex], p[newPartIndex+1:]
		return newPartString, p
	}
	return p, ""
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func checkFileFormatToJPEG(s string) bool {
	dotIndex := strings.Index(s, ".")
	if dotIndex > 0 {
		if s[dotIndex:] == ".jpg" || s[dotIndex:] == ".jpeg" {
			return true
		}
	}
	return false
}
