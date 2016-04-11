package dtos

import "io"

type HealthCheckOrBuilder struct {
	Command             *CommandInfo
	CommandOrBuilder    *CommandInfoOrBuilder
	ConsecutiveFailures int32
	DelaySeconds        float64
	GracePeriodSeconds  float64
	Http                *HTTP
	HttpOrBuilder       *HTTPOrBuilder
	IntervalSeconds     float64
	TimeoutSeconds      float64
}

func (self *HealthCheckOrBuilder) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *HealthCheckOrBuilder) FormatText() string {
	return FormatText(self)
}

func (self *HealthCheckOrBuilder) FormatJSON() string {
	return FormatJSON(self)
}

type HealthCheckOrBuilderList []*HealthCheckOrBuilder

func (list *HealthCheckOrBuilderList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *HealthCheckOrBuilderList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *HealthCheckOrBuilderList) FormatJSON() string {
	return FormatJSON(list)
}
