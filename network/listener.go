package network

import (
	"fmt"
	"log"

	"github.com/DarthPestilane/easytcp"
)

type Listener struct {
	tcp *easytcp.Server
}

func NewListener() *Listener {
	return &Listener{
		tcp: easytcp.NewServer(&easytcp.ServerOption{
			Packer:           easytcp.NewDefaultPacker(),
			Codec:            nil,
			DoNotPrintRoutes: true,
		}),
	}
}

func (l *Listener) Run(port int) {
	log.Printf("Start network server port: %d", port)
	if err := l.tcp.Run(fmt.Sprintf(":%d", port)); err != nil {
		log.Fatalf("serve err: %s", err)
	}
}

func (l *Listener) RegisterMiddlewares(middleware easytcp.MiddlewareFunc) {
	l.tcp.Use(middleware)
}

func (l *Listener) RegisterRoutes(routes map[uint16]RouteInfo) {
	// register all routes
	for key, route := range routes {
		l.tcp.AddRoute(key, route.HandleRequest)
	}
}
