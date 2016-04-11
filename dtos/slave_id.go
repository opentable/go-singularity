package dtos

import "io"

type SlaveID struct {
	//	AllFields *Map[FieldDescriptor,Object]
	DefaultInstanceForType    *SlaveID
	DescriptorForType         *Descriptor
	InitializationErrorString string
	Initialized               bool
	//	ParserForType *com.google.protobuf.Parser&lt;org.apache.mesos.Protos$SlaveID&gt;
	SerializedSize int32
	UnknownFields  *UnknownFieldSet
	Value          string
	ValueBytes     *ByteString
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
