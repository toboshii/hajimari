package handlers

import (
	"net/http"

	"github.com/toboshii/hajimari/internal/log"
)

var (
	logger = log.New()
)

type Handler interface {
	Handle(responseWriter http.ResponseWriter, request *http.Request)
}
