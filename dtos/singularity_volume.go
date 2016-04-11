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

type SingularityVolumeList []*SingularityVolume

func (list *SingularityVolumeList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *SingularityVolumeList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityVolumeList) FormatJSON() string {
	return FormatJSON(list)
}
