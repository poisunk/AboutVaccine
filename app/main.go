package main

import "vax/internal/config"

func main() {
	Execute()
}

func RunApp() {
	conf, err := config.ReadConfig()
	if err != nil {
		panic(err)
	}
	app, err := InitApplication(conf.Server, conf.Database)
	if err != nil {
		panic(err)
	}
	err = app.Run()
	if err != nil {
		panic(err)
	}
}
