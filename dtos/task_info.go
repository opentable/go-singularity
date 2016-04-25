package dtos

import "io"

type TaskInfo struct {
	//	AllFields *Map[FieldDescriptor,Object] `json:"allFields"`
	Command                   *CommandInfo            `json:"command"`
	CommandOrBuilder          *CommandInfoOrBuilder   `json:"commandOrBuilder"`
	Container                 *ContainerInfo          `json:"container"`
	ContainerOrBuilder        *ContainerInfoOrBuilder `json:"containerOrBuilder"`
	Data                      *ByteString             `json:"data"`
	DefaultInstanceForType    *TaskInfo               `json:"defaultInstanceForType"`
	DescriptorForType         *Descriptor             `json:"descriptorForType"`
	Discovery                 *DiscoveryInfo          `json:"discovery"`
	DiscoveryOrBuilder        *DiscoveryInfoOrBuilder `json:"discoveryOrBuilder"`
	Executor                  *ExecutorInfo           `json:"executor"`
	ExecutorOrBuilder         *ExecutorInfoOrBuilder  `json:"executorOrBuilder"`
	HealthCheck               *HealthCheck            `json:"healthCheck"`
	HealthCheckOrBuilder      *HealthCheckOrBuilder   `json:"healthCheckOrBuilder"`
	InitializationErrorString string                  `json:"initializationErrorString"`
	Initialized               bool                    `json:"initialized"`
	Labels                    *Labels                 `json:"labels"`
	LabelsOrBuilder           *LabelsOrBuilder        `json:"labelsOrBuilder"`
	Name                      string                  `json:"name"`
	NameBytes                 *ByteString             `json:"nameBytes"`
	//	ParserForType *com.google.protobuf.Parser&lt;org.apache.mesos.Protos$TaskInfo&gt; `json:"parserForType"`
	ResourcesCount int32 `json:"resourcesCount"`
	//	ResourcesList *List[Resource] `json:"resourcesList"`
	//	ResourcesOrBuilderList *List[? extends org.apache.mesos.Protos$ResourceOrBuilder] `json:"resourcesOrBuilderList"`
	SerializedSize   int32             `json:"serializedSize"`
	SlaveId          *SlaveID          `json:"slaveId"`
	SlaveIdOrBuilder *SlaveIDOrBuilder `json:"slaveIdOrBuilder"`
	TaskId           *TaskID           `json:"taskId"`
	TaskIdOrBuilder  *TaskIDOrBuilder  `json:"taskIdOrBuilder"`
	UnknownFields    *UnknownFieldSet  `json:"unknownFields"`
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
