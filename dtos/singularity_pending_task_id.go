package dtos

import "io"

type SingularityPendingTaskIdPendingType string

const (
	SingularityPendingTaskIdPendingTypeIMMEDIATE                   SingularityPendingTaskIdPendingType = "IMMEDIATE"
	SingularityPendingTaskIdPendingTypeONEOFF                      SingularityPendingTaskIdPendingType = "ONEOFF"
	SingularityPendingTaskIdPendingTypeBOUNCE                      SingularityPendingTaskIdPendingType = "BOUNCE"
	SingularityPendingTaskIdPendingTypeNEW_DEPLOY                  SingularityPendingTaskIdPendingType = "NEW_DEPLOY"
	SingularityPendingTaskIdPendingTypeNEXT_DEPLOY_STEP            SingularityPendingTaskIdPendingType = "NEXT_DEPLOY_STEP"
	SingularityPendingTaskIdPendingTypeUNPAUSED                    SingularityPendingTaskIdPendingType = "UNPAUSED"
	SingularityPendingTaskIdPendingTypeRETRY                       SingularityPendingTaskIdPendingType = "RETRY"
	SingularityPendingTaskIdPendingTypeUPDATED_REQUEST             SingularityPendingTaskIdPendingType = "UPDATED_REQUEST"
	SingularityPendingTaskIdPendingTypeDECOMISSIONED_SLAVE_OR_RACK SingularityPendingTaskIdPendingType = "DECOMISSIONED_SLAVE_OR_RACK"
	SingularityPendingTaskIdPendingTypeTASK_DONE                   SingularityPendingTaskIdPendingType = "TASK_DONE"
	SingularityPendingTaskIdPendingTypeSTARTUP                     SingularityPendingTaskIdPendingType = "STARTUP"
	SingularityPendingTaskIdPendingTypeCANCEL_BOUNCE               SingularityPendingTaskIdPendingType = "CANCEL_BOUNCE"
	SingularityPendingTaskIdPendingTypeTASK_BOUNCE                 SingularityPendingTaskIdPendingType = "TASK_BOUNCE"
	SingularityPendingTaskIdPendingTypeDEPLOY_CANCELLED            SingularityPendingTaskIdPendingType = "DEPLOY_CANCELLED"
)

type SingularityPendingTaskId struct {
	CreatedAt   int64                               `json:"createdAt"`
	DeployId    string                              `json:"deployId"`
	Id          string                              `json:"id"`
	InstanceNo  int32                               `json:"instanceNo"`
	NextRunAt   int64                               `json:"nextRunAt"`
	PendingType SingularityPendingTaskIdPendingType `json:"pendingType"`
	RequestId   string                              `json:"requestId"`
}

func (self *SingularityPendingTaskId) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SingularityPendingTaskId) FormatText() string {
	return FormatText(self)
}

func (self *SingularityPendingTaskId) FormatJSON() string {
	return FormatJSON(self)
}

type SingularityPendingTaskIdList []*SingularityPendingTaskId

func (list *SingularityPendingTaskIdList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *SingularityPendingTaskIdList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityPendingTaskIdList) FormatJSON() string {
	return FormatJSON(list)
}
