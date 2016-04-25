package dtos

import "io"

type SingularityRequestHistoryRequestHistoryType string

const (
	SingularityRequestHistoryRequestHistoryTypeCREATED             SingularityRequestHistoryRequestHistoryType = "CREATED"
	SingularityRequestHistoryRequestHistoryTypeUPDATED             SingularityRequestHistoryRequestHistoryType = "UPDATED"
	SingularityRequestHistoryRequestHistoryTypeDELETED             SingularityRequestHistoryRequestHistoryType = "DELETED"
	SingularityRequestHistoryRequestHistoryTypePAUSED              SingularityRequestHistoryRequestHistoryType = "PAUSED"
	SingularityRequestHistoryRequestHistoryTypeUNPAUSED            SingularityRequestHistoryRequestHistoryType = "UNPAUSED"
	SingularityRequestHistoryRequestHistoryTypeENTERED_COOLDOWN    SingularityRequestHistoryRequestHistoryType = "ENTERED_COOLDOWN"
	SingularityRequestHistoryRequestHistoryTypeEXITED_COOLDOWN     SingularityRequestHistoryRequestHistoryType = "EXITED_COOLDOWN"
	SingularityRequestHistoryRequestHistoryTypeFINISHED            SingularityRequestHistoryRequestHistoryType = "FINISHED"
	SingularityRequestHistoryRequestHistoryTypeDEPLOYED_TO_UNPAUSE SingularityRequestHistoryRequestHistoryType = "DEPLOYED_TO_UNPAUSE"
	SingularityRequestHistoryRequestHistoryTypeBOUNCED             SingularityRequestHistoryRequestHistoryType = "BOUNCED"
	SingularityRequestHistoryRequestHistoryTypeSCALED              SingularityRequestHistoryRequestHistoryType = "SCALED"
	SingularityRequestHistoryRequestHistoryTypeSCALE_REVERTED      SingularityRequestHistoryRequestHistoryType = "SCALE_REVERTED"
)

type SingularityRequestHistory struct {
	CreatedAt int64                                       `json:"createdAt"`
	EventType SingularityRequestHistoryRequestHistoryType `json:"eventType"`
	Message   string                                      `json:"message"`
	Request   *SingularityRequest                         `json:"request"`
	User      string                                      `json:"user"`
}

func (self *SingularityRequestHistory) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SingularityRequestHistory) FormatText() string {
	return FormatText(self)
}

func (self *SingularityRequestHistory) FormatJSON() string {
	return FormatJSON(self)
}

type SingularityRequestHistoryList []*SingularityRequestHistory

func (list *SingularityRequestHistoryList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *SingularityRequestHistoryList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityRequestHistoryList) FormatJSON() string {
	return FormatJSON(list)
}
