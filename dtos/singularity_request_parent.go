package dtos

import "io"

type SingularityRequestParentRequestState string

const (
	SingularityRequestParentRequestStateACTIVE               SingularityRequestParentRequestState = "ACTIVE"
	SingularityRequestParentRequestStateDELETED              SingularityRequestParentRequestState = "DELETED"
	SingularityRequestParentRequestStatePAUSED               SingularityRequestParentRequestState = "PAUSED"
	SingularityRequestParentRequestStateSYSTEM_COOLDOWN      SingularityRequestParentRequestState = "SYSTEM_COOLDOWN"
	SingularityRequestParentRequestStateFINISHED             SingularityRequestParentRequestState = "FINISHED"
	SingularityRequestParentRequestStateDEPLOYING_TO_UNPAUSE SingularityRequestParentRequestState = "DEPLOYING_TO_UNPAUSE"
)

type SingularityRequestParent struct {
	ActiveDeploy             *SingularityDeploy                   `json:"activeDeploy"`
	ExpiringBounce           *SingularityExpiringBounce           `json:"expiringBounce"`
	ExpiringPause            *SingularityExpiringPause            `json:"expiringPause"`
	ExpiringScale            *SingularityExpiringScale            `json:"expiringScale"`
	ExpiringSkipHealthchecks *SingularityExpiringSkipHealthchecks `json:"expiringSkipHealthchecks"`
	PendingDeploy            *SingularityDeploy                   `json:"pendingDeploy"`
	PendingDeployState       *SingularityPendingDeploy            `json:"pendingDeployState"`
	Request                  *SingularityRequest                  `json:"request"`
	RequestDeployState       *SingularityRequestDeployState       `json:"requestDeployState"`
	State                    SingularityRequestParentRequestState `json:"state"`
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
