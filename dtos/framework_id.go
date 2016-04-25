package dtos

import "io"

type FrameworkID struct {
	//	AllFields *Map[FieldDescriptor,Object] `json:"allFields"`
	DefaultInstanceForType    *FrameworkID `json:"defaultInstanceForType"`
	DescriptorForType         *Descriptor  `json:"descriptorForType"`
	InitializationErrorString string       `json:"initializationErrorString"`
	Initialized               bool         `json:"initialized"`
	//	ParserForType *com.google.protobuf.Parser&lt;org.apache.mesos.Protos$FrameworkID&gt; `json:"parserForType"`
	SerializedSize int32            `json:"serializedSize"`
	UnknownFields  *UnknownFieldSet `json:"unknownFields"`
	Value          string           `json:"value"`
	ValueBytes     *ByteString      `json:"valueBytes"`
}

func (self *FrameworkID) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *FrameworkID) FormatText() string {
	return FormatText(self)
}

func (self *FrameworkID) FormatJSON() string {
	return FormatJSON(self)
}

type FrameworkIDList []*FrameworkID

func (list *FrameworkIDList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *FrameworkIDList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *FrameworkIDList) FormatJSON() string {
	return FormatJSON(list)
}
