package dtos

import "io"

type SlaveID struct {
	//	AllFields *Map[FieldDescriptor,Object] `json:"allFields"`
	DefaultInstanceForType    *SlaveID    `json:"defaultInstanceForType"`
	DescriptorForType         *Descriptor `json:"descriptorForType"`
	InitializationErrorString string      `json:"initializationErrorString"`
	Initialized               bool        `json:"initialized"`
	//	ParserForType *com.google.protobuf.Parser&lt;org.apache.mesos.Protos$SlaveID&gt; `json:"parserForType"`
	SerializedSize int32            `json:"serializedSize"`
	UnknownFields  *UnknownFieldSet `json:"unknownFields"`
	Value          string           `json:"value"`
	ValueBytes     *ByteString      `json:"valueBytes"`
}

func (self *SlaveID) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SlaveID) FormatText() string {
	return FormatText(self)
}

func (self *SlaveID) FormatJSON() string {
	return FormatJSON(self)
}

type SlaveIDList []*SlaveID

func (list *SlaveIDList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *SlaveIDList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SlaveIDList) FormatJSON() string {
	return FormatJSON(list)
}
