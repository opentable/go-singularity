package dtos

import "io"

type SingularityRequestCleanupRequestCleanupType string

const (
	SingularityRequestCleanupRequestCleanupTypeDELETING           SingularityRequestCleanupRequestCleanupType = "DELETING"
	SingularityRequestCleanupRequestCleanupTypePAUSING            SingularityRequestCleanupRequestCleanupType = "PAUSING"
	SingularityRequestCleanupRequestCleanupTypeBOUNCE             SingularityRequestCleanupRequestCleanupType = "BOUNCE"
	SingularityRequestCleanupRequestCleanupTypeINCREMENTAL_BOUNCE SingularityRequestCleanupRequestCleanupType = "INCREMENTAL_BOUNCE"
)

type SingularityRequestCleanup struct {
	ActionId         string                                      `json:"actionId"`
	CleanupType      SingularityRequestCleanupRequestCleanupType `json:"cleanupType"`
	DeployId         string                                      `json:"deployId"`
	KillTasks        bool                                        `json:"killTasks"`
	Message          string                                      `json:"message"`
	RequestId        string                                      `json:"requestId"`
	SkipHealthchecks bool                                        `json:"skipHealthchecks"`
	Timestamp        int64                                       `json:"timestamp"`
	User             string                                      `json:"user"`
}

func (self *SingularityRequestCleanup) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SingularityRequestCleanup) FormatText() string {
	return FormatText(self)
}

func (self *SingularityRequestCleanup) FormatJSON() string {
	return FormatJSON(self)
}

type SingularityRequestCleanupList []*SingularityRequestCleanup

func (list *SingularityRequestCleanupList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *SingularityRequestCleanupList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityRequestCleanupList) FormatJSON() string {
	return FormatJSON(list)
}
