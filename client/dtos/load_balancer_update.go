package dtos

import "io"

type LoadBalancerUpdate struct {
	LoadBalancerRequestId LoadBalancerRequestId
//	LoadBalancerState BaragonRequestState
	Message string
//	Method LoadBalancerMethod
	Timestamp int64
	Uri string
}

func (self *LoadBalancerUpdate) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *LoadBalancerUpdate) FormatText() string {
	return FormatText(self)
}

func (self *LoadBalancerUpdate) FormatJSON() string {
	return FormatJSON(self)
}
