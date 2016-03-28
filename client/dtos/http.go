package dtos

import "io"

type HTTP struct {
	//	AllFields *Map[FieldDescriptor,Object]
	DefaultInstanceForType    *HTTP
	DescriptorForType         *Descriptor
	InitializationErrorString string
	Initialized               bool
	//	ParserForType *com.google.protobuf.Parser&lt;org.apache.mesos.Protos$HealthCheck$HTTP&gt;
	Path           string
	PathBytes      *ByteString
	Port           int32
	SerializedSize int32
	StatusesCount  int32
	StatusesList   int32
	UnknownFields  *UnknownFieldSet
}

func (self *HTTP) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *HTTP) FormatText() string {
	return FormatText(self)
}

func (self *HTTP) FormatJSON() string {
	return FormatJSON(self)
}
