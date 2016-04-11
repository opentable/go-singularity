package dtos

import "io"

type SingularityTaskShellCommandRequest struct {
	ShellCommand *SingularityShellCommand
	TaskId       *SingularityTaskId
	Timestamp    int64
	User         string
}

func (self *SingularityTaskShellCommandRequest) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SingularityTaskShellCommandRequest) FormatText() string {
	return FormatText(self)
}

func (self *SingularityTaskShellCommandRequest) FormatJSON() string {
	return FormatJSON(self)
}

type SingularityTaskShellCommandRequestList []*SingularityTaskShellCommandRequest

func (list *SingularityTaskShellCommandRequestList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *SingularityTaskShellCommandRequestList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityTaskShellCommandRequestList) FormatJSON() string {
	return FormatJSON(list)
}
