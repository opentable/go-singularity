package dtos

import "io"

type SingularityMachineChangeRequest struct {
	Message string `json:"message"`
}

func (self *SingularityMachineChangeRequest) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SingularityMachineChangeRequest) FormatText() string {
	return FormatText(self)
}

func (self *SingularityMachineChangeRequest) FormatJSON() string {
	return FormatJSON(self)
}

type SingularityMachineChangeRequestList []*SingularityMachineChangeRequest

func (list *SingularityMachineChangeRequestList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *SingularityMachineChangeRequestList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityMachineChangeRequestList) FormatJSON() string {
	return FormatJSON(list)
}
