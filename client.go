package singularity

import (
	"net/http"

	"github.com/opentable/swaggering"
)

//go:generate swagger-client-maker --client-package=singularity --import-name=github.com/opentable/go-singularity api-docs/ .

// Client is the top level singularity client.
// Wraps the swaggering GenericClient
type Client struct {
	swaggering.GenericClient
}

// NewClient builds a new Client
func NewClient(apiBase string) (client *Client) {
	return &Client{swaggering.GenericClient{apiBase, swaggering.NullLogger{}, http.Client{}}}
}
