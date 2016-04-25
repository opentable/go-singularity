package dtos

import "io"

type SingularityPendingRequestPendingType string

const (
	SingularityPendingRequestPendingTypeIMMEDIATE                   SingularityPendingRequestPendingType = "IMMEDIATE"
	SingularityPendingRequestPendingTypeONEOFF                      SingularityPendingRequestPendingType = "ONEOFF"
	SingularityPendingRequestPendingTypeBOUNCE                      SingularityPendingRequestPendingType = "BOUNCE"
	SingularityPendingRequestPendingTypeNEW_DEPLOY                  SingularityPendingRequestPendingType = "NEW_DEPLOY"
	SingularityPendingRequestPendingTypeNEXT_DEPLOY_STEP            SingularityPendingRequestPendingType = "NEXT_DEPLOY_STEP"
	SingularityPendingRequestPendingTypeUNPAUSED                    SingularityPendingRequestPendingType = "UNPAUSED"
	SingularityPendingRequestPendingTypeRETRY                       SingularityPendingRequestPendingType = "RETRY"
	SingularityPendingRequestPendingTypeUPDATED_REQUEST             SingularityPendingRequestPendingType = "UPDATED_REQUEST"
	SingularityPendingRequestPendingTypeDECOMISSIONED_SLAVE_OR_RACK SingularityPendingRequestPendingType = "DECOMISSIONED_SLAVE_OR_RACK"
	SingularityPendingRequestPendingTypeTASK_DONE                   SingularityPendingRequestPendingType = "TASK_DONE"
	SingularityPendingRequestPendingTypeSTARTUP                     SingularityPendingRequestPendingType = "STARTUP"
	SingularityPendingRequestPendingTypeCANCEL_BOUNCE               SingularityPendingRequestPendingType = "CANCEL_BOUNCE"
	SingularityPendingRequestPendingTypeTASK_BOUNCE                 SingularityPendingRequestPendingType = "TASK_BOUNCE"
	SingularityPendingRequestPendingTypeDEPLOY_CANCELLED            SingularityPendingRequestPendingType = "DEPLOY_CANCELLED"
)

type SingularityPendingRequest struct {
	ActionId         string                               `json:"actionId"`
	CmdLineArgsList  StringList                           `json:"cmdLineArgsList"`
	DeployId         string                               `json:"deployId"`
	Message          string                               `json:"message"`
	PendingType      SingularityPendingRequestPendingType `json:"pendingType"`
	RequestId        string                               `json:"requestId"`
	RunId            string                               `json:"runId"`
	SkipHealthchecks bool                                 `json:"skipHealthchecks"`
	Timestamp        int64                                `json:"timestamp"`
	User             string                               `json:"user"`
}

func (self *SingularityPendingRequest) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SingularityPendingRequest) FormatText() string {
	return FormatText(self)
}

func (self *SingularityPendingRequest) FormatJSON() string {
	return FormatJSON(self)
}

type SingularityPendingRequestList []*SingularityPendingRequest

func (list *SingularityPendingRequestList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *SingularityPendingRequestList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityPendingRequestList) FormatJSON() string {
	return FormatJSON(list)
}
