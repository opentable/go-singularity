package dtos

import "io"

type TaskInfo struct {
	//	AllFields *Map[FieldDescriptor,Object]
	Command                   *CommandInfo
	CommandOrBuilder          *CommandInfoOrBuilder
	Container                 *ContainerInfo
	ContainerOrBuilder        *ContainerInfoOrBuilder
	Data                      *ByteString
	DefaultInstanceForType    *TaskInfo
	DescriptorForType         *Descriptor
	Discovery                 *DiscoveryInfo
	DiscoveryOrBuilder        *DiscoveryInfoOrBuilder
	Executor                  *ExecutorInfo
	ExecutorOrBuilder         *ExecutorInfoOrBuilder
	HealthCheck               *HealthCheck
	HealthCheckOrBuilder      *HealthCheckOrBuilder
	InitializationErrorString string
	Initialized               bool
	Labels                    *Labels
	LabelsOrBuilder           *LabelsOrBuilder
	Name                      string
	NameBytes                 *ByteString
	//	ParserForType *com.google.protobuf.Parser&lt;org.apache.mesos.Protos$TaskInfo&gt;
	ResourcesCount int32
	//	ResourcesList *List[Resource]
	//	ResourcesOrBuilderList *List[? extends org.apache.mesos.Protos$ResourceOrBuilder]
	SerializedSize   int32
	SlaveId          *SlaveID
	SlaveIdOrBuilder *SlaveIDOrBuilder
	TaskId           *TaskID
	TaskIdOrBuilder  *TaskIDOrBuilder
	UnknownFields    *UnknownFieldSet
}

func (self *TaskInfo) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *TaskInfo) FormatText() string {
	return FormatText(self)
}

func (self *TaskInfo) FormatJSON() string {
	return FormatJSON(self)
}

type TaskInfoList []*TaskInfo

func (list *TaskInfoList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *TaskInfoList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *TaskInfoList) FormatJSON() string {
	return FormatJSON(list)
}
