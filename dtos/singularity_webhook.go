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

type SingularityWebhookList []*SingularityWebhook

func (list *SingularityWebhookList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *SingularityWebhookList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityWebhookList) FormatJSON() string {
	return FormatJSON(list)
}
