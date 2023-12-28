package di

import (
	wireminio "github.com/RyuChk/wire-provider/minio"
	wiremongo "github.com/RyuChk/wire-provider/mongodb"
	"github.com/ZecretBone/ips-bff/internal/config"
	"github.com/ZecretBone/ips-bff/internal/repository/cache"
	"github.com/ZecretBone/ips-bff/internal/repository/minio"
	"github.com/ZecretBone/ips-bff/internal/repository/mongodb"
	"github.com/ZecretBone/ips-bff/internal/services/rssi"
	"github.com/google/wire"
)

var DatabaseSet = wire.NewSet(
	minio.ProvideMinioService,
	mongodb.ProvideMongoDBService,
	cache.ProvideCacheService,
)

var ProviderSet = wire.NewSet(
	rssi.ProvideRSSIService,
)

var ConfigSet = wire.NewSet(
	config.ProvideMongoxConfig,
	config.ProvideMinioXConfig,
	config.ProvideCacheConfig,
	config.ProvideRSSIConfig,
)

type Locator struct {
	MongoDBConn  wiremongo.Connection
	MinioXConn   wireminio.Connection
	RSSIService  rssi.Service
	CacheService cache.Service
}
