package dtos

import "io"

type SingularityKilledTaskIdRecord struct {
	OriginalTimestamp int64 `json:"originalTimestamp"`
	//	RequestCleanupType *RequestCleanupType `json:"requestCleanupType"`
	Retries int32 `json:"retries"`
	//	TaskCleanupType *TaskCleanupType `json:"taskCleanupType"`
	TaskId    *SingularityTaskId `json:"taskId"`
	Timestamp int64              `json:"timestamp"`
}

func (self *SingularityKilledTaskIdRecord) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SingularityKilledTaskIdRecord) FormatText() string {
	return FormatText(self)
}

func (self *SingularityKilledTaskIdRecord) FormatJSON() string {
	return FormatJSON(self)
}

type SingularityKilledTaskIdRecordList []*SingularityKilledTaskIdRecord

func (list *SingularityKilledTaskIdRecordList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *SingularityKilledTaskIdRecordList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityKilledTaskIdRecordList) FormatJSON() string {
	return FormatJSON(list)
}
