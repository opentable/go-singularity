package dtos

import "io"

type ContainerInfoType string

const (
	ContainerInfoTypeDOCKER ContainerInfoType = "DOCKER"
	ContainerInfoTypeMESOS  ContainerInfoType = "MESOS"
)

type ContainerInfo struct {
	//	AllFields *Map[FieldDescriptor,Object] `json:"allFields"`
	DefaultInstanceForType    *ContainerInfo       `json:"defaultInstanceForType"`
	DescriptorForType         *Descriptor          `json:"descriptorForType"`
	Docker                    *DockerInfo          `json:"docker"`
	DockerOrBuilder           *DockerInfoOrBuilder `json:"dockerOrBuilder"`
	Hostname                  string               `json:"hostname"`
	HostnameBytes             *ByteString          `json:"hostnameBytes"`
	InitializationErrorString string               `json:"initializationErrorString"`
	Initialized               bool                 `json:"initialized"`
	//	ParserForType *com.google.protobuf.Parser&lt;org.apache.mesos.Protos$ContainerInfo&gt; `json:"parserForType"`
	SerializedSize int32             `json:"serializedSize"`
	Type           ContainerInfoType `json:"type"`
	UnknownFields  *UnknownFieldSet  `json:"unknownFields"`
	VolumesCount   int32             `json:"volumesCount"`
	//	VolumesList *List[Volume] `json:"volumesList"`
	//	VolumesOrBuilderList *List[? extends org.apache.mesos.Protos$VolumeOrBuilder] `json:"volumesOrBuilderList"`

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

type ContainerInfoList []*ContainerInfo

func (list *ContainerInfoList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *ContainerInfoList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *ContainerInfoList) FormatJSON() string {
	return FormatJSON(list)
}
