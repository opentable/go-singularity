package dtos

import "io"

type ExecutorIDOrBuilder struct {
	Value string
	ValueBytes ByteString
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
