package services

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/k2sebeom/go-echo-server-template/config"
	"github.com/k2sebeom/go-echo-server-template/models"
	"github.com/k2sebeom/go-echo-server-template/models/schema"
	"github.com/k2sebeom/go-echo-server-template/repositories"
)

type OAuthService struct {
	userRepo repositories.IUserRepository
}

func NewOAuthService(
	userRepo repositories.IUserRepository,
) *OAuthService {
	return &OAuthService{
		userRepo: userRepo,
	}
}

const (
	kakaoAuthUrl = "https://kauth.kakao.com"
	kakaoApiUrl = "https://kapi.kakao.com"

	naverAuthUrl = "https://nid.naver.com"
	naverApiUrl = "https://openapi.naver.com"
)

type SocialHandleResult uint

const (
	SOCIAL_USER_CREATED SocialHandleResult = 1
	SOCIAL_USER_FOUND SocialHandleResult = 2
)

type (
	KakaoAuth struct {
		AccessToken string `json:"access_token"`
	}

	KakaoAccount struct {
		Email string `json:"email"`
	}

	KakaoInfo struct {
		Id uint `json:"id"`
		KakaoAccount KakaoAccount `json:"kakao_account"`
	}

	NaverAuth struct {
		AccessToken string `json:"access_token"`
	}

	NaverAccount struct {
		Id uint `json:"id"`
		Email string `json:"email"`
	}

	NaverInfo struct {
		Response NaverAccount `json:"response"`
	}
)

func (s *OAuthService) handleSocialUser(provider schema.SocialProvider, hash string, email string) (*schema.User, SocialHandleResult, error) {
	user, err := s.userRepo.GetUserBySocialHash(hash)
	if err != nil && err.Error() != models.ErrUserNotFound {
		return nil, 0, err
	}

	if user != nil {
		return user, SOCIAL_USER_FOUND, nil
	}
	user = &schema.User {
		SocialCredential: &schema.SocialCredential{
			IdentifierHash: hash,
			Provider: provider,
		},
	}
	user, err = s.userRepo.CreateUser(user)
	if err != nil {
		return nil, 0, err
	}
	return user, SOCIAL_USER_CREATED, nil
}

func (s *OAuthService) HandleKakao(code string) (*schema.User, SocialHandleResult, error) {
	clientId := config.GlobalConfig.Auth.KakaoClientId
	resp, err := http.Get(fmt.Sprintf(
		"%s/oauth/token?grant_type=authorization_code&client_id=%s&code=%s",
		kakaoAuthUrl,
		clientId,
		code,
	))
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()
	auth := new(KakaoAuth)
	respBody, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(respBody, auth)

	if err != nil {
		return nil, 0, err
	}
	
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/v2/user/me", kakaoApiUrl), nil)
	if err != nil {
		return nil, 0, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", auth.AccessToken))
	
	client := &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()
	respBody, _ = ioutil.ReadAll(resp.Body)
	info := new(KakaoInfo)
	err = json.Unmarshal(respBody, info)
	if err != nil {
		return nil, 0, err
	}
	data := fmt.Sprintf("kakao-%d-%s", info.Id, info.KakaoAccount.Email)
	hash := fmt.Sprintf("%x", sha256.Sum256([]byte(data)))
	user, result, err := s.handleSocialUser(schema.KAKAO, hash, info.KakaoAccount.Email)
	if err != nil {
		return nil, 0, err
	}
	return user, result, nil;
}


func (s *OAuthService) HandleNaver(code string, state string) (*schema.User, SocialHandleResult, error) {
	clientId := config.GlobalConfig.Auth.NaverClientId
	clientSecret := config.GlobalConfig.Auth.NaverClientSecret
	resp, err := http.Get(fmt.Sprintf(
		"%s/oauth2.0/token?grant_type=authorization_code&client_id=%s&client_secret=%s&code=%s&state=%s",
		naverAuthUrl,
		clientId,
		clientSecret,
		code,
		state,
	))
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()
	auth := new(NaverAuth)
	respBody, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(respBody, auth)

	if err != nil {
		return nil, 0, err
	}

	resp, err = http.Get(fmt.Sprintf(
		"%s/v1/nid/me",
		naverApiUrl,
	))
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()
	respBody, _ = ioutil.ReadAll(resp.Body)
	info := new(NaverInfo)
	err = json.Unmarshal(respBody, info)
	if err != nil {
		return nil, 0, err
	}
	data := fmt.Sprintf("naver-%d-%s", info.Response.Id, info.Response.Email)
	hash := fmt.Sprintf("%x", sha256.Sum256([]byte(data)))
	user, result, err := s.handleSocialUser(schema.NAVER, hash, info.Response.Email)
	if err != nil {
		return nil, 0, err
	}
	return user, result, nil;
}
