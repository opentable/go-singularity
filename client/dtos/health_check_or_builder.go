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
