package dtos

import "io"

type SingularityDeployResultDeployState string

const (
	SingularityDeployResultDeployStateSUCCEEDED             SingularityDeployResultDeployState = "SUCCEEDED"
	SingularityDeployResultDeployStateFAILED_INTERNAL_STATE SingularityDeployResultDeployState = "FAILED_INTERNAL_STATE"
	SingularityDeployResultDeployStateCANCELING             SingularityDeployResultDeployState = "CANCELING"
	SingularityDeployResultDeployStateWAITING               SingularityDeployResultDeployState = "WAITING"
	SingularityDeployResultDeployStateOVERDUE               SingularityDeployResultDeployState = "OVERDUE"
	SingularityDeployResultDeployStateFAILED                SingularityDeployResultDeployState = "FAILED"
	SingularityDeployResultDeployStateCANCELED              SingularityDeployResultDeployState = "CANCELED"
)

type SingularityDeployResult struct {
	DeployFailures SingularityDeployFailureList       `json:"deployFailures"`
	DeployState    SingularityDeployResultDeployState `json:"deployState"`
	LbUpdate       *SingularityLoadBalancerUpdate     `json:"lbUpdate"`
	Message        string                             `json:"message"`
	Timestamp      int64                              `json:"timestamp"`
}

func (self *SingularityDeployResult) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SingularityDeployResult) FormatText() string {
	return FormatText(self)
}

func (self *SingularityDeployResult) FormatJSON() string {
	return FormatJSON(self)
}

type SingularityDeployResultList []*SingularityDeployResult

func (list *SingularityDeployResultList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *SingularityDeployResultList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityDeployResultList) FormatJSON() string {
	return FormatJSON(list)
}
