package server

import (
	wiregin "github.com/RyuChk/wire-provider/gin"
	"github.com/RyuChk/wire-provider/gin/provider"

	"github.com/ZecretBone/ips-bff/cmd/bff-api/handler"
	"github.com/ZecretBone/ips-bff/routers"
)

type ginRouterCustomizer struct {
	handler *handler.Handlers
}

func (gc *ginRouterCustomizer) Register(server *wiregin.Server) error {
	routers.RegisterRouter(server.Engine, *gc.handler)
	return nil
}

func (gc *ginRouterCustomizer) Configure(builder wiregin.Builder) error {
	return nil
}

func ProvideGinRouterCustomizer(handler handler.Handlers) provider.RouterCustomizer {
	return &ginRouterCustomizer{
		handler: &handler,
	}
}
