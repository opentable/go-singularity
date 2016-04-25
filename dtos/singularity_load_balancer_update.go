package dtos

import "io"

type SingularityLoadBalancerUpdateBaragonRequestState string

const (
	SingularityLoadBalancerUpdateBaragonRequestStateUNKNOWN              SingularityLoadBalancerUpdateBaragonRequestState = "UNKNOWN"
	SingularityLoadBalancerUpdateBaragonRequestStateFAILED               SingularityLoadBalancerUpdateBaragonRequestState = "FAILED"
	SingularityLoadBalancerUpdateBaragonRequestStateWAITING              SingularityLoadBalancerUpdateBaragonRequestState = "WAITING"
	SingularityLoadBalancerUpdateBaragonRequestStateSUCCESS              SingularityLoadBalancerUpdateBaragonRequestState = "SUCCESS"
	SingularityLoadBalancerUpdateBaragonRequestStateCANCELING            SingularityLoadBalancerUpdateBaragonRequestState = "CANCELING"
	SingularityLoadBalancerUpdateBaragonRequestStateCANCELED             SingularityLoadBalancerUpdateBaragonRequestState = "CANCELED"
	SingularityLoadBalancerUpdateBaragonRequestStateINVALID_REQUEST_NOOP SingularityLoadBalancerUpdateBaragonRequestState = "INVALID_REQUEST_NOOP"
)

type SingularityLoadBalancerUpdateLoadBalancerMethod string

const (
	SingularityLoadBalancerUpdateLoadBalancerMethodPRE_ENQUEUE SingularityLoadBalancerUpdateLoadBalancerMethod = "PRE_ENQUEUE"
	SingularityLoadBalancerUpdateLoadBalancerMethodENQUEUE     SingularityLoadBalancerUpdateLoadBalancerMethod = "ENQUEUE"
	SingularityLoadBalancerUpdateLoadBalancerMethodCHECK_STATE SingularityLoadBalancerUpdateLoadBalancerMethod = "CHECK_STATE"
	SingularityLoadBalancerUpdateLoadBalancerMethodCANCEL      SingularityLoadBalancerUpdateLoadBalancerMethod = "CANCEL"
	SingularityLoadBalancerUpdateLoadBalancerMethodDELETE      SingularityLoadBalancerUpdateLoadBalancerMethod = "DELETE"
)

type SingularityLoadBalancerUpdate struct {
	LoadBalancerRequestId *LoadBalancerRequestId                           `json:"loadBalancerRequestId"`
	LoadBalancerState     SingularityLoadBalancerUpdateBaragonRequestState `json:"loadBalancerState"`
	Message               string                                           `json:"message"`
	Method                SingularityLoadBalancerUpdateLoadBalancerMethod  `json:"method"`
	Timestamp             int64                                            `json:"timestamp"`
	Uri                   string                                           `json:"uri"`
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

func (list *SingularityLoadBalancerUpdateList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *SingularityLoadBalancerUpdateList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityLoadBalancerUpdateList) FormatJSON() string {
	return FormatJSON(list)
}
