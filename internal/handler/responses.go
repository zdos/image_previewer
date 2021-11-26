package handler

import (
	"net/http"

	"go.uber.org/zap"
)

func sendImage(w http.ResponseWriter, h http.Header, data []byte) {
	pasteOriginURLHeaders(w, h)
	w.Header().Set("Content-Disposition", "attachment")
	w.Write(data)
}

func pasteOriginURLHeaders(w http.ResponseWriter, h http.Header) {
	nv := 0
	for _, vv := range h {
		nv += len(vv)
	}
	for k, vv := range h {
		for _, v := range vv {
			w.Header().Add(k, v)
		}
	}
}

func httpBadRequest(w http.ResponseWriter, msg string, err error) {
	httpError(w, http.StatusBadRequest, msg, err)
}

func httpInternalServerError(w http.ResponseWriter, msg string, err error) {
	httpError(w, http.StatusInternalServerError, msg, err)
}

func httpBadGatewayError(w http.ResponseWriter, msg string, err error) {
	httpError(w, http.StatusBadGateway, msg, err)
}

func httpError(w http.ResponseWriter, httpStatus int, msg string, err error) {
	http.Error(w, msg, httpStatus)
	zap.L().Warn(msg, zap.Error(err))
}
