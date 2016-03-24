package dtos

import "io"

type MachineChangeRequest struct {
	Message string
}

func (self *MachineChangeRequest) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *MachineChangeRequest) FormatText() string {
	return FormatText(self)
}

func (self *MachineChangeRequest) FormatJSON() string {
	return FormatJSON(self)
}
