package singularity

import (
	"net/http"

	"github.com/opentable/sous/util/logging"
	"github.com/opentable/swaggering"
)

//go:generate swagger-client-maker --client-package=singularity --import-name=github.com/opentable/go-singularity api-docs/ .

// Client is the top level singularity client.
// Wraps the swaggering GenericClient
type Client struct {
	swaggering.Requester
}

// NewClient builds a new Client
func NewClient(apiBase string, loggerOpt ...logging.LogSink) (client *Client) {
	var logger logging.LogSink
	if len(loggerOpt) > 0 {
		logger = loggerOpt[0]
	} else {
		logger = logging.SilentLogSet()
	}

	return &Client{&swaggering.GenericClient{
		BaseURL: apiBase,
		Logger:  logger,
		HTTP:    http.Client{},
	}}
}

// NewDummyClient builds a client/control pair for testing
func NewDummyClient(apiBase string) (*Client, swaggering.DummyControl) {
	sc, ctrl := swaggering.NewChannelDummy()
	return &Client{&sc}, ctrl
}
