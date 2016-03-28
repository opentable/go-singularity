package dtos

import "io"

type SingularityExpiringSkipHealthchecks struct {
	ActionId string
	//	ExpiringAPIRequestObject *T
	RequestId                string
	RevertToSkipHealthchecks bool
	StartMillis              int64
	User                     string
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
