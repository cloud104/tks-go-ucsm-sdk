package util

import (
	"crypto/tls"
	"net/http"

	"github.com/cloud104/tks-go-ucsm-sdk/api"
)

func AaaLogin(endPoint string, username string, password string) (client *api.Client, err error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	httpClient := &http.Client{Transport: tr}
	config := api.Config{
		Endpoint:   endPoint,
		Username:   username,
		Password:   password,
		HttpClient: httpClient,
		Debug:      false,
	}
	if client, err = api.NewClient(config); err == nil {
		_, err = client.AaaLogin()
	}
	return
}
