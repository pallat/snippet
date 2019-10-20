package echo_test

import (
	"net/http"
	"testing"

	"github.com/labstack/echo/v4/middleware"
	"github.com/pallat/snippet/go/echo"
	"github.com/sirupsen/logrus"
)

func hello(c echo.Context) error {
	c.Log(logrus.InfoLevel, "this is logrus")
	return c.String(http.StatusOK, "Hello, World!")
}

func TestUsage(t *testing.T) {
	e := echo.New()
	e.SetLogrus(logrus.New())

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Add(http.MethodGet, "/", hello)

	e.Logger.Fatal(e.Start(":1323"))
}
