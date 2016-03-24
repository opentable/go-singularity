package dtos

import "io"

type Ports struct {
//	AllFields Map[FieldDescriptor,Object]
	DefaultInstanceForType Ports
	DescriptorForType Descriptor
	InitializationErrorString string
	Initialized bool
//	ParserForType com.google.protobuf.Parser&lt;org.apache.mesos.Protos$Ports&gt;
	PortsCount int32
//	PortsList List[Port]
//	PortsOrBuilderList List[? extends org.apache.mesos.Protos$PortOrBuilder]
	SerializedSize int32
	UnknownFields UnknownFieldSet
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
