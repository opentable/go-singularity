package dtos

import "io"

type SingularityExpiringScale struct {
	ActionId string `json:"actionId"`
	//	ExpiringAPIRequestObject *T `json:"expiringAPIRequestObject"`
	RequestId         string `json:"requestId"`
	RevertToInstances int32  `json:"revertToInstances"`
	StartMillis       int64  `json:"startMillis"`
	User              string `json:"user"`
}

func (self *SingularityExpiringScale) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SingularityExpiringScale) FormatText() string {
	return FormatText(self)
}

func (self *SingularityExpiringScale) FormatJSON() string {
	return FormatJSON(self)
}

type SingularityExpiringScaleList []*SingularityExpiringScale

func (list *SingularityExpiringScaleList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *SingularityExpiringScaleList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityExpiringScaleList) FormatJSON() string {
	return FormatJSON(list)
}
