package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/slovak12345/DroneMavlinkApp/pkg/graceful"
	"go.uber.org/zap"
)

type Handler struct {
	log *zap.Logger
}

func NewHandler(log *zap.Logger) *Handler {
	return &Handler{
		log: log,
	}
}

func (h *Handler) StartHandle(httpHost, httpPort string) {
	server := &http.Server{
		Addr:    httpHost + ":" + httpPort,
		Handler: h.newRouter(),
	}

	h.log.Info(fmt.Sprintf("Server is running %v", server.Addr))

	graceful.AddCallback(func() error {
		return server.Shutdown(context.Background())
	})

	err := server.ListenAndServe()
	if err != nil {
		h.log.Error("Server shutdown failed", zap.Error(err))
		graceful.ShutdownNow()

		return
	}
}
