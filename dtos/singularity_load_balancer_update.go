package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

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
	LoadBalancerState     *SingularityLoadBalancerUpdateBaragonRequestState `json:"loadBalancerState,omitempty"`
	Message               *string                                           `json:"message,omitempty"`
	Timestamp             *int64                                            `json:"timestamp,omitempty"`
	Uri                   *string                                           `json:"uri,omitempty"`
	Method                *SingularityLoadBalancerUpdateLoadBalancerMethod  `json:"method,omitempty"`
	LoadBalancerRequestId *LoadBalancerRequestId                            `json:"loadBalancerRequestId,omitempty"`
}

func (self *SingularityLoadBalancerUpdate) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *SingularityLoadBalancerUpdate) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityLoadBalancerUpdate); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityLoadBalancerUpdate cannot copy the values from %#v", other)
}

func (self *SingularityLoadBalancerUpdate) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityLoadBalancerUpdate) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityLoadBalancerUpdate) SetField(name string, value interface{}) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityLoadBalancerUpdate", name)

	case "loadBalancerState", "LoadBalancerState":
		v, ok := value.(SingularityLoadBalancerUpdateBaragonRequestState)
		if ok {
			self.LoadBalancerState = &v
			return nil
		}
		return fmt.Errorf("Field loadBalancerState/LoadBalancerState: value %v(%T) couldn't be cast to type SingularityLoadBalancerUpdateBaragonRequestState", value, value)

	case "message", "Message":
		v, ok := value.(string)
		if ok {
			self.Message = &v
			return nil
		}
		return fmt.Errorf("Field message/Message: value %v(%T) couldn't be cast to type string", value, value)

	case "timestamp", "Timestamp":
		v, ok := value.(int64)
		if ok {
			self.Timestamp = &v
			return nil
		}
		return fmt.Errorf("Field timestamp/Timestamp: value %v(%T) couldn't be cast to type int64", value, value)

	case "uri", "Uri":
		v, ok := value.(string)
		if ok {
			self.Uri = &v
			return nil
		}
		return fmt.Errorf("Field uri/Uri: value %v(%T) couldn't be cast to type string", value, value)

	case "method", "Method":
		v, ok := value.(SingularityLoadBalancerUpdateLoadBalancerMethod)
		if ok {
			self.Method = &v
			return nil
		}
		return fmt.Errorf("Field method/Method: value %v(%T) couldn't be cast to type SingularityLoadBalancerUpdateLoadBalancerMethod", value, value)

	case "loadBalancerRequestId", "LoadBalancerRequestId":
		v, ok := value.(*LoadBalancerRequestId)
		if ok {
			self.LoadBalancerRequestId = v
			return nil
		}
		return fmt.Errorf("Field loadBalancerRequestId/LoadBalancerRequestId: value %v(%T) couldn't be cast to type *LoadBalancerRequestId", value, value)

	}
}

func (self *SingularityLoadBalancerUpdate) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityLoadBalancerUpdate", name)

	case "loadBalancerState", "LoadBalancerState":
		return *self.LoadBalancerState, nil
		return nil, fmt.Errorf("Field LoadBalancerState no set on LoadBalancerState %+v", self)

	case "message", "Message":
		return *self.Message, nil
		return nil, fmt.Errorf("Field Message no set on Message %+v", self)

	case "timestamp", "Timestamp":
		return *self.Timestamp, nil
		return nil, fmt.Errorf("Field Timestamp no set on Timestamp %+v", self)

	case "uri", "Uri":
		return *self.Uri, nil
		return nil, fmt.Errorf("Field Uri no set on Uri %+v", self)

	case "method", "Method":
		return *self.Method, nil
		return nil, fmt.Errorf("Field Method no set on Method %+v", self)

	case "loadBalancerRequestId", "LoadBalancerRequestId":
		return self.LoadBalancerRequestId, nil
		return nil, fmt.Errorf("Field LoadBalancerRequestId no set on LoadBalancerRequestId %+v", self)

	}
}

func (self *SingularityLoadBalancerUpdate) ClearField(name string) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityLoadBalancerUpdate", name)

	case "loadBalancerState", "LoadBalancerState":
		self.LoadBalancerState = nil

	case "message", "Message":
		self.Message = nil

	case "timestamp", "Timestamp":
		self.Timestamp = nil

	case "uri", "Uri":
		self.Uri = nil

	case "method", "Method":
		self.Method = nil

	case "loadBalancerRequestId", "LoadBalancerRequestId":
		self.LoadBalancerRequestId = nil

	}

	return nil
}

func (self *SingularityLoadBalancerUpdate) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type SingularityLoadBalancerUpdateList []*SingularityLoadBalancerUpdate

func (self *SingularityLoadBalancerUpdateList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityLoadBalancerUpdateList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityLoadBalancerUpdateList cannot copy the values from %#v", other)
}

func (list *SingularityLoadBalancerUpdateList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
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
	return swaggering.FormatJSON(list)
}
