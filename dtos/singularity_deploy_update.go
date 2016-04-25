package dtos

import "io"

type SingularityDeployUpdateDeployEventType string

const (
	SingularityDeployUpdateDeployEventTypeSTARTING SingularityDeployUpdateDeployEventType = "STARTING"
	SingularityDeployUpdateDeployEventTypeFINISHED SingularityDeployUpdateDeployEventType = "FINISHED"
)

type SingularityDeployUpdate struct {
	Deploy       *SingularityDeploy                     `json:"deploy"`
	DeployMarker *SingularityDeployMarker               `json:"deployMarker"`
	DeployResult *SingularityDeployResult               `json:"deployResult"`
	EventType    SingularityDeployUpdateDeployEventType `json:"eventType"`
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

func (list *SingularityDeployUpdateList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *SingularityDeployUpdateList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityDeployUpdateList) FormatJSON() string {
	return FormatJSON(list)
}
