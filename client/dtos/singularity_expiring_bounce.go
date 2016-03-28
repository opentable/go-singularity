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
