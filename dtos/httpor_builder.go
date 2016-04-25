package dtos

import "io"

type HTTPOrBuilder struct {
	Path          string      `json:"path"`
	PathBytes     *ByteString `json:"pathBytes"`
	Port          int32       `json:"port"`
	StatusesCount int32       `json:"statusesCount"`
	StatusesList  []int32     `json:"statusesList"`
}

func (self *HTTPOrBuilder) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *HTTPOrBuilder) FormatText() string {
	return FormatText(self)
}

func (self *HTTPOrBuilder) FormatJSON() string {
	return FormatJSON(self)
}

type HTTPOrBuilderList []*HTTPOrBuilder

func (list *HTTPOrBuilderList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *HTTPOrBuilderList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *HTTPOrBuilderList) FormatJSON() string {
	return FormatJSON(list)
}
