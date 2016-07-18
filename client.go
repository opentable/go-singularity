package singularity

import (
	"net/http"

	"github.com/opentable/swaggering"
)

//go:generate swagger-client-maker --client-package=singularity --import-name=github.com/opentable/go-singularity api-docs/ .

// Client is the top level singularity client.
// Wraps the swaggering GenericClient
type Client struct {
	swaggering.Requester
}

// NewClient builds a new Client
func NewClient(apiBase string) (client *Client) {
	return &Client{&swaggering.GenericClient{apiBase, swaggering.NullLogger{}, http.Client{}}}
}

// NewDummyClient builds a client/control pair for testing
func NewDummyClient(apiBase string) (*Client, swaggering.DummyControl) {
	sc, ctrl := swaggering.NewChannelDummy()
	return &Client{&sc}, ctrl
}
