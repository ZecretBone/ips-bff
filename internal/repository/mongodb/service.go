package mongodb

import (
	wiremongo "github.com/RyuChk/wire-provider/mongodb"
)

// ProvideMongoDBService
// @@wire-set-name@@ name:"DatabaseSet"
func ProvideMongoDBService(config wiremongo.Config) (wiremongo.Connection, func(), error) {
	conn, err := wiremongo.NewConnection(&config)
	if err != nil {
		return nil, func() {}, nil
	}
	if conn == nil {
		panic(err)
	}
	return conn, func() {
		conn.Disconnect()
	}, nil
}
