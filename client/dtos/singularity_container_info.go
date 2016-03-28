package dtos

import "io"

type SingularityContainerInfo struct {
	Docker *SingularityDockerInfo
	//	Type *SingularityContainerType
	Volumes *SingularityVolume
}

func (self *SingularityContainerInfo) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SingularityContainerInfo) FormatText() string {
	return FormatText(self)
}

func (self *SingularityContainerInfo) FormatJSON() string {
	return FormatJSON(self)
}
