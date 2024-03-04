package models

import (
	"errors"

	"gorm.io/gorm"
)

type Url struct {
	gorm.Model
	FullUrl  string
	ShortUrl string `gorm: "index; unique"`
}

var (
	ErrUrlNotFound = errors.New("Url not found")
)

const (
	Alpabet = "0123456789aAbBcCdDeEfFgGhHiIjJkKlLmMnNoOpPqQrRsStTuUvVwWxXyYzZ"
)
