package dtos

import "io"

type SingularityTaskCleanup struct {
	ActionId string
	//	CleanupType *TaskCleanupType
	Message   string
	TaskId    *SingularityTaskId
	Timestamp int64
	User      string
}

func (self *SingularityTaskCleanup) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SingularityTaskCleanup) FormatText() string {
	return FormatText(self)
}

func (self *SingularityTaskCleanup) FormatJSON() string {
	return FormatJSON(self)
}

type SingularityTaskCleanupList []*SingularityTaskCleanup

func (list *SingularityTaskCleanupList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *SingularityTaskCleanupList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityTaskCleanupList) FormatJSON() string {
	return FormatJSON(list)
}
