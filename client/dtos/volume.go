package dtos

import "io"

type Volume struct {
	ContainerPath string
	HostPath string
//	Mode DockerVolumeMode
}

func (self *Volume) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *Volume) FormatText() string {
	return FormatText(self)
}

func (self *Volume) FormatJSON() string {
	return FormatJSON(self)
}
