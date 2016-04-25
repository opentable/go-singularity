package dtos

import "io"

type SingularityWebhookWebhookType string

const (
	SingularityWebhookWebhookTypeTASK    SingularityWebhookWebhookType = "TASK"
	SingularityWebhookWebhookTypeREQUEST SingularityWebhookWebhookType = "REQUEST"
	SingularityWebhookWebhookTypeDEPLOY  SingularityWebhookWebhookType = "DEPLOY"
)

type SingularityWebhook struct {
	Id        string                        `json:"id"`
	Timestamp int64                         `json:"timestamp"`
	Type      SingularityWebhookWebhookType `json:"type"`
	Uri       string                        `json:"uri"`
	User      string                        `json:"user"`
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
