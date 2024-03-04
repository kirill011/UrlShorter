package transport

import (
	"UrlShorter/internal/services"
	"UrlShorter/internal/transport/rest"
)

type Transport struct {
	H rest.Handlers
}

func NewTransport(s services.Service) *Transport {
	return &Transport{
		H: rest.NewHandlers(s),
	}
}
