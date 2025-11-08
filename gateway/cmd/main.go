package main

import (
	_ "gateway/cmd/docs"
	"gateway/config"
	"gateway/copyrights"
	"gateway/internal/grpc/products"
	"gateway/internal/net/echo"
	"gateway/internal/pkg/redis"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// @title Volha gateway REST API
// @version 0.1

// @contact.name Alex "bustard" Provor
// @contact.url https://breezy.su
// @contact.email help@breezy.su

// @host localhost:8080
// @BasePath /
// @schemes http
func main() {

	const op = "main"

	cfg := config.MustSetup()
	p, err := products.New(cfg)
	if err != nil {
		log.Panic(err)
	}

	rds := redis.New(cfg)

	e := echo.New(rds, p, cfg)
	go e.MustRun()

	if cfg.Mode != "DEV" {
		if err := copyrights.Info(); err != nil {
			log.Println(err)
		}
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	sign := <-stop

	if err := e.Stop(); err != nil {
		log.Println(err)
	}

	log.Println(op, "stop signal", sign)
}
