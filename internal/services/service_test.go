package services

import (
	"UrlShorter/internal/config"
	"UrlShorter/internal/models"
	"UrlShorter/internal/services/mocks"
	"strconv"
	"testing"
)

func TestService_CreateUrl(t *testing.T) {
	type args struct {
		fullUrl string
	}
	tests := []struct {
		s       Service
		args    args
		want    string
		wantErr bool
	}{
		// {
		// 	args: args{
		// 		fullUrl: "http://google.com/",
		// 	},
		// 	want:    "http://localhost:8080/0",
		// 	wantErr: false,
		// },
		{
			args: args{
				fullUrl: "http://google.com/",
			},
			want:    "http://localhost:8080/AAbB",
			wantErr: false,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {

			creator := mocks.NewUrlCreator(t)
			//creator.On("GetMaxUrl").Return(models.Url{}, gorm.ErrRecordNotFound)
			//creator.On("CreateUrl", models.Url{FullUrl: "http://google.com/", ShortUrl: "0"}).Return(nil)

			creator.On("GetMaxUrl").Return(models.Url{ShortUrl: "AAbb"}, nil)
			creator.On("CreateUrl", models.Url{FullUrl: "http://google.com/", ShortUrl: "AAbB"}).Return(nil)

			tt.s = NewService(&config.Config{RunIp: "localhost", RunPort: "8080"}, creator, mocks.NewUrlGetter(t))

			got, err := tt.s.CreateUrl(tt.args.fullUrl)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.CreateUrl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Service.CreateUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_GetUrl(t *testing.T) {
	type args struct {
		shortUrl string
	}
	tests := []struct {
		name    string
		s       Service
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "base",
			args: args{
				shortUrl: "http://localhost:8080/0",
			},
			want:    "http://google.com/",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			getter := mocks.NewUrlGetter(t)
			getter.On("GetUrl", models.Url{ShortUrl: "http://localhost:8080/0"}).Return(models.Url{ShortUrl: "http://localhost:8080/0", FullUrl: "http://google.com/"}, nil)
			tt.s = NewService(&config.Config{RunIp: "localhost", RunPort: "8080"}, mocks.NewUrlCreator(t), getter)

			got, err := tt.s.GetUrl(tt.args.shortUrl)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GetUrl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Service.GetUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_shortString(t *testing.T) {
	type args struct {
		maxShort string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				maxShort: "",
			},
			want: "0",
		},
		{
			args: args{
				maxShort: "AAbb",
			},
			want: "AAbB",
		},
		{
			args: args{
				maxShort: "ZZZZZZZ",
			},
			want: "00000000",
		},
		{
			args: args{
				maxShort: "Z",
			},
			want: "00",
		},
		{
			args: args{
				maxShort: "ZZZzZZZ",
			},
			want: "ZZZZZZZ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortString(tt.args.maxShort); got != tt.want {
				t.Errorf("shortString() = %v, want %v", got, tt.want)
			}
		})
	}
}
