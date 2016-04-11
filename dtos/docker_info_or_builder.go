package dtos

import "io"

type DockerInfoOrBuilder struct {
	ForcePullImage bool
	Image          string
	ImageBytes     *ByteString
	//	Network *Network
	ParametersCount int32
	//	ParametersList *List[Parameter]
	//	ParametersOrBuilderList *List[? extends org.apache.mesos.Protos$ParameterOrBuilder]
	PortMappingsCount int32
	//	PortMappingsList *List[PortMapping]
	//	PortMappingsOrBuilderList *List[? extends org.apache.mesos.Protos$ContainerInfo$DockerInfo$PortMappingOrBuilder]
	Privileged bool
}

func (self *DockerInfoOrBuilder) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *DockerInfoOrBuilder) FormatText() string {
	return FormatText(self)
}

func (self *DockerInfoOrBuilder) FormatJSON() string {
	return FormatJSON(self)
}

type DockerInfoOrBuilderList []*DockerInfoOrBuilder

func (list *DockerInfoOrBuilderList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *DockerInfoOrBuilderList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *DockerInfoOrBuilderList) FormatJSON() string {
	return FormatJSON(list)
}
