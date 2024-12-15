package handler

import (
	"net/http"
)

type LimiterHandler struct{}

func NewLimiterHandler() *LimiterHandler {
	return &LimiterHandler{}
}

func (h *LimiterHandler) HandleRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Request accepted"))
}
