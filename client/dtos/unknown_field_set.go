package dtos

import "io"

type UnknownFieldSet struct {
	DefaultInstanceForType UnknownFieldSet
	Initialized bool
//	ParserForType Parser
	SerializedSize int32
	SerializedSizeAsMessageSet int32
}

func (self *UnknownFieldSet) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *UnknownFieldSet) FormatText() string {
	return FormatText(self)
}

func (self *UnknownFieldSet) FormatJSON() string {
	return FormatJSON(self)
}
