package dtos

import "io"

type SingularityRack struct {
	CurrentState *SingularityMachineStateHistoryUpdate `json:"currentState"`
	FirstSeenAt  int64                                 `json:"firstSeenAt"`
	Id           string                                `json:"id"`
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

type SingularityRackList []*SingularityRack

func (list *SingularityRackList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *SingularityRackList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityRackList) FormatJSON() string {
	return FormatJSON(list)
}
