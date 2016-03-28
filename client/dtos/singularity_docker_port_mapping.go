package dtos

import "io"

type SingularityDockerPortMapping struct {
	ContainerPort int32
	//	ContainerPortType *SingularityPortMappingType
	HostPort int32
	//	HostPortType *SingularityPortMappingType
	Protocol string
}

func (self *SingularityDockerPortMapping) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SingularityDockerPortMapping) FormatText() string {
	return FormatText(self)
}

func (self *SingularityDockerPortMapping) FormatJSON() string {
	return FormatJSON(self)
}
