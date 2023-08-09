package handlers

import (
	"net/http"

	"github.com/k2sebeom/go-echo-server-template/config"
	"github.com/k2sebeom/go-echo-server-template/models"
	"github.com/k2sebeom/go-echo-server-template/services"
	"github.com/k2sebeom/go-echo-server-template/utils"

	"github.com/labstack/echo/v4"
)

type (
	SignUpRequest struct {
		Email string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}

	SignUpResponse struct {
		UserId uint `json:"user_id" validate:"required"`
		AccessToken string `json:"access_token" validate:"required,jwt"`
		RefreshToken string `json:"refresh_token" validate:"required,jwt"`
	}

	SignInRequest struct {
		Email string `json:"id" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}

	SignInResponse struct {
		UserId uint `json:"user_id" validate:"required"`
		AccessToken string `json:"access_token" validate:"required,jwt"`
		RefreshToken string `json:"refresh_token" validate:"required,jwt"`
	}

	PostTokenRequest struct {
		RefreshToken string `json:"refresh_token" validate:"required,jwt"`
	}

	PostTokenResponse struct {
		UserId uint `json:"user_id" validate:"required"`
		AccessToken string `json:"access_token" validate:"required,jwt"`
		RefreshToken string `json:"refresh_token" validate:"required,jwt"`
	}
)

//  Signup godoc
//	@Summary		Sign Up
//	@Description	Sign Up using email and password
//	@Accept			json
//	@Produce		json
//	@Param			info	body		SignUpRequest	true	"SignUp Info"
//	@Success		200		{object}	SignUpResponse
//	@Failure		400		{object}	echo.HTTPError	"Invalid Request"
//	@Failure		500		{object}	echo.HTTPError	"Failed to sign up"
//	@Router			/api/auth/signup [post]
//	@Tags			Auth
func SignUpHandler(authService *services.AuthService) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Get Request Body
		req := new(SignUpRequest)
		if err := utils.Validate(c, req); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
		}

		// Create User on Table with PasswordCredentials
		user, err := authService.SignUpNewUser(req.Email, req.Password)
		if err != nil {
			if err.Error() == models.ErrUserConflict {
				return echo.NewHTTPError(http.StatusConflict, "User with email already exists")
			}
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to Create User")
		}

		tokens, err := utils.IssueCredentials(user.Id)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to Generate Credentials")
		}
		// Return accessToken and refreshToken		
		return c.JSON(http.StatusOK, SignUpResponse{
			UserId: user.Id,
			AccessToken: tokens.AccessToken,
			RefreshToken: tokens.RefreshToken,
		})
	}
}


//  Signin godoc
//	@Summary		Sign In
//	@Description	Sign In using email and password
//	@Accept			json
//	@Produce		json
//	@Param			info	body		SignInRequest	true	"SignIn Info"
//	@Success		200		{object}	SignInResponse
//	@Failure		400		{object}	echo.HTTPError	"Invalid Request"
//	@Failure		500		{object}	echo.HTTPError	"Failed to sign up"
//	@Router			/api/auth/signin [post]
//	@Tags			Auth
func SignInHandler(authService *services.AuthService) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Get Request Body
		req := new(SignInRequest)
		if err := utils.Validate(c, req); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
		}

		// Get User on Table with PasswordCredentials
		user, err := authService.SignInWithPassword(req.Email, req.Password)
		if err != nil {
			if err.Error() == models.ErrUserNotFound {
				return echo.NewHTTPError(http.StatusNotFound, "User not found")
			} else if err.Error() == models.ErrPasswordMismatch {
				return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
			}
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to fetch credentials")
		}

		tokens, err := utils.IssueCredentials(user.Id)

		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to Generate Credentials")
		}

		// Return accessToken and refreshToken		
		return c.JSON(http.StatusOK, SignInResponse{
			UserId: user.Id,
			AccessToken: tokens.AccessToken,
			RefreshToken: tokens.RefreshToken,
		})
	}
}

//  RefreshToken godoc
//	@Summary		Refresh Token
//	@Description	Issue new token using refresh token
//	@Accept			json
//	@Produce		json
//	@Param			info	body		SignInRequest	true	"SignIn Info"
//	@Success		200		{object}	SignInResponse
//	@Failure		400		{object}	echo.HTTPError	"Invalid Request"
//	@Failure		401		{object}	echo.HTTPError	"Token Invalid"
//	@Failure		500		{object}	echo.HTTPError	"Failed to refresh token"
//	@Router			/api/auth/token [post]
//	@Tags			Auth
func RefreshTokenHandler(authService *services.AuthService) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Get Request Body
		req := new(PostTokenRequest)
		if err := utils.Validate(c, req); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
		}

		// Validate Refresh Token
		userId, err := utils.ValidateToken(req.RefreshToken, config.GlobalConfig.Auth.RefreshSecret)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
		}

		// Create New Tokens
		tokens, err := utils.IssueCredentials(userId)

		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to Generate Credentials")
		}

		// Return accessToken and refreshToken		
		return c.JSON(http.StatusOK, PostTokenResponse{
			UserId: userId,
			AccessToken: tokens.AccessToken,
			RefreshToken: tokens.RefreshToken,
		})
	}
}