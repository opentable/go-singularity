package dtos

import "io"

type FrameworkID struct {
	//	AllFields *Map[FieldDescriptor,Object]
	DefaultInstanceForType    *FrameworkID
	DescriptorForType         *Descriptor
	InitializationErrorString string
	Initialized               bool
	//	ParserForType *com.google.protobuf.Parser&lt;org.apache.mesos.Protos$FrameworkID&gt;
	SerializedSize int32
	UnknownFields  *UnknownFieldSet
	Value          string
	ValueBytes     *ByteString
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
