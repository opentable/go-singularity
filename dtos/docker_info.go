package dtos

import "io"

type DockerInfoNetwork string

const (
	DockerInfoNetworkHOST   DockerInfoNetwork = "HOST"
	DockerInfoNetworkBRIDGE DockerInfoNetwork = "BRIDGE"
	DockerInfoNetworkNONE   DockerInfoNetwork = "NONE"
)

type DockerInfo struct {
	//	AllFields *Map[FieldDescriptor,Object] `json:"allFields"`
	DefaultInstanceForType    *DockerInfo       `json:"defaultInstanceForType"`
	DescriptorForType         *Descriptor       `json:"descriptorForType"`
	ForcePullImage            bool              `json:"forcePullImage"`
	Image                     string            `json:"image"`
	ImageBytes                *ByteString       `json:"imageBytes"`
	InitializationErrorString string            `json:"initializationErrorString"`
	Initialized               bool              `json:"initialized"`
	Network                   DockerInfoNetwork `json:"network"`
	ParametersCount           int32             `json:"parametersCount"`
	//	ParametersList *List[Parameter] `json:"parametersList"`
	//	ParametersOrBuilderList *List[? extends org.apache.mesos.Protos$ParameterOrBuilder] `json:"parametersOrBuilderList"`
	//	ParserForType *com.google.protobuf.Parser&lt;org.apache.mesos.Protos$ContainerInfo$DockerInfo&gt; `json:"parserForType"`
	PortMappingsCount int32 `json:"portMappingsCount"`
	//	PortMappingsList *List[PortMapping] `json:"portMappingsList"`
	//	PortMappingsOrBuilderList *List[? extends org.apache.mesos.Protos$ContainerInfo$DockerInfo$PortMappingOrBuilder] `json:"portMappingsOrBuilderList"`
	Privileged     bool             `json:"privileged"`
	SerializedSize int32            `json:"serializedSize"`
	UnknownFields  *UnknownFieldSet `json:"unknownFields"`
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

type DockerInfoList []*DockerInfo

func (list *DockerInfoList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *DockerInfoList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *DockerInfoList) FormatJSON() string {
	return FormatJSON(list)
}
