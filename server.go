package application

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, hander http.Handler) error {
	s.httpServer = &http.Server{
		Addr: ":" + port,
		Handler: hander, 
		MaxHeaderBytes: 1 << 20,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}