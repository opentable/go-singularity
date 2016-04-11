package dtos

import "io"

type SingularityTaskHistoryUpdate struct {
	StatusMessage string
	StatusReason  string
	TaskId        *SingularityTaskId
	//	TaskState *ExtendedTaskState
	Timestamp int64
}

func (self *SingularityTaskHistoryUpdate) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SingularityTaskHistoryUpdate) FormatText() string {
	return FormatText(self)
}

func (self *SingularityTaskHistoryUpdate) FormatJSON() string {
	return FormatJSON(self)
}

type SingularityTaskHistoryUpdateList []*SingularityTaskHistoryUpdate

func (list *SingularityTaskHistoryUpdateList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *SingularityTaskHistoryUpdateList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityTaskHistoryUpdateList) FormatJSON() string {
	return FormatJSON(list)
}
