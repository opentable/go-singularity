package dtos

import "io"

type LoadBalancerRequestId struct {
	AttemptNumber int32
	Id            string
	//	RequestType *LoadBalancerRequestType
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
