package dtos

import "io"

type HTTP struct {
	//	AllFields *Map[FieldDescriptor,Object] `json:"allFields"`
	DefaultInstanceForType    *HTTP       `json:"defaultInstanceForType"`
	DescriptorForType         *Descriptor `json:"descriptorForType"`
	InitializationErrorString string      `json:"initializationErrorString"`
	Initialized               bool        `json:"initialized"`
	//	ParserForType *com.google.protobuf.Parser&lt;org.apache.mesos.Protos$HealthCheck$HTTP&gt; `json:"parserForType"`
	Path           string           `json:"path"`
	PathBytes      *ByteString      `json:"pathBytes"`
	Port           int32            `json:"port"`
	SerializedSize int32            `json:"serializedSize"`
	StatusesCount  int32            `json:"statusesCount"`
	StatusesList   []int32          `json:"statusesList"`
	UnknownFields  *UnknownFieldSet `json:"unknownFields"`
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

type HTTPList []*HTTP

func (list *HTTPList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *HTTPList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *HTTPList) FormatJSON() string {
	return FormatJSON(list)
}
