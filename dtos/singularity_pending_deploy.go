package dtos

import "io"

type SingularityPendingDeployDeployState string

const (
	SingularityPendingDeployDeployStateSUCCEEDED             SingularityPendingDeployDeployState = "SUCCEEDED"
	SingularityPendingDeployDeployStateFAILED_INTERNAL_STATE SingularityPendingDeployDeployState = "FAILED_INTERNAL_STATE"
	SingularityPendingDeployDeployStateCANCELING             SingularityPendingDeployDeployState = "CANCELING"
	SingularityPendingDeployDeployStateWAITING               SingularityPendingDeployDeployState = "WAITING"
	SingularityPendingDeployDeployStateOVERDUE               SingularityPendingDeployDeployState = "OVERDUE"
	SingularityPendingDeployDeployStateFAILED                SingularityPendingDeployDeployState = "FAILED"
	SingularityPendingDeployDeployStateCANCELED              SingularityPendingDeployDeployState = "CANCELED"
)

type SingularityPendingDeploy struct {
	CurrentDeployState     SingularityPendingDeployDeployState `json:"currentDeployState"`
	DeployMarker           *SingularityDeployMarker            `json:"deployMarker"`
	DeployProgress         *SingularityDeployProgress          `json:"deployProgress"`
	LastLoadBalancerUpdate *SingularityLoadBalancerUpdate      `json:"lastLoadBalancerUpdate"`
}

func (self *SingularityPendingDeploy) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SingularityPendingDeploy) FormatText() string {
	return FormatText(self)
}

func (self *SingularityPendingDeploy) FormatJSON() string {
	return FormatJSON(self)
}

type SingularityPendingDeployList []*SingularityPendingDeploy

func (list *SingularityPendingDeployList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *SingularityPendingDeployList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityPendingDeployList) FormatJSON() string {
	return FormatJSON(list)
}
