package dtos

import "io"

type LoadBalancerRequestId struct {
	AttemptNumber int32
	Id string
//	RequestType LoadBalancerRequestType
}

func (self *LoadBalancerRequestId) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *LoadBalancerRequestId) FormatText() string {
	return FormatText(self)
}

func (self *LoadBalancerRequestId) FormatJSON() string {
	return FormatJSON(self)
}
