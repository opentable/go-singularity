package dtos

import "io"

type HealthCheckOrBuilder struct {
	Command             *CommandInfo          `json:"command"`
	CommandOrBuilder    *CommandInfoOrBuilder `json:"commandOrBuilder"`
	ConsecutiveFailures int32                 `json:"consecutiveFailures"`
	DelaySeconds        float64               `json:"delaySeconds"`
	GracePeriodSeconds  float64               `json:"gracePeriodSeconds"`
	Http                *HTTP                 `json:"http"`
	HttpOrBuilder       *HTTPOrBuilder        `json:"httpOrBuilder"`
	IntervalSeconds     float64               `json:"intervalSeconds"`
	TimeoutSeconds      float64               `json:"timeoutSeconds"`
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
