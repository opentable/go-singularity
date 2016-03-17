package client

import (
	"encoding/json"
	"io"
)

type SingularityState struct {
}

func (client *Client) GetState() (ss SingularityState, err error) {
	ss = SingularityState{}
	response, err := client.APIGet("state")
	if err != nil {
		return
	}
	ss.Populate(response)
	return
}

func (ss *SingularityState) Populate(jsonReader io.ReadCloser) (err error) {
	data := make([]byte, 1024)
	chunk := make([]byte, 1024)
	for {
		_, err = jsonReader.Read(chunk)
		data = append(data, chunk...)
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

func (ss *SingularityState) FormatText() string {
	return "California. Uber alles."
}
