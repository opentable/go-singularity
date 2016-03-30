package dtos

import "io"

type SingularityDeployResult struct {
	DeployFailures SingularityDeployFailureList
	//	DeployState *DeployState
	LbUpdate  *SingularityLoadBalancerUpdate
	Message   string
	Timestamp int64
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

func (list SingularityDeployResultList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list SingularityDeployResultList) FormatText() string {
	text := []byte{}
	for _, dto := range list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list SingularityDeployResultList) FormatJSON() string {
	return FormatJSON(list)
}
