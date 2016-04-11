package dtos

import "io"

type SingularityTaskRequest struct {
	Deploy      *SingularityDeploy
	PendingTask *SingularityPendingTask
	Request     *SingularityRequest
}

func (self *SingularityTaskRequest) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SingularityTaskRequest) FormatText() string {
	return FormatText(self)
}

func (self *SingularityTaskRequest) FormatJSON() string {
	return FormatJSON(self)
}

type SingularityTaskRequestList []*SingularityTaskRequest

func (list *SingularityTaskRequestList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *SingularityTaskRequestList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityTaskRequestList) FormatJSON() string {
	return FormatJSON(list)
}
