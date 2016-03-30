package dtos

import "io"

type HealthCheck struct {
	//	AllFields *Map[FieldDescriptor,Object]
	Command                   *CommandInfo
	CommandOrBuilder          *CommandInfoOrBuilder
	ConsecutiveFailures       int32
	DefaultInstanceForType    *HealthCheck
	DelaySeconds              float64
	DescriptorForType         *Descriptor
	GracePeriodSeconds        float64
	Http                      *HTTP
	HttpOrBuilder             *HTTPOrBuilder
	InitializationErrorString string
	Initialized               bool
	IntervalSeconds           float64
	//	ParserForType *com.google.protobuf.Parser&lt;org.apache.mesos.Protos$HealthCheck&gt;
	SerializedSize int32
	TimeoutSeconds float64
	UnknownFields  *UnknownFieldSet
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

func (list HealthCheckList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list HealthCheckList) FormatText() string {
	text := []byte{}
	for _, dto := range list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list HealthCheckList) FormatJSON() string {
	return FormatJSON(list)
}
