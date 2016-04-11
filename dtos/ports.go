package dtos

import "io"

type Ports struct {
	//	AllFields *Map[FieldDescriptor,Object]
	DefaultInstanceForType    *Ports
	DescriptorForType         *Descriptor
	InitializationErrorString string
	Initialized               bool
	//	ParserForType *com.google.protobuf.Parser&lt;org.apache.mesos.Protos$Ports&gt;
	PortsCount int32
	//	PortsList *List[Port]
	//	PortsOrBuilderList *List[? extends org.apache.mesos.Protos$PortOrBuilder]
	SerializedSize int32
	UnknownFields  *UnknownFieldSet
}

func (self *Ports) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *Ports) FormatText() string {
	return FormatText(self)
}

func (self *Ports) FormatJSON() string {
	return FormatJSON(self)
}

type PortsList []*Ports

func (list *PortsList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *PortsList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *PortsList) FormatJSON() string {
	return FormatJSON(list)
}
