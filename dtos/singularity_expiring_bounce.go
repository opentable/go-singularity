package dtos

import "io"

type SingularityExpiringBounce struct {
	ActionId string `json:"actionId"`
	DeployId string `json:"deployId"`
	//	ExpiringAPIRequestObject *T `json:"expiringAPIRequestObject"`
	RequestId   string `json:"requestId"`
	StartMillis int64  `json:"startMillis"`
	User        string `json:"user"`
}

func (self *SingularityExpiringBounce) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SingularityExpiringBounce) FormatText() string {
	return FormatText(self)
}

func (self *SingularityExpiringBounce) FormatJSON() string {
	return FormatJSON(self)
}

type SingularityExpiringBounceList []*SingularityExpiringBounce

func (list *SingularityExpiringBounceList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *SingularityExpiringBounceList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityExpiringBounceList) FormatJSON() string {
	return FormatJSON(list)
}
