package dtos

import "io"

type ExecutorInfo struct {
//	AllFields Map[FieldDescriptor,Object]
	Command CommandInfo
	CommandOrBuilder CommandInfoOrBuilder
	Container ContainerInfo
	ContainerOrBuilder ContainerInfoOrBuilder
	Data ByteString
	DefaultInstanceForType ExecutorInfo
	DescriptorForType Descriptor
	Discovery DiscoveryInfo
	DiscoveryOrBuilder DiscoveryInfoOrBuilder
	ExecutorId ExecutorID
	ExecutorIdOrBuilder ExecutorIDOrBuilder
	FrameworkId FrameworkID
	FrameworkIdOrBuilder FrameworkIDOrBuilder
	InitializationErrorString string
	Initialized bool
	Name string
	NameBytes ByteString
//	ParserForType com.google.protobuf.Parser&lt;org.apache.mesos.Protos$ExecutorInfo&gt;
	ResourcesCount int32
//	ResourcesList List[Resource]
//	ResourcesOrBuilderList List[? extends org.apache.mesos.Protos$ResourceOrBuilder]
	SerializedSize int32
	Source string
	SourceBytes ByteString
	UnknownFields UnknownFieldSet
}

func (self *ExecutorInfo) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *ExecutorInfo) FormatText() string {
	return FormatText(self)
}

func (self *ExecutorInfo) FormatJSON() string {
	return FormatJSON(self)
}
