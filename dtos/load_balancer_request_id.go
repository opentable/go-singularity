package dtos

import "io"

type LoadBalancerRequestIdLoadBalancerRequestType string

const (
	LoadBalancerRequestIdLoadBalancerRequestTypeADD    LoadBalancerRequestIdLoadBalancerRequestType = "ADD"
	LoadBalancerRequestIdLoadBalancerRequestTypeREMOVE LoadBalancerRequestIdLoadBalancerRequestType = "REMOVE"
	LoadBalancerRequestIdLoadBalancerRequestTypeDEPLOY LoadBalancerRequestIdLoadBalancerRequestType = "DEPLOY"
	LoadBalancerRequestIdLoadBalancerRequestTypeDELETE LoadBalancerRequestIdLoadBalancerRequestType = "DELETE"
)

type LoadBalancerRequestId struct {
	AttemptNumber int32                                        `json:"attemptNumber"`
	Id            string                                       `json:"id"`
	RequestType   LoadBalancerRequestIdLoadBalancerRequestType `json:"requestType"`
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

type LoadBalancerRequestIdList []*LoadBalancerRequestId

func (list *LoadBalancerRequestIdList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *LoadBalancerRequestIdList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *LoadBalancerRequestIdList) FormatJSON() string {
	return FormatJSON(list)
}
