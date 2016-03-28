package dtos

import "io"

type SingularityTaskHealthcheckResult struct {
	DurationMillis int64
	ErrorMessage   string
	ResponseBody   string
	StatusCode     int32
	TaskId         *SingularityTaskId
	Timestamp      int64
}

func (self *SingularityTaskHealthcheckResult) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SingularityTaskHealthcheckResult) FormatText() string {
	return FormatText(self)
}

func (self *SingularityTaskHealthcheckResult) FormatJSON() string {
	return FormatJSON(self)
}
