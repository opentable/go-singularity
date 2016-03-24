package dtos

import "io"

type RequestParent struct {
	ActiveDeploy Deploy
	ExpiringBounce ExpiringBounce
	ExpiringPause ExpiringPause
	ExpiringScale ExpiringScale
	ExpiringSkipHealthchecks ExpiringSkipHealthchecks
	PendingDeploy Deploy
	PendingDeployState PendingDeploy
	Request Request
	RequestDeployState RequestDeployState
//	State RequestState
}

func (self *RequestParent) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *RequestParent) FormatText() string {
	return FormatText(self)
}

func (self *RequestParent) FormatJSON() string {
	return FormatJSON(self)
}
