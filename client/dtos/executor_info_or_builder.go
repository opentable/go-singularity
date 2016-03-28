package dtos

import "io"

type ExecutorInfoOrBuilder struct {
	Command              *CommandInfo
	CommandOrBuilder     *CommandInfoOrBuilder
	Container            *ContainerInfo
	ContainerOrBuilder   *ContainerInfoOrBuilder
	Data                 *ByteString
	Discovery            *DiscoveryInfo
	DiscoveryOrBuilder   *DiscoveryInfoOrBuilder
	ExecutorId           *ExecutorID
	ExecutorIdOrBuilder  *ExecutorIDOrBuilder
	FrameworkId          *FrameworkID
	FrameworkIdOrBuilder *FrameworkIDOrBuilder
	Name                 string
	NameBytes            *ByteString
	ResourcesCount       int32
	//	ResourcesList *List[Resource]
	//	ResourcesOrBuilderList *List[? extends org.apache.mesos.Protos$ResourceOrBuilder]
	Source      string
	SourceBytes *ByteString
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
