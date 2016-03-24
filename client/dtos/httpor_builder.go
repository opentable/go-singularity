package dtos

import "io"

type HTTPOrBuilder struct {
	Path string
	PathBytes ByteString
	Port int32
	StatusesCount int32
	StatusesList []int32
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
