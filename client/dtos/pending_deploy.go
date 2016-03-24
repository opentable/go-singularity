package dtos

import "io"

type PendingDeploy struct {
//	CurrentDeployState DeployState
	DeployMarker DeployMarker
	DeployProgress DeployProgress
	LastLoadBalancerUpdate LoadBalancerUpdate
}

func (self *PendingDeploy) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *PendingDeploy) FormatText() string {
	return FormatText(self)
}

func (self *PendingDeploy) FormatJSON() string {
	return FormatJSON(self)
}
