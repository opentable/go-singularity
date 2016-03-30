package dtos

import "io"

type CommandInfo struct {
	//	AllFields *Map[FieldDescriptor,Object]
	ArgumentsCount            int32
	ArgumentsList             StringList
	Container                 *ContainerInfo
	ContainerOrBuilder        *ContainerInfoOrBuilder
	DefaultInstanceForType    *CommandInfo
	DescriptorForType         *Descriptor
	Environment               *Environment
	EnvironmentOrBuilder      *EnvironmentOrBuilder
	InitializationErrorString string
	Initialized               bool
	//	ParserForType *com.google.protobuf.Parser&lt;org.apache.mesos.Protos$CommandInfo&gt;
	SerializedSize int32
	Shell          bool
	UnknownFields  *UnknownFieldSet
	UrisCount      int32
	//	UrisList *List[URI]
	//	UrisOrBuilderList *List[? extends org.apache.mesos.Protos$CommandInfo$URIOrBuilder]
	User       string
	UserBytes  *ByteString
	Value      string
	ValueBytes *ByteString
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

func (list CommandInfoList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list CommandInfoList) FormatText() string {
	text := []byte{}
	for _, dto := range list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list CommandInfoList) FormatJSON() string {
	return FormatJSON(list)
}
