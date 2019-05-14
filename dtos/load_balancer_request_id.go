package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type LoadBalancerRequestIdLoadBalancerRequestType string

const (
	LoadBalancerRequestIdLoadBalancerRequestTypeADD    LoadBalancerRequestIdLoadBalancerRequestType = "ADD"
	LoadBalancerRequestIdLoadBalancerRequestTypeREMOVE LoadBalancerRequestIdLoadBalancerRequestType = "REMOVE"
	LoadBalancerRequestIdLoadBalancerRequestTypeDEPLOY LoadBalancerRequestIdLoadBalancerRequestType = "DEPLOY"
	LoadBalancerRequestIdLoadBalancerRequestTypeDELETE LoadBalancerRequestIdLoadBalancerRequestType = "DELETE"
)

type LoadBalancerRequestId struct {
	RequestType   *LoadBalancerRequestIdLoadBalancerRequestType `json:"requestType,omitempty"`
	AttemptNumber *int32                                        `json:"attemptNumber,omitempty"`
	Id            *string                                       `json:"id,omitempty"`
}

func (self *LoadBalancerRequestId) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *LoadBalancerRequestId) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*LoadBalancerRequestId); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A LoadBalancerRequestId cannot copy the values from %#v", other)
}

func (self *LoadBalancerRequestId) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *LoadBalancerRequestId) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *LoadBalancerRequestId) SetField(name string, value interface{}) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on LoadBalancerRequestId", name)

	case "requestType", "RequestType":
		v, ok := value.(LoadBalancerRequestIdLoadBalancerRequestType)
		if ok {
			self.RequestType = &v
			return nil
		}
		return fmt.Errorf("Field requestType/RequestType: value %v(%T) couldn't be cast to type LoadBalancerRequestIdLoadBalancerRequestType", value, value)

	case "attemptNumber", "AttemptNumber":
		v, ok := value.(int32)
		if ok {
			self.AttemptNumber = &v
			return nil
		}
		return fmt.Errorf("Field attemptNumber/AttemptNumber: value %v(%T) couldn't be cast to type int32", value, value)

	case "id", "Id":
		v, ok := value.(string)
		if ok {
			self.Id = &v
			return nil
		}
		return fmt.Errorf("Field id/Id: value %v(%T) couldn't be cast to type string", value, value)

	}
}

func (self *LoadBalancerRequestId) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on LoadBalancerRequestId", name)

	case "requestType", "RequestType":
		return *self.RequestType, nil
		return nil, fmt.Errorf("Field RequestType no set on RequestType %+v", self)

	case "attemptNumber", "AttemptNumber":
		return *self.AttemptNumber, nil
		return nil, fmt.Errorf("Field AttemptNumber no set on AttemptNumber %+v", self)

	case "id", "Id":
		return *self.Id, nil
		return nil, fmt.Errorf("Field Id no set on Id %+v", self)

	}
}

func (self *LoadBalancerRequestId) ClearField(name string) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on LoadBalancerRequestId", name)

	case "requestType", "RequestType":
		self.RequestType = nil

	case "attemptNumber", "AttemptNumber":
		self.AttemptNumber = nil

	case "id", "Id":
		self.Id = nil

	}

	return nil
}

func (self *LoadBalancerRequestId) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type LoadBalancerRequestIdList []*LoadBalancerRequestId

func (self *LoadBalancerRequestIdList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*LoadBalancerRequestIdList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A LoadBalancerRequestIdList cannot copy the values from %#v", other)
}

func (list *LoadBalancerRequestIdList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
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
	return swaggering.FormatJSON(list)
}
