package dtos

import "io"

type Resources struct {
	Cpus float64
	MemoryMb float64
	NumPorts int32
}

func (self *Resources) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *Resources) FormatText() string {
	return FormatText(self)
}

func (self *Resources) FormatJSON() string {
	return FormatJSON(self)
}
