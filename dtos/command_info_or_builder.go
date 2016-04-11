package dtos

import "io"

type CommandInfoOrBuilder struct {
	ArgumentsCount       int32
	ArgumentsList        StringList
	Container            *ContainerInfo
	ContainerOrBuilder   *ContainerInfoOrBuilder
	Environment          *Environment
	EnvironmentOrBuilder *EnvironmentOrBuilder
	Shell                bool
	UrisCount            int32
	//	UrisList *List[URI]
	//	UrisOrBuilderList *List[? extends org.apache.mesos.Protos$CommandInfo$URIOrBuilder]
	User       string
	UserBytes  *ByteString
	Value      string
	ValueBytes *ByteString
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

type CommandInfoOrBuilderList []*CommandInfoOrBuilder

func (list *CommandInfoOrBuilderList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *CommandInfoOrBuilderList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *CommandInfoOrBuilderList) FormatJSON() string {
	return FormatJSON(list)
}
