package utils

import (
	"errors"
	"strings"
	"time"

	"github.com/k2sebeom/go-echo-server-template/config"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

const (
	TokenIssuer = "k2sebeom"
)

type TokenCredentials struct {
	AccessToken string;
	RefreshToken string;
}

func issueToken(id uint, expiresIn time.Duration, secret []byte) (string, error) {
	issuedAt := time.Now()
	t := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id": id,
			"issuer": TokenIssuer,
			"exp": issuedAt.Add(expiresIn).Unix(),
		},
	)
	token, err := t.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return token, nil
}

func ValidatePassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func IssueCredentials(id uint) (TokenCredentials, error) {
	c := config.GlobalConfig.Auth
	accessToken, err := issueToken(id, time.Hour * 24, []byte(c.AccessSecret));
	if err != nil {
		return TokenCredentials{}, err
	}
	refreshToken, err := issueToken(id, time.Hour * 336, []byte(c.RefreshSecret));
	if err != nil {
		return TokenCredentials{}, err
	}
	return TokenCredentials{
		AccessToken: accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func ValidateToken(tokenStr string, secret string) (uint, error) {
	tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")
	keyFunc := func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != "HS256" {
			return nil, errors.New("TOKEN INVALID")
		}
		return []byte(secret), nil
	}

	token, err := jwt.Parse(tokenStr, keyFunc)
	if err != nil {
		return 0, errors.New("TOKEN INVALID")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("TOKEN INVALID")
	}
	userId, err := verifyClaims(claims)
	if err != nil {
		return 0, errors.New("TOKEN INVALID")
	}

	return userId, nil
}

func verifyClaims(claims jwt.MapClaims) (uint, error) {
	// Check if token expired
	exp, ok := claims["exp"].(float64)
	if !ok {
		return 0, errors.New("Invalid Token")
	}
	if int64(exp) < time.Now().Unix() {
		return 0, errors.New("Token Expired")
	}

	// Check token Issuer
	issuer, ok := claims["issuer"].(string)
	if !ok {
		return 0, errors.New("Invalid Token")
	}
	if issuer != TokenIssuer {
		return 0, errors.New("Issuer Mismatch")
	}

	// Extract UserId
	userId, ok := claims["id"].(float64)
	if !ok {
		return 0, errors.New("Invalid Token")
	}

	return uint(userId), nil
}
