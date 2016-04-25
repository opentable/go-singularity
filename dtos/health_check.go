package dtos

import "io"

type HealthCheck struct {
	//	AllFields *Map[FieldDescriptor,Object] `json:"allFields"`
	Command                   *CommandInfo          `json:"command"`
	CommandOrBuilder          *CommandInfoOrBuilder `json:"commandOrBuilder"`
	ConsecutiveFailures       int32                 `json:"consecutiveFailures"`
	DefaultInstanceForType    *HealthCheck          `json:"defaultInstanceForType"`
	DelaySeconds              float64               `json:"delaySeconds"`
	DescriptorForType         *Descriptor           `json:"descriptorForType"`
	GracePeriodSeconds        float64               `json:"gracePeriodSeconds"`
	Http                      *HTTP                 `json:"http"`
	HttpOrBuilder             *HTTPOrBuilder        `json:"httpOrBuilder"`
	InitializationErrorString string                `json:"initializationErrorString"`
	Initialized               bool                  `json:"initialized"`
	IntervalSeconds           float64               `json:"intervalSeconds"`
	//	ParserForType *com.google.protobuf.Parser&lt;org.apache.mesos.Protos$HealthCheck&gt; `json:"parserForType"`
	SerializedSize int32            `json:"serializedSize"`
	TimeoutSeconds float64          `json:"timeoutSeconds"`
	UnknownFields  *UnknownFieldSet `json:"unknownFields"`
}

func (self *HealthCheck) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *HealthCheck) FormatText() string {
	return FormatText(self)
}

func (self *HealthCheck) FormatJSON() string {
	return FormatJSON(self)
}

type HealthCheckList []*HealthCheck

func (list *HealthCheckList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *HealthCheckList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *HealthCheckList) FormatJSON() string {
	return FormatJSON(list)
}
