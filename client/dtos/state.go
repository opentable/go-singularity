package dtos

import "io"

type State struct {
	ActiveSlaves              int
	ActiveRacks               int
	AuthDatastoreHealthy      bool
	OverProvisionedRequestIds []string
	HostStates                []HostState
}

type HostState struct {
	HostAddress, HostName, DriverStatus, MesosMaster string
	Uptime, MillisSinceLastOffer                     int64
	Master                                           bool
}

type APIClient interface {
	APIGet(string, DTO) error
}

func (ss *State) Get(client APIClient) (err error) {
	return client.APIGet("state", ss)
}

func (ss *State) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, ss)
	return
}

func (ss *State) FormatText() string {
	return FormatText(ss)
}

func (ss *State) FormatJSON() string {
	return FormatJSON(ss)
}
