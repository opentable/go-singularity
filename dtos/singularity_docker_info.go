package dtos

import "io"

type SingularityDockerInfo struct {
	ForcePullImage bool
	Image          string
	//	Network *SingularityDockerNetworkType
	//	Parameters *Map[string,string]
	PortMappings SingularityDockerPortMappingList
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

type SingularityDockerInfoList []*SingularityDockerInfo

func (list SingularityDockerInfoList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list SingularityDockerInfoList) FormatText() string {
	text := []byte{}
	for _, dto := range list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list SingularityDockerInfoList) FormatJSON() string {
	return FormatJSON(list)
}
