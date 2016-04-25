package dtos

import "io"

type SingularityShellCommand struct {
	LogfileName string     `json:"logfileName"`
	Name        string     `json:"name"`
	Options     StringList `json:"options"`
	User        string     `json:"user"`
}

func (self *SingularityShellCommand) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SingularityShellCommand) FormatText() string {
	return FormatText(self)
}

func (self *SingularityShellCommand) FormatJSON() string {
	return FormatJSON(self)
}

type SingularityShellCommandList []*SingularityShellCommand

func (list *SingularityShellCommandList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *SingularityShellCommandList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityShellCommandList) FormatJSON() string {
	return FormatJSON(list)
}
