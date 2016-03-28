package dtos

import "io"

type SingularityVolume struct {
	ContainerPath string
	HostPath      string
	//	Mode *SingularityDockerVolumeMode
}

func (self *SingularityVolume) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SingularityVolume) FormatText() string {
	return FormatText(self)
}

func (self *SingularityVolume) FormatJSON() string {
	return FormatJSON(self)
}
