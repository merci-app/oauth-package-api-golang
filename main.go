package main

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/merci-app/oauth-sample-api-golang/client"
	"net/http"
	"time"
)

var (
	username = "<USERNAME>"
	password = "<PASSWORD>"
)

func main() {
	caradhras := NewAuthorization(username, password)
	err := caradhras.Authenticate()
	if err != nil {
		panic(err.Error())
	}

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

	valid := time.Now().Before(a.expiresIn)
	if !valid {
		oauth, err := a.oauth()
		if err != nil {
			return err
		}
		a.accessToken = oauth.AccessToken
		a.expiresIn = time.Now().Add(time.Duration(oauth.ExpiresIn-10) * time.Second)
	}
	return nil
}

func (a *Authorization) oauth() (authResponse, error) {

	var response authResponse
	url := "https://auth.hml.caradhras.io/oauth2/token?grant_type=client_credentials"
	basicAuth := base64.StdEncoding.EncodeToString([]byte(a.username + ":" + a.password))

	req := client.NewClient()
	resp, _, err := req.Post(url).
		Set("Content-Type", "applicaion/x-www-form-urlencoded").
		Set("Authorizaion", "Basic "+basicAuth).
		Do(&response)

	if err != nil {
		return response, err
	}
	if resp.StatusCode != http.StatusOK {
		return response, errors.New(response.Error)
	}

	return response, nil
}

func (a *Authorization) GetExpireIn() time.Time {
	return a.expiresIn
}

func (a *Authorization) GetAccessToken() string {
	return a.accessToken
}
