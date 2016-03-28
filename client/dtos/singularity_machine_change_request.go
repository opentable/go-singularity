package dtos

import "io"

type SingularityMachineChangeRequest struct {
	Message string
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
