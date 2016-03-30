package dtos

import "io"

type SingularityContainerInfo struct {
	Docker *SingularityDockerInfo
	//	Type *SingularityContainerType
	Volumes SingularityVolumeList
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

type SingularityContainerInfoList []*SingularityContainerInfo

func (list SingularityContainerInfoList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list SingularityContainerInfoList) FormatText() string {
	text := []byte{}
	for _, dto := range list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list SingularityContainerInfoList) FormatJSON() string {
	return FormatJSON(list)
}
