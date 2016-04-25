package dtos

import "io"

type SingularityTaskId struct {
	DeployId        string `json:"deployId"`
	Host            string `json:"host"`
	Id              string `json:"id"`
	InstanceNo      int32  `json:"instanceNo"`
	RackId          string `json:"rackId"`
	RequestId       string `json:"requestId"`
	SanitizedHost   string `json:"sanitizedHost"`
	SanitizedRackId string `json:"sanitizedRackId"`
	StartedAt       int64  `json:"startedAt"`
}

func (self *SingularityTaskId) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SingularityTaskId) FormatText() string {
	return FormatText(self)
}

func (self *SingularityTaskId) FormatJSON() string {
	return FormatJSON(self)
}

type SingularityTaskIdList []*SingularityTaskId

func (list *SingularityTaskIdList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *SingularityTaskIdList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityTaskIdList) FormatJSON() string {
	return FormatJSON(list)
}
