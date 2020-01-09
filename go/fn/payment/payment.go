package payment

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HandlerFunc func(request, FavoriteFunc) (response, error)

func New(h HandlerFunc, fav FavoriteFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req request
		err := c.Bind(&req)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
		}

		err = req.validate()
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
		}

		res, err := h(req, fav)
		if err != nil {
			return c.JSON(500, map[string]string{"message": err.Error()})
		}
		return c.JSON(http.StatusOK, res)
	}
}

type request struct{}

func (r *request) validate() error {
	return nil
}

type response struct {
	Message string `json:"message"`
}

func Handler(req request, fav FavoriteFunc) (response, error) {
	msg, err := fav("ok")
	if err != nil {
		return response{}, err
	}
	return response{Message: msg}, nil
}
