package dtos

import "io"

type CommandInfo struct {
	//	AllFields *Map[FieldDescriptor,Object] `json:"allFields"`
	ArgumentsCount            int32                   `json:"argumentsCount"`
	ArgumentsList             StringList              `json:"argumentsList"`
	Container                 *ContainerInfo          `json:"container"`
	ContainerOrBuilder        *ContainerInfoOrBuilder `json:"containerOrBuilder"`
	DefaultInstanceForType    *CommandInfo            `json:"defaultInstanceForType"`
	DescriptorForType         *Descriptor             `json:"descriptorForType"`
	Environment               *Environment            `json:"environment"`
	EnvironmentOrBuilder      *EnvironmentOrBuilder   `json:"environmentOrBuilder"`
	InitializationErrorString string                  `json:"initializationErrorString"`
	Initialized               bool                    `json:"initialized"`
	//	ParserForType *com.google.protobuf.Parser&lt;org.apache.mesos.Protos$CommandInfo&gt; `json:"parserForType"`
	SerializedSize int32            `json:"serializedSize"`
	Shell          bool             `json:"shell"`
	UnknownFields  *UnknownFieldSet `json:"unknownFields"`
	UrisCount      int32            `json:"urisCount"`
	//	UrisList *List[URI] `json:"urisList"`
	//	UrisOrBuilderList *List[? extends org.apache.mesos.Protos$CommandInfo$URIOrBuilder] `json:"urisOrBuilderList"`
	User       string      `json:"user"`
	UserBytes  *ByteString `json:"userBytes"`
	Value      string      `json:"value"`
	ValueBytes *ByteString `json:"valueBytes"`
}

func (self *CommandInfo) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *CommandInfo) FormatText() string {
	return FormatText(self)
}

func (self *CommandInfo) FormatJSON() string {
	return FormatJSON(self)
}

type CommandInfoList []*CommandInfo

func (list *CommandInfoList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *CommandInfoList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *CommandInfoList) FormatJSON() string {
	return FormatJSON(list)
}
