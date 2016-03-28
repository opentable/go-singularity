package dtos

import "io"

type SingularityDockerInfo struct {
	ForcePullImage bool
	Image          string
	//	Network *SingularityDockerNetworkType
	//	Parameters *Map[string,string]
	PortMappings *SingularityDockerPortMapping
	Privileged   bool
}

func (self *SingularityDockerInfo) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SingularityDockerInfo) FormatText() string {
	return FormatText(self)
}

func (self *SingularityDockerInfo) FormatJSON() string {
	return FormatJSON(self)
}
