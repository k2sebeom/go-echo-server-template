package routes

import (
	"github.com/k2sebeom/go-echo-server-template/api/handlers"
	"github.com/k2sebeom/go-echo-server-template/services"

	"github.com/labstack/echo/v4"
)

func AuthRoutes(router *echo.Group, authService *services.AuthService, oauthService *services.OAuthService) *echo.Group {
	auth := router.Group("/auth")
	{
		auth.POST("/signup", handlers.SignUpHandler(authService));
		auth.POST("/signin", handlers.SignInHandler(authService));
		auth.POST("/token", handlers.RefreshTokenHandler(authService));
	}
	social := auth.Group("/social")
	{
		social.GET("/kakao", handlers.SocialKakaoHandler(oauthService))
		social.GET("/naver", handlers.SocialNaverHandler(oauthService))
		social.GET("/google", handlers.SocialGoogleHandler(oauthService))
	}
	return auth
}