package handlers

import (
	"fmt"
	"net/http"

	"github.com/k2sebeom/go-echo-server-template/services"
	"github.com/k2sebeom/go-echo-server-template/utils"
	"github.com/labstack/echo/v4"
)

func SocialKakaoHandler(oauthService *services.OAuthService) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Get Request Query
		code := c.QueryParam("code")
		if len(code) == 0 {
			return echo.NewHTTPError(http.StatusBadRequest, "Missing code in query")
		}
		user, result, err := oauthService.HandleKakao(code)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		if result == services.SOCIAL_USER_CREATED {
			return c.Redirect(http.StatusFound, "https://google.com?hash=" + user.SocialCredential.IdentifierHash)
		} else if result == services.SOCIAL_USER_FOUND {
			tokens, err := utils.IssueCredentials(user.Id)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, "Failed to Generate Credentials")
			}
			return c.Redirect(http.StatusFound, fmt.Sprintf("https://google.com?id=%d&access_token=%s&refresh_token=%s", user.Id, tokens.AccessToken, tokens.RefreshToken))
		} else {
			return echo.NewHTTPError(http.StatusInternalServerError, "Social Handle Error")
		}
	}
}

func SocialNaverHandler(oauthService *services.OAuthService) echo.HandlerFunc {
	return func(c echo.Context) error {
		code := c.QueryParam("code")
		state := c.QueryParam("state")
		if len(code) == 0 {
			return echo.NewHTTPError(http.StatusBadRequest, "Missing code in query")
		}
		if len(state) == 0 {
			return echo.NewHTTPError(http.StatusBadRequest, "Missing state in query")
		}
		user, result, err := oauthService.HandleNaver(code, state)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		if result == services.SOCIAL_USER_CREATED {
			return c.Redirect(http.StatusFound, "https://google.com?hash=" + user.SocialCredential.IdentifierHash)
		} else if result == services.SOCIAL_USER_FOUND {
			tokens, err := utils.IssueCredentials(user.Id)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, "Failed to Generate Credentials")
			}
			return c.Redirect(http.StatusFound, fmt.Sprintf("https://google.com?id=%d&access_token=%s&refresh_token=%s", user.Id, tokens.AccessToken, tokens.RefreshToken))
		} else {
			return echo.NewHTTPError(http.StatusInternalServerError, "Social Handle Error")
		}
	}
}

func SocialGoogleHandler(oauthService *services.OAuthService) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.Redirect(http.StatusFound, "https://google.com")
	}
}

