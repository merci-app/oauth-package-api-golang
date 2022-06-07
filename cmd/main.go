package main

import (
	"fmt"
	"github.com/merci-app/oauth-sample-api-golang/authorization"
)

var (
	username = "username"
	password = "password"
)

func main() {
	caradhras := authorization.NewAuthorization(username, password)
	token, err := caradhras.GetAccessToken()
	if err != nil {
		panic(err.Error())
	}

	// Return access token
	fmt.Println(token)
	fmt.Println(caradhras.IsExpired())

	token, err = caradhras.GetAccessToken()
	if err != nil {
		panic(err.Error())
	}

	// Return the same access token
	fmt.Println(token)
	fmt.Println(caradhras.IsExpired())

	// Expire the access token
	caradhras.ExpireAccessToken()
	fmt.Println(caradhras.IsExpired())

	token, err = caradhras.GetAccessToken()
	if err != nil {
		panic(err.Error())
	}

	// Return new access token
	fmt.Println(token)
	fmt.Println(caradhras.IsExpired())
}
