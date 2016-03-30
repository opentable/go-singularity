package dtos

import "io"

type SingularityDeployUpdate struct {
	Deploy       *SingularityDeploy
	DeployMarker *SingularityDeployMarker
	DeployResult *SingularityDeployResult
	//	EventType *DeployEventType
}

func (self *SingularityDeployUpdate) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SingularityDeployUpdate) FormatText() string {
	return FormatText(self)
}

func (self *SingularityDeployUpdate) FormatJSON() string {
	return FormatJSON(self)
}

type SingularityDeployUpdateList []*SingularityDeployUpdate

func (list SingularityDeployUpdateList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list SingularityDeployUpdateList) FormatText() string {
	text := []byte{}
	for _, dto := range list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list SingularityDeployUpdateList) FormatJSON() string {
	return FormatJSON(list)
}
