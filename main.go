package main

import (
	"github.com/lcaparada/gopportunities/config"
	"github.com/lcaparada/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	err := config.Init()

	if err != nil {
		logger.Errorf("Config initialization error: %v", err)
		panic(err)
	}

	router.Initialize()
}
