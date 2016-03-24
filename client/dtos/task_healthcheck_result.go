package dtos

import "io"

type TaskHealthcheckResult struct {
	DurationMillis int64
	ErrorMessage string
	ResponseBody string
	StatusCode int32
	TaskId TaskId
	Timestamp int64
}

func (self *TaskHealthcheckResult) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *TaskHealthcheckResult) FormatText() string {
	return FormatText(self)
}

func (self *TaskHealthcheckResult) FormatJSON() string {
	return FormatJSON(self)
}
