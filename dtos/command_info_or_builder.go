package dtos

import "io"

type CommandInfoOrBuilder struct {
	ArgumentsCount       int32                   `json:"argumentsCount"`
	ArgumentsList        StringList              `json:"argumentsList"`
	Container            *ContainerInfo          `json:"container"`
	ContainerOrBuilder   *ContainerInfoOrBuilder `json:"containerOrBuilder"`
	Environment          *Environment            `json:"environment"`
	EnvironmentOrBuilder *EnvironmentOrBuilder   `json:"environmentOrBuilder"`
	Shell                bool                    `json:"shell"`
	UrisCount            int32                   `json:"urisCount"`
	//	UrisList *List[URI] `json:"urisList"`
	//	UrisOrBuilderList *List[? extends org.apache.mesos.Protos$CommandInfo$URIOrBuilder] `json:"urisOrBuilderList"`
	User       string      `json:"user"`
	UserBytes  *ByteString `json:"userBytes"`
	Value      string      `json:"value"`
	ValueBytes *ByteString `json:"valueBytes"`
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
