package dtos

import "io"

type TaskID struct {
	//	AllFields *Map[FieldDescriptor,Object] `json:"allFields"`
	DefaultInstanceForType    *TaskID     `json:"defaultInstanceForType"`
	DescriptorForType         *Descriptor `json:"descriptorForType"`
	InitializationErrorString string      `json:"initializationErrorString"`
	Initialized               bool        `json:"initialized"`
	//	ParserForType *com.google.protobuf.Parser&lt;org.apache.mesos.Protos$TaskID&gt; `json:"parserForType"`
	SerializedSize int32            `json:"serializedSize"`
	UnknownFields  *UnknownFieldSet `json:"unknownFields"`
	Value          string           `json:"value"`
	ValueBytes     *ByteString      `json:"valueBytes"`
}

func (self *TaskID) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *TaskID) FormatText() string {
	return FormatText(self)
}

func (self *TaskID) FormatJSON() string {
	return FormatJSON(self)
}

type TaskIDList []*TaskID

func (list *TaskIDList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *TaskIDList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *TaskIDList) FormatJSON() string {
	return FormatJSON(list)
}
