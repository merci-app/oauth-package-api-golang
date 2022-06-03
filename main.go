package main

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/merci-app/oauth-sample-api-golang/client"
	"net/http"
	"sync"
	"time"
)

var (
	username = "username"
	password = "password"
)

func main() {
	caradhras := NewAuthorization(username, password)
	err := caradhras.Authenticate()
	if err != nil {
		panic(err.Error())
	}

	// Result is a new token with full expiration time
	fmt.Println(caradhras.GetAccessToken())
	fmt.Println(caradhras.GetExpireIn().Format(time.RFC3339))

	err = caradhras.Authenticate()
	if err != nil {
		panic(err.Error())
	}

	// Result is the same token with smaller expiration time
	fmt.Println(caradhras.GetAccessToken())
	fmt.Println(caradhras.GetExpireIn().Format(time.RFC3339))

	caradhras.ExpireAccessToken()

	// Result is no token
	fmt.Println(caradhras.GetAccessToken())
	fmt.Println(caradhras.GetExpireIn().Format(time.RFC3339))

	err = caradhras.Authenticate()
	if err != nil {
		panic(err.Error())
	}

	// Result is a new token with full expiration time
	fmt.Println(caradhras.GetAccessToken())
	fmt.Println(caradhras.GetExpireIn().Format(time.RFC3339))
}

func NewAuthorization(username, password string) *Authorization {
	return &Authorization{
		username: username,
		password: password,
	}
}

func (a *Authorization) Authenticate() error {
	a.lock.Lock()
	defer a.lock.Unlock()

	expired := time.Now().After(a.expiresIn)
	if expired {
		oauth, err := a.oauth()
		if err != nil {
			return err
		}
		a.accessToken = oauth.AccessToken
		a.expiresIn = time.Now().Add(time.Duration(oauth.ExpiresIn-10) * time.Second)
	}
	return nil
}

func (a *Authorization) oauth() (oauthResponse, error) {

	var response oauthResponse
	url := "https://auth.hml.caradhras.io/oauth2/token?grant_type=client_credentials"
	basicAuth := base64.StdEncoding.EncodeToString([]byte(a.username + ":" + a.password))

	req := client.NewClient()
	resp, _, err := req.Post(url).
		Set("Content-Type", "applicaion/x-www-form-urlencoded").
		Set("Authorization", "Basic "+basicAuth).
		Do(&response)

	if err != nil {
		return response, err
	}
	if resp.StatusCode != http.StatusOK {
		return response, errors.New(response.Error)
	}

	return response, nil
}

func (a *Authorization) ExpireAccessToken() {
	a.accessToken = ""
	a.expiresIn = time.Time{}
}

func (a *Authorization) GetExpireIn() time.Time {
	return a.expiresIn
}

func (a *Authorization) GetAccessToken() string {
	return a.accessToken
}

type Authorization struct {
	username    string
	password    string
	accessToken string
	expiresIn   time.Time
	lock        sync.Mutex
}

type oauthResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	Error       string `json:"error"`
}
