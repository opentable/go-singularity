package client

import (
	"encoding/json"
	"fmt"
	"io"
)

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

func (client *Client) GetState() (ss *State, err error) {
	ss = &State{}
	err = client.APIGet("state", ss)
	return
}

func (ss *State) Populate(jsonReader io.ReadCloser) (err error) {
	data := make([]byte, 0, 1024)
	chunk := make([]byte, 1024)
	for {
		var count int
		count, err = jsonReader.Read(chunk)
		data = append(data, chunk[:count]...)
		if err == io.EOF {
			jsonReader.Close()
			break
		}
		if err != nil {
			return
		}
	}

	err = json.Unmarshal(data, ss)
	return
}

func (ss *State) FormatText() string {
	return fmt.Sprintf("%+v", ss)
}
