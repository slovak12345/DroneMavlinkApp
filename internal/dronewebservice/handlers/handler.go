package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/bluenviron/gomavlib/v2"
	"github.com/bluenviron/gomavlib/v2/pkg/dialects/ardupilotmega"
	"github.com/bluenviron/gomavlib/v2/pkg/dialects/asluav"
	"github.com/bluenviron/gomavlib/v2/pkg/dialects/pythonarraytest"
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

func (h *Handler) StartEnpointUdpServerHandle(httpHost, httpPort, dialect string, OutSystemID byte) {

	Address := fmt.Sprintf("%s:%s", httpHost, httpPort)

	Dialect := ardupilotmega.Dialect

	switch dialect {

	case "ardupilotmega":
		Dialect = ardupilotmega.Dialect
	case "pythonarraytest":
		Dialect = pythonarraytest.Dialect
	case "asluav":
		Dialect = asluav.Dialect
	}

	// create a node which communicates with a UDP endpoint in server mode
	node, err := gomavlib.NewNode(gomavlib.NodeConf{
		Endpoints: []gomavlib.EndpointConf{
			gomavlib.EndpointUDPServer{Address: Address},
		},
		Dialect:     Dialect,
		OutVersion:  gomavlib.V2, // change to V1 if you're unable to communicate with the target
		OutSystemID: OutSystemID,
	})
	if err != nil {
		h.log.Error(fmt.Sprintf("Endpoint UDP Server with %v couldn't start:", Address), zap.Error(err))
	}

	h.log.Info(fmt.Sprintf("Udp endpoint server with id:%v and adress:%v starting with dialect:%v", OutSystemID, Address, dialect))

	// print incoming messages
	for evt := range node.Events() {
		if frm, ok := evt.(*gomavlib.EventFrame); ok {
			log.Printf("received: id=%d, %+v\n", frm.Message().GetID(), frm.Message())
		}
	}
}
