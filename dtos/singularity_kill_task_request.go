package dtos

import "io"

type SingularityKillTaskRequest struct {
	ActionId               string `json:"actionId"`
	Message                string `json:"message"`
	Override               bool   `json:"override"`
	WaitForReplacementTask bool   `json:"waitForReplacementTask"`
}

func (self *SingularityKillTaskRequest) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SingularityKillTaskRequest) FormatText() string {
	return FormatText(self)
}

func (self *SingularityKillTaskRequest) FormatJSON() string {
	return FormatJSON(self)
}

type SingularityKillTaskRequestList []*SingularityKillTaskRequest

func (list *SingularityKillTaskRequestList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *SingularityKillTaskRequestList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityKillTaskRequestList) FormatJSON() string {
	return FormatJSON(list)
}
