package dao

import "github.com/google/wire"

var ProviderSetDao = wire.NewSet(
	NewDB,
	NewEngine,
)
