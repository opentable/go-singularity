package dtos

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

type DTO interface {
	Populate(io.ReadCloser) error
	FormatText() string
	FormatJSON() string
}

func ReadPopulate(jsonReader io.ReadCloser, target interface{}) (err error) {
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

	if len(data) == 0 {
		err = nil
		return
	}

	err = json.Unmarshal(data, target)
	return
}

func FormatText(dto interface{}) string {
	return fmt.Sprintf("%+v", dto)
}

func FormatJSON(dto interface{}) string {
	str, err := json.Marshal(dto)
	if err != nil {
		return "<<XXXX>>"
	} else {
		buf := bytes.Buffer{}
		json.Indent(&buf, str, "", "  ")
		return buf.String()
	}
}
