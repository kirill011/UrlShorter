package rest

import (
	"UrlShorter/internal/services"
	"net/http"
	"regexp"

	"github.com/labstack/echo/v4"
)

type Handlers struct {
	s services.Service
}

func NewHandlers(s services.Service) Handlers {
	return Handlers{s: s}
}

func (h *Handlers) CreateUrl(c echo.Context) error {

	var FullUrl string

	err := c.Bind(&FullUrl)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	//Валидация полученного URL
	matched, _ := regexp.MatchString(`[-a-zA-Z0-9@:%_\+.~#?&\/=]{2,256}\.[a-z]{2,4}\b(\/[-a-zA-Z0-9@:%_\+.~#?&\/=]*)?`, FullUrl)

	if !matched {
		return c.JSON(http.StatusBadRequest, "Request body does not contain a URL")
	}

	short, err := h.s.CreateUrl(FullUrl)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.String(http.StatusOK, short)
}

func (h *Handlers) GetUrl(c echo.Context) error {

	shortUrl := c.Param("shortUrl")

	FullUrl, err := h.s.GetUrl(shortUrl)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.String(http.StatusOK, FullUrl)
}
