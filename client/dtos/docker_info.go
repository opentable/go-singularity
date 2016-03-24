package dtos

import "io"

type DockerInfo struct {
//	AllFields Map[FieldDescriptor,Object]
	DefaultInstanceForType DockerInfo
	DescriptorForType Descriptor
	ForcePullImage bool
	Image string
	ImageBytes ByteString
	InitializationErrorString string
	Initialized bool
//	Network Network
	ParametersCount int32
//	ParametersList List[Parameter]
//	ParametersOrBuilderList List[? extends org.apache.mesos.Protos$ParameterOrBuilder]
//	ParserForType com.google.protobuf.Parser&lt;org.apache.mesos.Protos$ContainerInfo$DockerInfo&gt;
	PortMappingsCount int32
//	PortMappingsList List[PortMapping]
//	PortMappingsOrBuilderList List[? extends org.apache.mesos.Protos$ContainerInfo$DockerInfo$PortMappingOrBuilder]
	Privileged bool
	SerializedSize int32
	UnknownFields UnknownFieldSet
}

func (self *DockerInfo) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *DockerInfo) FormatText() string {
	return FormatText(self)
}

func (self *DockerInfo) FormatJSON() string {
	return FormatJSON(self)
}
