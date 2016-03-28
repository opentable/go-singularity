package dtos

import "io"

type SingularityExpiringPause struct {
	ActionId string
	//	ExpiringAPIRequestObject *T
	RequestId   string
	StartMillis int64
	User        string
}

func (self *SingularityExpiringPause) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SingularityExpiringPause) FormatText() string {
	return FormatText(self)
}

func (self *SingularityExpiringPause) FormatJSON() string {
	return FormatJSON(self)
}
