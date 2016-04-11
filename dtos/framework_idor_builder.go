package dtos

import "io"

type FrameworkIDOrBuilder struct {
	Value      string
	ValueBytes *ByteString
}

func (self *FrameworkIDOrBuilder) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *FrameworkIDOrBuilder) FormatText() string {
	return FormatText(self)
}

func (self *FrameworkIDOrBuilder) FormatJSON() string {
	return FormatJSON(self)
}

type FrameworkIDOrBuilderList []*FrameworkIDOrBuilder

func (list *FrameworkIDOrBuilderList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *FrameworkIDOrBuilderList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *FrameworkIDOrBuilderList) FormatJSON() string {
	return FormatJSON(list)
}
