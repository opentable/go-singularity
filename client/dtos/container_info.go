package dtos

import "io"

type ContainerInfo struct {
	//	AllFields *Map[FieldDescriptor,Object]
	DefaultInstanceForType    *ContainerInfo
	DescriptorForType         *Descriptor
	Docker                    *DockerInfo
	DockerOrBuilder           *DockerInfoOrBuilder
	Hostname                  string
	HostnameBytes             *ByteString
	InitializationErrorString string
	Initialized               bool
	//	ParserForType *com.google.protobuf.Parser&lt;org.apache.mesos.Protos$ContainerInfo&gt;
	SerializedSize int32
	//	Type *Type
	UnknownFields *UnknownFieldSet
	VolumesCount  int32
	//	VolumesList *List[Volume]
	//	VolumesOrBuilderList *List[? extends org.apache.mesos.Protos$VolumeOrBuilder]
}

func (self *ContainerInfo) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *ContainerInfo) FormatText() string {
	return FormatText(self)
}

func (self *ContainerInfo) FormatJSON() string {
	return FormatJSON(self)
}
