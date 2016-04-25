package dtos

import "io"

type ExecutorIDOrBuilder struct {
	Value      string      `json:"value"`
	ValueBytes *ByteString `json:"valueBytes"`
}

func (self *ExecutorIDOrBuilder) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *ExecutorIDOrBuilder) FormatText() string {
	return FormatText(self)
}

func (self *ExecutorIDOrBuilder) FormatJSON() string {
	return FormatJSON(self)
}

type ExecutorIDOrBuilderList []*ExecutorIDOrBuilder

func (list *ExecutorIDOrBuilderList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *ExecutorIDOrBuilderList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *ExecutorIDOrBuilderList) FormatJSON() string {
	return FormatJSON(list)
}
