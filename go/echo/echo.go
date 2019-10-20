package echo

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type Context struct {
	echo.Context
	*logrus.Logger
}

type HandlerFunc func(Context) error

type Echo struct {
	*echo.Echo
	logger *logrus.Logger
}

func (e *Echo) SetLogrus(logger *logrus.Logger) {
	e.logger = logger
}

func (e *Echo) Add(method, path string, handler HandlerFunc, middleware ...echo.MiddlewareFunc) *echo.Route {
	next := construct(handler, e.logger)
	return e.Echo.Add(method, path, next, middleware...)
}

func New() *Echo {
	return &Echo{Echo: echo.New()}
}

func construct(h HandlerFunc, logger *logrus.Logger) echo.HandlerFunc {
	if logger == nil {
		logger = logrus.New()
	}
	return func(c echo.Context) error {
		nc := Context{
			Context: c,
			Logger:  logger,
		}
		return h(nc)
	}
}
