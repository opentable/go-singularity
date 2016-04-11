package dtos

import "io"

type TaskIDOrBuilder struct {
	Value      string
	ValueBytes *ByteString
}

func (self *TaskIDOrBuilder) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *TaskIDOrBuilder) FormatText() string {
	return FormatText(self)
}

func (self *TaskIDOrBuilder) FormatJSON() string {
	return FormatJSON(self)
}

type TaskIDOrBuilderList []*TaskIDOrBuilder

func (list *TaskIDOrBuilderList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *TaskIDOrBuilderList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *TaskIDOrBuilderList) FormatJSON() string {
	return FormatJSON(list)
}
