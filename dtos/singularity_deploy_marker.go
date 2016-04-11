package dtos

import "io"

type SingularityDeployMarker struct {
	DeployId  string
	Message   string
	RequestId string
	Timestamp int64
	User      string
}

func (self *SingularityDeployMarker) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SingularityDeployMarker) FormatText() string {
	return FormatText(self)
}

func (self *SingularityDeployMarker) FormatJSON() string {
	return FormatJSON(self)
}

type SingularityDeployMarkerList []*SingularityDeployMarker

func (list *SingularityDeployMarkerList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *SingularityDeployMarkerList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityDeployMarkerList) FormatJSON() string {
	return FormatJSON(list)
}
