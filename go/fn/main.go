package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/pallat/snippet/fn/payment"
)

var router = echo.New()

func init() {
	router.Logger.SetLevel(log.INFO)
}

func main() {
	router.GET("/", payment.New())

	go start()

	shutdown()
}

func start() {
	if err := router.Start(":1323"); err != nil {
		router.Logger.Info("shutting down the server")
	}
}

func shutdown() {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	if err := router.Shutdown(context.Background()); err != nil {
		router.Logger.Fatal(err)
	}
}
