package server

import "github.com/google/wire"

var ProviderSetServer = wire.NewSet(
	NewAppServer,
)
