package inMem

import (
	"UrlShorter/internal/models"
	"strings"
)

type InMemory struct {
	M map[string]string
}

func (base *InMemory) GetUrl(url models.Url) (models.Url, error) {

	full, ok := base.M[url.ShortUrl]
	if !ok {
		return models.Url{}, models.ErrUrlNotFound
	}

	url.FullUrl = full

	return url, nil
}

func (base *InMemory) CreateUrl(url models.Url) error {
	base.M[url.ShortUrl] = url.FullUrl

	return nil
}

func (base *InMemory) GetMaxUrl() (models.Url, error) {

	var maxUrl string = ""

	for val := range base.M {
		if len(val) > len(maxUrl) {
			maxUrl = val
			continue
		} else if len(val) < len(maxUrl) {
			continue
		} else {
			for i, v := range val {
				if strings.Index(models.Alpabet, string(v)) > strings.Index(models.Alpabet, string(maxUrl[i])) {
					maxUrl = val
				}
			}
		}
	}
	return models.Url{ShortUrl: maxUrl}, nil
}

func NewInMemory() *InMemory {
	return &InMemory{M: make(map[string]string)}
}
