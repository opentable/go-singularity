package dtos

import "io"

type SingularityRequestParent struct {
	ActiveDeploy             *SingularityDeploy
	ExpiringBounce           *SingularityExpiringBounce
	ExpiringPause            *SingularityExpiringPause
	ExpiringScale            *SingularityExpiringScale
	ExpiringSkipHealthchecks *SingularityExpiringSkipHealthchecks
	PendingDeploy            *SingularityDeploy
	PendingDeployState       *SingularityPendingDeploy
	Request                  *SingularityRequest
	RequestDeployState       *SingularityRequestDeployState
	//	State *RequestState
}

func (self *SingularityRequestParent) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SingularityRequestParent) FormatText() string {
	return FormatText(self)
}

func (self *SingularityRequestParent) FormatJSON() string {
	return FormatJSON(self)
}
