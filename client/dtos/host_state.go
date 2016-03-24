package dtos

import "io"

type HostState struct {
	DriverStatus string
	HostAddress string
	Hostname string
	Master bool
	MesosMaster string
	MillisSinceLastOffer int64
	Uptime int64
}

func (self *HostState) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *HostState) FormatText() string {
	return FormatText(self)
}

func (self *HostState) FormatJSON() string {
	return FormatJSON(self)
}
