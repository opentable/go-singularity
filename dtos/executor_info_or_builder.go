package dtos

import "io"

type ExecutorInfoOrBuilder struct {
	Command              *CommandInfo            `json:"command"`
	CommandOrBuilder     *CommandInfoOrBuilder   `json:"commandOrBuilder"`
	Container            *ContainerInfo          `json:"container"`
	ContainerOrBuilder   *ContainerInfoOrBuilder `json:"containerOrBuilder"`
	Data                 *ByteString             `json:"data"`
	Discovery            *DiscoveryInfo          `json:"discovery"`
	DiscoveryOrBuilder   *DiscoveryInfoOrBuilder `json:"discoveryOrBuilder"`
	ExecutorId           *ExecutorID             `json:"executorId"`
	ExecutorIdOrBuilder  *ExecutorIDOrBuilder    `json:"executorIdOrBuilder"`
	FrameworkId          *FrameworkID            `json:"frameworkId"`
	FrameworkIdOrBuilder *FrameworkIDOrBuilder   `json:"frameworkIdOrBuilder"`
	Name                 string                  `json:"name"`
	NameBytes            *ByteString             `json:"nameBytes"`
	ResourcesCount       int32                   `json:"resourcesCount"`
	//	ResourcesList *List[Resource] `json:"resourcesList"`
	//	ResourcesOrBuilderList *List[? extends org.apache.mesos.Protos$ResourceOrBuilder] `json:"resourcesOrBuilderList"`
	Source      string      `json:"source"`
	SourceBytes *ByteString `json:"sourceBytes"`
}

func (self *ExecutorInfoOrBuilder) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *ExecutorInfoOrBuilder) FormatText() string {
	return FormatText(self)
}

func (self *ExecutorInfoOrBuilder) FormatJSON() string {
	return FormatJSON(self)
}

type ExecutorInfoOrBuilderList []*ExecutorInfoOrBuilder

func (list *ExecutorInfoOrBuilderList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *ExecutorInfoOrBuilderList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *ExecutorInfoOrBuilderList) FormatJSON() string {
	return FormatJSON(list)
}
