package client

import "net/http"

type Config struct {
    APIKey     string
    APIURL     string
    HTTPClient *http.Client
}
