package service

import (
	"gitlab.com/dpcat237/flisy/config"
	"gitlab.com/dpcat237/flisy/src/adapter/repository"
	"gitlab.com/dpcat237/flisy/src/module/flight"
	"gitlab.com/dpcat237/flisy/src/module/seat"
)

type Collector struct {
	FlHnd flight.Handler
	StHnd seat.Handler
}

func Init(cfg config.Config) Collector {
	db := repository.InitConnectionDb(cfg.DbHost, cfg.DbUser, cfg.DbPassword, cfg.DbName)

	//Init repositories
	fRepo := flight.NewRepository(db)
	stRepo := seat.NewRepository(db)

	//Init handlers
	fHnd := flight.NewHandler(fRepo)
	stHnd := seat.NewHandler(stRepo)

	return Collector{
		FlHnd: fHnd,
		StHnd: stHnd,
	}
}
