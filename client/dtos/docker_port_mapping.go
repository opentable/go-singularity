package dtos

import "io"

type DockerPortMapping struct {
	ContainerPort int32
//	ContainerPortType PortMappingType
	HostPort int32
//	HostPortType PortMappingType
	Protocol string
}

func (self *DockerPortMapping) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *DockerPortMapping) FormatText() string {
	return FormatText(self)
}

func (self *DockerPortMapping) FormatJSON() string {
	return FormatJSON(self)
}
