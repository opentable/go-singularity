package dtos

import "io"

type DockerInfoOrBuilderNetwork string

const (
	DockerInfoOrBuilderNetworkHOST   DockerInfoOrBuilderNetwork = "HOST"
	DockerInfoOrBuilderNetworkBRIDGE DockerInfoOrBuilderNetwork = "BRIDGE"
	DockerInfoOrBuilderNetworkNONE   DockerInfoOrBuilderNetwork = "NONE"
)

type DockerInfoOrBuilder struct {
	ForcePullImage  bool                       `json:"forcePullImage"`
	Image           string                     `json:"image"`
	ImageBytes      *ByteString                `json:"imageBytes"`
	Network         DockerInfoOrBuilderNetwork `json:"network"`
	ParametersCount int32                      `json:"parametersCount"`
	//	ParametersList *List[Parameter] `json:"parametersList"`
	//	ParametersOrBuilderList *List[? extends org.apache.mesos.Protos$ParameterOrBuilder] `json:"parametersOrBuilderList"`
	PortMappingsCount int32 `json:"portMappingsCount"`
	//	PortMappingsList *List[PortMapping] `json:"portMappingsList"`
	//	PortMappingsOrBuilderList *List[? extends org.apache.mesos.Protos$ContainerInfo$DockerInfo$PortMappingOrBuilder] `json:"portMappingsOrBuilderList"`
	Privileged bool `json:"privileged"`
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
