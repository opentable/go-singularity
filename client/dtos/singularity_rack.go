package dtos

import "io"

type SingularityRack struct {
	CurrentState *SingularityMachineStateHistoryUpdate
	FirstSeenAt  int64
	Id           string
}

func (self *SingularityRack) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SingularityRack) FormatText() string {
	return FormatText(self)
}

func (self *SingularityRack) FormatJSON() string {
	return FormatJSON(self)
}
