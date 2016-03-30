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

type SingularityLoadBalancerUpdateList []*SingularityLoadBalancerUpdate

func (list SingularityLoadBalancerUpdateList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list SingularityLoadBalancerUpdateList) FormatText() string {
	text := []byte{}
	for _, dto := range list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list SingularityLoadBalancerUpdateList) FormatJSON() string {
	return FormatJSON(list)
}
