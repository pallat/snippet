package payment

import "net/http"

type clienter interface {
	Do(*http.Request) (*http.Response, error)
}

func NewClient() clienter {
	return http.DefaultClient
}
