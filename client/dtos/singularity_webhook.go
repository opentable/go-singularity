package dtos

import "io"

type SingularityWebhook struct {
	Id        string
	Timestamp int64
	//	Type *WebhookType
	Uri  string
	User string
}

func (self *SingularityWebhook) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SingularityWebhook) FormatText() string {
	return FormatText(self)
}

func (self *SingularityWebhook) FormatJSON() string {
	return FormatJSON(self)
}
