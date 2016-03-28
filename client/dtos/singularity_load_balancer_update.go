package dtos

import "io"

type SingularityLoadBalancerUpdate struct {
	LoadBalancerRequestId *LoadBalancerRequestId
	//	LoadBalancerState *BaragonRequestState
	Message string
	//	Method *LoadBalancerMethod
	Timestamp int64
	Uri       string
}

func (self *SingularityLoadBalancerUpdate) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SingularityLoadBalancerUpdate) FormatText() string {
	return FormatText(self)
}

func (self *SingularityLoadBalancerUpdate) FormatJSON() string {
	return FormatJSON(self)
}
