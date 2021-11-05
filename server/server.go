package server

import (
	"conjur/logger"
	"conjur/module"
	"context"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

type Server struct {
	*module.Module
	Service *http.Server
}

func MakeServer() *Server {
	server := Server{}
	server.Module = module.MakeModule("Server", server.Init, server.DeInit)
	router := mux.NewRouter()
	router.StrictSlash(true)
	RegisterRoutes(router.PathPrefix("/api").Subrouter())
	router.PathPrefix("/").Handler(SpaHandler{
		StaticPath: "client/build",
		IndexPath:  "index.html",
	})
	server.Service = &http.Server{
		Handler: handlers.CombinedLoggingHandler(os.Stdout, router),
		Addr:    ":8080",
	}
	return &server
}

func (server *Server) Init() {
	go func() {
		logger.Info(fmt.Sprintf("Server starting at port %s", server.Service.Addr))
		shutdown := server.Service.ListenAndServe()
		if shutdown != nil {
			if shutdown == http.ErrServerClosed {
				return
			}
			logger.Error(fmt.Sprintf("Server closed with exceptions: %s", shutdown.Error()))
			return
		}
		logger.Error("Server closed with exceptions")
	}()
}

func (server *Server) DeInit() {
	ctx := context.Background()
	err := server.Service.Shutdown(ctx)
	if err != nil {
		logger.Error(fmt.Sprintf("Server closed with exceptions: %s", err.Error()))
	}
}
