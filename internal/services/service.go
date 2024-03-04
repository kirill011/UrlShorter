package services

import (
	"UrlShorter/internal/config"
	"UrlShorter/internal/models"
	"strings"

	"gorm.io/gorm"
)

type Service struct {
	creator UrlCreator
	getter  UrlGetter
}

var (
	StringUrl = "http://"
)

//go:generate go run github.com/vektra/mockery/v2@v2.40.0 --name=UrlGetter
type UrlGetter interface {
	GetUrl(models.Url) (models.Url, error)
}

//go:generate go run github.com/vektra/mockery/v2@v2.40.0 --name=UrlCreator
type UrlCreator interface {
	CreateUrl(models.Url) error
	GetMaxUrl() (models.Url, error)
}

func NewService(cfg *config.Config, creator UrlCreator, getter UrlGetter) Service {
	StringUrl += cfg.RunIp + ":" + cfg.RunPort + "/"

	return Service{creator: creator, getter: getter}
}

// Функция помещает в хранилище полученную полную ссылку и вычисленную сокращённую
func (s *Service) CreateUrl(fullUrl string) (string, error) {
	maxUrl, err := s.creator.GetMaxUrl()
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return "", err
		}
		maxUrl = models.Url{ShortUrl: ""}
	}

	url := models.Url{
		ShortUrl: shortString(maxUrl.ShortUrl),
		FullUrl:  fullUrl,
	}
	if err := s.creator.CreateUrl(url); err != nil {
		return "", err
	}

	return StringUrl + url.ShortUrl, nil
}

// Функция получает из хранилища полную ссылку по сокращённой
func (s *Service) GetUrl(shortUrl string) (string, error) {
	url, err := s.getter.GetUrl(models.Url{ShortUrl: shortUrl})
	if err != nil {
		return "", err
	}
	return url.FullUrl, nil
}

// Функция для увеличения строки maxShort на 1
func shortString(maxShort string) string {
	//С конца проходимся по максимальной строке
	for i := len(maxShort) - 1; i >= 0; i-- {
		//если текущий символ это не последний символ в строке alpabet
		if maxShort[i] != models.Alpabet[len(models.Alpabet)-1] {
			//увеличиваем первый попавшийся символ != последнему символу в строке alpabet на 1
			return maxShort[:i] + string(models.Alpabet[strings.Index(models.Alpabet, string(maxShort[i]))+1]) + maxShort[i+1:]
		}
	}
	//Если все символы в строке = последнему символу в строке alpabet, то просто выводим len(maxShort) + 1 1-х символов alphabet
	return strings.Repeat(string(models.Alpabet[0]), len(maxShort)+1)
}
