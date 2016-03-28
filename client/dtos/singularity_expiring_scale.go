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
