package dtos

import "io"

type SingularityUpdatePendingDeployRequest struct {
	DeployId              string `json:"deployId"`
	RequestId             string `json:"requestId"`
	TargetActiveInstances int32  `json:"targetActiveInstances"`
}

func (self *SingularityUpdatePendingDeployRequest) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SingularityUpdatePendingDeployRequest) FormatText() string {
	return FormatText(self)
}

func (self *SingularityUpdatePendingDeployRequest) FormatJSON() string {
	return FormatJSON(self)
}

type SingularityUpdatePendingDeployRequestList []*SingularityUpdatePendingDeployRequest

func (list *SingularityUpdatePendingDeployRequestList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *SingularityUpdatePendingDeployRequestList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityUpdatePendingDeployRequestList) FormatJSON() string {
	return FormatJSON(list)
}
