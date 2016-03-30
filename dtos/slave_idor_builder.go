package dtos

import "io"

type SlaveIDOrBuilder struct {
	Value      string
	ValueBytes *ByteString
}

func (self *SlaveIDOrBuilder) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SlaveIDOrBuilder) FormatText() string {
	return FormatText(self)
}

func (self *SlaveIDOrBuilder) FormatJSON() string {
	return FormatJSON(self)
}

type SlaveIDOrBuilderList []*SlaveIDOrBuilder

func (list SlaveIDOrBuilderList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list SlaveIDOrBuilderList) FormatText() string {
	text := []byte{}
	for _, dto := range list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list SlaveIDOrBuilderList) FormatJSON() string {
	return FormatJSON(list)
}
