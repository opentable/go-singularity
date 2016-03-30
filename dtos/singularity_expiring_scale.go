package dtos

import "io"

type SingularityExpiringScale struct {
	ActionId string
	//	ExpiringAPIRequestObject *T
	RequestId         string
	RevertToInstances int32
	StartMillis       int64
	User              string
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

func (list SingularityExpiringScaleList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list SingularityExpiringScaleList) FormatText() string {
	text := []byte{}
	for _, dto := range list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list SingularityExpiringScaleList) FormatJSON() string {
	return FormatJSON(list)
}
