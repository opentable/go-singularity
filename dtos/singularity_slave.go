package dtos

import "io"

type SingularitySlave struct {
	//	Attributes *Map[string,string]
	CurrentState *SingularityMachineStateHistoryUpdate
	FirstSeenAt  int64
	Host         string
	Id           string
	RackId       string
}

func (self *SingularitySlave) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SingularitySlave) FormatText() string {
	return FormatText(self)
}

func (self *SingularitySlave) FormatJSON() string {
	return FormatJSON(self)
}
