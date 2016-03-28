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
