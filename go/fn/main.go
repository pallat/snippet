package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

var router = echo.New()

func init() {
	router.Logger.SetLevel(log.INFO)
}

func main() {
	router.GET("/", func(c echo.Context) error {
		time.Sleep(5 * time.Second)
		return c.JSON(http.StatusOK, "OK")
	})

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
