package network

import (
	"fmt"
	"log"

	"github.com/DarthPestilane/easytcp"
)

type Server struct {
	listener *easytcp.Server
}

func NewListener() *Server {
	return &Server{
		listener: easytcp.NewServer(&easytcp.ServerOption{
			Packer:           easytcp.NewDefaultPacker(),
			Codec:            nil,
			DoNotPrintRoutes: true,
		}),
	}
}

func (s *Server) Run(port int) {
	log.Printf("Start network server port: %d", port)
	if err := s.listener.Run(fmt.Sprintf(":%d", port)); err != nil {
		log.Fatalf("serve err: %s", err)
	}
}

func (s *Server) RegisterMiddlewares(middleware easytcp.MiddlewareFunc) {
	s.listener.Use(middleware)
}

func (s *Server) RegisterRoutes(routes map[uint16]RouteInfo) {
	// register all routes
	for key, route := range routes {
		s.listener.AddRoute(key, route.HandleRequest)
	}
}
