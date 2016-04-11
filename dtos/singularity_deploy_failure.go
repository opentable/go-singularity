package dtos

import "io"

type SingularityDeployFailure struct {
	Message string
	//	Reason *SingularityDeployFailureReason
	TaskId *SingularityTaskId
}

func (self *SingularityDeployFailure) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SingularityDeployFailure) FormatText() string {
	return FormatText(self)
}

func (self *SingularityDeployFailure) FormatJSON() string {
	return FormatJSON(self)
}

type SingularityDeployFailureList []*SingularityDeployFailure

func (list *SingularityDeployFailureList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *SingularityDeployFailureList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityDeployFailureList) FormatJSON() string {
	return FormatJSON(list)
}
