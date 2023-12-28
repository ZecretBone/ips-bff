package di

import (
	"github.com/RyuChk/wire-provider/gin/provider"
	"github.com/ZecretBone/ips-bff/cmd/bff-api/handler"
	"github.com/ZecretBone/ips-bff/cmd/bff-api/server"
	"github.com/google/wire"
)

var CustomizerSet = wire.NewSet(
	server.ProvideGinRouterCustomizer,
)

var ProviderSet = wire.NewSet(
	handler.ProvideRSSIHandler,
	wire.Struct(new(handler.Handlers), "*"),
)

type Locator struct {
	Handlers            *handler.Handlers
	GinServerCustomizer provider.RouterCustomizer
	RSSIHandler         handler.RSSIHandler
}
