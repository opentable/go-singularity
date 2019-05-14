package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityWebhookSummary struct {
	QueueSize *int32              `json:"queueSize,omitempty"`
	Webhook   *SingularityWebhook `json:"webhook,omitempty"`
}

func (self *SingularityWebhookSummary) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *SingularityWebhookSummary) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityWebhookSummary); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityWebhookSummary cannot copy the values from %#v", other)
}

func (self *SingularityWebhookSummary) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityWebhookSummary) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityWebhookSummary) SetField(name string, value interface{}) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityWebhookSummary", name)

	case "queueSize", "QueueSize":
		v, ok := value.(int32)
		if ok {
			self.QueueSize = &v
			return nil
		}
		return fmt.Errorf("Field queueSize/QueueSize: value %v(%T) couldn't be cast to type int32", value, value)

	case "webhook", "Webhook":
		v, ok := value.(*SingularityWebhook)
		if ok {
			self.Webhook = v
			return nil
		}
		return fmt.Errorf("Field webhook/Webhook: value %v(%T) couldn't be cast to type *SingularityWebhook", value, value)

	}
}

func (self *SingularityWebhookSummary) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityWebhookSummary", name)

	case "queueSize", "QueueSize":
		return *self.QueueSize, nil
		return nil, fmt.Errorf("Field QueueSize no set on QueueSize %+v", self)

	case "webhook", "Webhook":
		return self.Webhook, nil
		return nil, fmt.Errorf("Field Webhook no set on Webhook %+v", self)

	}
}

func (self *SingularityWebhookSummary) ClearField(name string) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityWebhookSummary", name)

	case "queueSize", "QueueSize":
		self.QueueSize = nil

	case "webhook", "Webhook":
		self.Webhook = nil

	}

	return nil
}

func (self *SingularityWebhookSummary) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type SingularityWebhookSummaryList []*SingularityWebhookSummary

func (self *SingularityWebhookSummaryList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityWebhookSummaryList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityWebhookSummaryList cannot copy the values from %#v", other)
}

func (list *SingularityWebhookSummaryList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *SingularityWebhookSummaryList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityWebhookSummaryList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
