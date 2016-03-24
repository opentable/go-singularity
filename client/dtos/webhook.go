package dtos

import "io"

type Webhook struct {
	Id string
	Timestamp int64
//	Type WebhookType
	Uri string
	User string
}

func (self *Webhook) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *Webhook) FormatText() string {
	return FormatText(self)
}

func (self *Webhook) FormatJSON() string {
	return FormatJSON(self)
}
