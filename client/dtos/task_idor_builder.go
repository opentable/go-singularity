package dtos

import "io"

type TaskIDOrBuilder struct {
	Value string
	ValueBytes ByteString
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
