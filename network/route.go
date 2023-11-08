package network

import "github.com/DarthPestilane/easytcp"

type RouteInfo interface {
	HandleRequest(c easytcp.Context)
}
