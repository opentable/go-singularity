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

type SingularityRequestParentList []*SingularityRequestParent

func (list *SingularityRequestParentList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *SingularityRequestParentList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityRequestParentList) FormatJSON() string {
	return FormatJSON(list)
}
