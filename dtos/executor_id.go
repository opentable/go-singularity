package dtos

import "io"

type ExecutorID struct {
	//	AllFields *Map[FieldDescriptor,Object] `json:"allFields"`
	DefaultInstanceForType    *ExecutorID `json:"defaultInstanceForType"`
	DescriptorForType         *Descriptor `json:"descriptorForType"`
	InitializationErrorString string      `json:"initializationErrorString"`
	Initialized               bool        `json:"initialized"`
	//	ParserForType *com.google.protobuf.Parser&lt;org.apache.mesos.Protos$ExecutorID&gt; `json:"parserForType"`
	SerializedSize int32            `json:"serializedSize"`
	UnknownFields  *UnknownFieldSet `json:"unknownFields"`
	Value          string           `json:"value"`
	ValueBytes     *ByteString      `json:"valueBytes"`
}

func (self *ExecutorID) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *ExecutorID) FormatText() string {
	return FormatText(self)
}

func (self *ExecutorID) FormatJSON() string {
	return FormatJSON(self)
}

type ExecutorIDList []*ExecutorID

func (list *ExecutorIDList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *ExecutorIDList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *ExecutorIDList) FormatJSON() string {
	return FormatJSON(list)
}
