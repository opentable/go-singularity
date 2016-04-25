package dtos

import "io"

type ExecutorInfo struct {
	//	AllFields *Map[FieldDescriptor,Object] `json:"allFields"`
	Command                   *CommandInfo            `json:"command"`
	CommandOrBuilder          *CommandInfoOrBuilder   `json:"commandOrBuilder"`
	Container                 *ContainerInfo          `json:"container"`
	ContainerOrBuilder        *ContainerInfoOrBuilder `json:"containerOrBuilder"`
	Data                      *ByteString             `json:"data"`
	DefaultInstanceForType    *ExecutorInfo           `json:"defaultInstanceForType"`
	DescriptorForType         *Descriptor             `json:"descriptorForType"`
	Discovery                 *DiscoveryInfo          `json:"discovery"`
	DiscoveryOrBuilder        *DiscoveryInfoOrBuilder `json:"discoveryOrBuilder"`
	ExecutorId                *ExecutorID             `json:"executorId"`
	ExecutorIdOrBuilder       *ExecutorIDOrBuilder    `json:"executorIdOrBuilder"`
	FrameworkId               *FrameworkID            `json:"frameworkId"`
	FrameworkIdOrBuilder      *FrameworkIDOrBuilder   `json:"frameworkIdOrBuilder"`
	InitializationErrorString string                  `json:"initializationErrorString"`
	Initialized               bool                    `json:"initialized"`
	Name                      string                  `json:"name"`
	NameBytes                 *ByteString             `json:"nameBytes"`
	//	ParserForType *com.google.protobuf.Parser&lt;org.apache.mesos.Protos$ExecutorInfo&gt; `json:"parserForType"`
	ResourcesCount int32 `json:"resourcesCount"`
	//	ResourcesList *List[Resource] `json:"resourcesList"`
	//	ResourcesOrBuilderList *List[? extends org.apache.mesos.Protos$ResourceOrBuilder] `json:"resourcesOrBuilderList"`
	SerializedSize int32            `json:"serializedSize"`
	Source         string           `json:"source"`
	SourceBytes    *ByteString      `json:"sourceBytes"`
	UnknownFields  *UnknownFieldSet `json:"unknownFields"`
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

type ExecutorInfoList []*ExecutorInfo

func (list *ExecutorInfoList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *ExecutorInfoList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *ExecutorInfoList) FormatJSON() string {
	return FormatJSON(list)
}
