package main

import (
	"github.com/starptech/go-web/server"
)

func main() {
	config := server.NewConfig()
	logger := server.NewLogger(config.GrayLogAddr, config.IsProduction)

	engine := server.NewEngine(config)
	engine.SetLogger(logger)
	engine.ServeStaticFiles()

	m := server.Migration{Db: engine.Db}
	m.Up()

	go func() {
		logger.Fatal(engine.Start(config.Address))
	}()

	engine.GracefulShutdown()
}
