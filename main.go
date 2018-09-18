package main

import (
	logger "github.com/sirupsen/logrus"

	"gitlab.com/dpcat237/flisy/config"
	"gitlab.com/dpcat237/flisy/src/router"
	"gitlab.com/dpcat237/flisy/src/router/controller"
	"gitlab.com/dpcat237/flisy/src/service"
)

func main() {
	cfg := config.LoadConfigData()
	sCll := service.Init(cfg)
	cc := controller.Init(sCll)
	r := router.New(cc, cfg.Port, cfg.CertFile, cfg.KeyFile)

	logger.Infoln("Router running at port", cfg.Port)
	r.Init()
}
