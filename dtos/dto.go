package dtos

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"
)

type DTO interface {
	Populate(io.ReadCloser) error
	FormatText() string
	FormatJSON() string
}

type Fielder interface {
	FieldsPresent() []string
	GetField(string) (interface{}, error)
	SetField(string, interface{}) error
	ClearField(string) error
	LoadMap(map[string]interface{}) error
}

type StringList []string

func (list StringList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list StringList) FormatText() string {
	return strings.Join(list, "\n")
}

func (list StringList) FormatJSON() string {
	return FormatJSON(list)
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

func MarshalJSON(dto Fielder) (buf []byte, err error) {
	data := make(map[string]interface{})
	for _, name := range dto.FieldsPresent() {
		data[name], _ = dto.GetField(name)
	}
	return json.Marshal(data)
}

func presenceFromMap(m map[string]bool) []string {
	presence := make([]string, 0)
	for name, present := range m {
		if present {
			presence = append(presence, name)
		}
	}
	return presence
}

func LoadMap(dto Fielder, from map[string]interface{}) (Fielder, error) {
	return dto, dto.LoadMap(from)
}

func loadMapIntoDTO(from map[string]interface{}, dto Fielder) error {
	errs := make([]string, 0)
	for name, value := range from {
		if err := dto.SetField(name, value); err != nil {
			errs = append(errs, err.Error())
		}
	}
	if len(errs) > 0 {
		return errors.New(strings.Join(errs, "\n"))
	}
	return nil
}

func FormatText(dto interface{}) string {
	return fmt.Sprintf("%+v", dto)
}

func FormatJSON(dto interface{}) string {
	str, err := json.Marshal(dto)
	if err != nil {
		return "&lt;<XXXX>>"
	} else {
		buf := bytes.Buffer{}
		json.Indent(&buf, str, "", "  ")
		return buf.String()
	}
}

// vim: set ft=go:
