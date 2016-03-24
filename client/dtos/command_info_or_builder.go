package dtos

import "io"

type CommandInfoOrBuilder struct {
	ArgumentsCount int32
	ArgumentsList []string
	Container ContainerInfo
	ContainerOrBuilder ContainerInfoOrBuilder
	Environment Environment
	EnvironmentOrBuilder EnvironmentOrBuilder
	Shell bool
	UrisCount int32
//	UrisList List[URI]
//	UrisOrBuilderList List[? extends org.apache.mesos.Protos$CommandInfo$URIOrBuilder]
	User string
	UserBytes ByteString
	Value string
	ValueBytes ByteString
}

func (self *CommandInfoOrBuilder) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *CommandInfoOrBuilder) FormatText() string {
	return FormatText(self)
}

func (self *CommandInfoOrBuilder) FormatJSON() string {
	return FormatJSON(self)
}
