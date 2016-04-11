package dtos

import "io"

type SingularityExpiringBounce struct {
	ActionId string
	DeployId string
	//	ExpiringAPIRequestObject *T
	RequestId   string
	StartMillis int64
	User        string
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
