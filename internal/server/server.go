package server

import (
	"vax/internal/config"
	"vax/internal/router"
)

type AppServer struct {
	ApiRouter  *router.ApiRouter
	ServerConf *config.Server
}

func NewAppServer(
	apiRouter *router.ApiRouter,
	serverConf *config.Server,
) *AppServer {
	return &AppServer{
		ApiRouter:  apiRouter,
		ServerConf: serverConf,
	}
}

func (_this *AppServer) Run() error {
	engine := _this.ApiRouter.InitEngine()
	err := engine.Run(_this.ServerConf.Port)
	return err
}
