package dtos

import "io"

type SingularityExpiringSkipHealthchecks struct {
	ActionId string `json:"actionId"`
	//	ExpiringAPIRequestObject *T `json:"expiringAPIRequestObject"`
	RequestId                string `json:"requestId"`
	RevertToSkipHealthchecks bool   `json:"revertToSkipHealthchecks"`
	StartMillis              int64  `json:"startMillis"`
	User                     string `json:"user"`
}

func (self *SingularityExpiringSkipHealthchecks) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SingularityExpiringSkipHealthchecks) FormatText() string {
	return FormatText(self)
}

func (self *SingularityExpiringSkipHealthchecks) FormatJSON() string {
	return FormatJSON(self)
}

type SingularityExpiringSkipHealthchecksList []*SingularityExpiringSkipHealthchecks

func (list *SingularityExpiringSkipHealthchecksList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *SingularityExpiringSkipHealthchecksList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityExpiringSkipHealthchecksList) FormatJSON() string {
	return FormatJSON(list)
}
