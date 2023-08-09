package config

import (
	"os"
)


type AuthConfig struct {
	AccessSecret string
	RefreshSecret string

	KakaoClientId string

	NaverClientId string
	NaverClientSecret string
}

func NewAuthConfig() AuthConfig {
	return AuthConfig{
		AccessSecret: os.Getenv("ACCESS_SECRET"),
		RefreshSecret: os.Getenv("REFRESH_SECRET"),
		KakaoClientId: os.Getenv("KAKAO_CLIENT_ID"),

		NaverClientId: os.Getenv("NAVER_CLIENT_ID"),
		NaverClientSecret: os.Getenv("NAVER_CLIENT_SECRET"),
	}
}