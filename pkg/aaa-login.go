package ucsm

import (
	"crypto/tls"
	"fmt"
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
		HTTPClient: httpClient,
		Debug:      false,
	}

	client, err = api.NewClient(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create API client: %w", err)
	}

	if _, err = client.AaaLogin(); err != nil { //nolint:errcheck
		return nil, fmt.Errorf("failed to login: %w", err)
	}
	return client, nil
}
