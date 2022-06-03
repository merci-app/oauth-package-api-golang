package main

import (
	"sync"
	"time"
)

type Authorization struct {
	username    string
	password    string
	accessToken string
	expiresIn   time.Time
	lock        sync.Mutex
}

type authResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`

	Error string `json:"error"`
}
