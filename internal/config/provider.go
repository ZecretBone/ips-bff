package config

import (
	wireminio "github.com/RyuChk/wire-provider/minio"
	wiremongo "github.com/RyuChk/wire-provider/mongodb"
	"github.com/kelseyhightower/envconfig"
)

func ProvideMongoxConfig() wiremongo.Config {
	return provideConfig(wiremongo.Config{})
}

func ProvideGRPCServiceConfig() GRPCConfig {
	return provideConfig(GRPCConfig{})
}

func ProvideMinioXConfig() wireminio.Config {
	return provideConfig(wireminio.Config{})
}

func ProvideCacheConfig() CacheConfig {
	return provideConfig(CacheConfig{})
}

func ProvideRSSIConfig() RSSIServiceConfig {
	return provideConfig(RSSIServiceConfig{})
}

func provideConfig[T any](cfg T) T {
	envconfig.MustProcess("", &cfg)
	return cfg
}
