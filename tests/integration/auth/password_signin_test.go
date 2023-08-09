package auth_test

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/k2sebeom/go-echo-server-template/api/handlers"
	"github.com/k2sebeom/go-echo-server-template/models/schema"
	integration_test "github.com/k2sebeom/go-echo-server-template/tests/integration"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"golang.org/x/crypto/bcrypt"
)


type PasswordSigninSuite struct {
	integration_test.IntegrationSuite

	User schema.User
	Password string
}

func (suite *PasswordSigninSuite) SetupTest() {
	suite.Setup(integration_test.IntegrationSettings{
		EnvPath: "../../../envs",
		MigrationPath: "../../../migrations",
	})
	suite.M.Down()
	suite.M.Steps(2)

	suite.Password = "password"
	hash, err := bcrypt.GenerateFromPassword([]byte(suite.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	suite.User = schema.User{
		PasswordCredential: &schema.PasswordCredential{
			Email: "dummy@gmail.com",
			PasswordHash: string(hash),
		},
	}
	suite.NoError(suite.DB.Create(&suite.User).Error, "Should create user correctly")
}

func (suite *PasswordSigninSuite) TestSigninCorrectly() {
	payload := handlers.SignInRequest {
		Email: suite.User.PasswordCredential.Email,
		Password: suite.Password,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}
	req := httptest.NewRequest(http.MethodPost, "/api/auth/signin", bytes.NewBuffer(b))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	suite.Server.Echo.ServeHTTP(rec, req)
	
	assert.Equal(suite.T(), http.StatusOK, rec.Code)
	
	resp := &handlers.SignInResponse{}
	suite.NoError(json.Unmarshal(rec.Body.Bytes(), &resp), "Response should be in format")

	suite.Equal(suite.User.Id, resp.UserId, "User id should be correct")
}


func (suite *PasswordSigninSuite) TestSigninFail() {
	payload := handlers.SignInRequest {
		Email: suite.User.PasswordCredential.Email,
		Password: suite.Password + "Wrong",
	}
	b, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}
	req := httptest.NewRequest(http.MethodPost, "/api/auth/signin", bytes.NewBuffer(b))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	suite.Server.Echo.ServeHTTP(rec, req)
	
	assert.Equal(suite.T(), http.StatusUnauthorized, rec.Code)
}

func (suite *PasswordSigninSuite) TestSigninNotExist() {
	payload := handlers.SignInRequest {
		Email: suite.User.PasswordCredential.Email + "dummy",
		Password: suite.Password,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}
	req := httptest.NewRequest(http.MethodPost, "/api/auth/signin", bytes.NewBuffer(b))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	suite.Server.Echo.ServeHTTP(rec, req)
	
	assert.Equal(suite.T(), http.StatusNotFound, rec.Code)
}

func TestPasswordSigninSuite(t *testing.T) {
	suite.Run(t, new(PasswordSigninSuite))
}
