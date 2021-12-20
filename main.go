package main

import (
	"github.com/Striker87/Banking/app"
	"github.com/Striker87/Banking/logger"
)

func main() {
	logger.Info("Starting the application...")
	app.Start()
}
