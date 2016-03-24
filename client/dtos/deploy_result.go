package dtos

import "io"

type DeployResult struct {
	DeployFailures []
//	DeployState DeployState
	LbUpdate LoadBalancerUpdate
	Message string
	Timestamp int64
}

func (self *DeployResult) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *DeployResult) FormatText() string {
	return FormatText(self)
}

func (self *DeployResult) FormatJSON() string {
	return FormatJSON(self)
}
