package dtos

import "io"

type SingularityHostState struct {
	DriverStatus         string
	HostAddress          string
	Hostname             string
	Master               bool
	MesosMaster          string
	MillisSinceLastOffer int64
	Uptime               int64
}

func (self *SingularityHostState) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SingularityHostState) FormatText() string {
	return FormatText(self)
}

func (self *SingularityHostState) FormatJSON() string {
	return FormatJSON(self)
}

type SingularityHostStateList []*SingularityHostState

func (list *SingularityHostStateList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *SingularityHostStateList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityHostStateList) FormatJSON() string {
	return FormatJSON(list)
}
