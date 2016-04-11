package dtos

import "io"

type SingularityPendingDeploy struct {
	//	CurrentDeployState *DeployState
	DeployMarker           *SingularityDeployMarker
	DeployProgress         *SingularityDeployProgress
	LastLoadBalancerUpdate *SingularityLoadBalancerUpdate
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
