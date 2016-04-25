package dtos

import "io"

type SingularityRequestDeployState struct {
	ActiveDeploy  *SingularityDeployMarker `json:"activeDeploy"`
	PendingDeploy *SingularityDeployMarker `json:"pendingDeploy"`
	RequestId     string                   `json:"requestId"`
}

func (self *SingularityRequestDeployState) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SingularityRequestDeployState) FormatText() string {
	return FormatText(self)
}

func (self *SingularityRequestDeployState) FormatJSON() string {
	return FormatJSON(self)
}

type SingularityRequestDeployStateList []*SingularityRequestDeployState

func (list *SingularityRequestDeployStateList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *SingularityRequestDeployStateList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityRequestDeployStateList) FormatJSON() string {
	return FormatJSON(list)
}
