package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityPendingTaskIdPendingType string

const (
	SingularityPendingTaskIdPendingTypeIMMEDIATE                   SingularityPendingTaskIdPendingType = "IMMEDIATE"
	SingularityPendingTaskIdPendingTypeONEOFF                      SingularityPendingTaskIdPendingType = "ONEOFF"
	SingularityPendingTaskIdPendingTypeBOUNCE                      SingularityPendingTaskIdPendingType = "BOUNCE"
	SingularityPendingTaskIdPendingTypeNEW_DEPLOY                  SingularityPendingTaskIdPendingType = "NEW_DEPLOY"
	SingularityPendingTaskIdPendingTypeNEXT_DEPLOY_STEP            SingularityPendingTaskIdPendingType = "NEXT_DEPLOY_STEP"
	SingularityPendingTaskIdPendingTypeUNPAUSED                    SingularityPendingTaskIdPendingType = "UNPAUSED"
	SingularityPendingTaskIdPendingTypeRETRY                       SingularityPendingTaskIdPendingType = "RETRY"
	SingularityPendingTaskIdPendingTypeUPDATED_REQUEST             SingularityPendingTaskIdPendingType = "UPDATED_REQUEST"
	SingularityPendingTaskIdPendingTypeDECOMISSIONED_SLAVE_OR_RACK SingularityPendingTaskIdPendingType = "DECOMISSIONED_SLAVE_OR_RACK"
	SingularityPendingTaskIdPendingTypeTASK_DONE                   SingularityPendingTaskIdPendingType = "TASK_DONE"
	SingularityPendingTaskIdPendingTypeSTARTUP                     SingularityPendingTaskIdPendingType = "STARTUP"
	SingularityPendingTaskIdPendingTypeCANCEL_BOUNCE               SingularityPendingTaskIdPendingType = "CANCEL_BOUNCE"
	SingularityPendingTaskIdPendingTypeTASK_BOUNCE                 SingularityPendingTaskIdPendingType = "TASK_BOUNCE"
	SingularityPendingTaskIdPendingTypeDEPLOY_CANCELLED            SingularityPendingTaskIdPendingType = "DEPLOY_CANCELLED"
)

type SingularityPendingTaskId struct {
	present map[string]bool

	CreatedAt int64 `json:"createdAt"`

	DeployId string `json:"deployId,omitempty"`

	Id string `json:"id,omitempty"`

	InstanceNo int32 `json:"instanceNo"`

	NextRunAt int64 `json:"nextRunAt"`

	PendingType SingularityPendingTaskIdPendingType `json:"pendingType"`

	RequestId string `json:"requestId,omitempty"`
}

func (self *SingularityPendingTaskId) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *SingularityPendingTaskId) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityPendingTaskId); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityPendingTaskId cannot copy the values from %#v", other)
}

func (self *SingularityPendingTaskId) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *SingularityPendingTaskId) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityPendingTaskId) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityPendingTaskId) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *SingularityPendingTaskId) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityPendingTaskId", name)

	case "createdAt", "CreatedAt":
		v, ok := value.(int64)
		if ok {
			self.CreatedAt = v
			self.present["createdAt"] = true
			return nil
		} else {
			return fmt.Errorf("Field createdAt/CreatedAt: value %v(%T) couldn't be cast to type int64", value, value)
		}

	case "deployId", "DeployId":
		v, ok := value.(string)
		if ok {
			self.DeployId = v
			self.present["deployId"] = true
			return nil
		} else {
			return fmt.Errorf("Field deployId/DeployId: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "id", "Id":
		v, ok := value.(string)
		if ok {
			self.Id = v
			self.present["id"] = true
			return nil
		} else {
			return fmt.Errorf("Field id/Id: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "instanceNo", "InstanceNo":
		v, ok := value.(int32)
		if ok {
			self.InstanceNo = v
			self.present["instanceNo"] = true
			return nil
		} else {
			return fmt.Errorf("Field instanceNo/InstanceNo: value %v(%T) couldn't be cast to type int32", value, value)
		}

	case "nextRunAt", "NextRunAt":
		v, ok := value.(int64)
		if ok {
			self.NextRunAt = v
			self.present["nextRunAt"] = true
			return nil
		} else {
			return fmt.Errorf("Field nextRunAt/NextRunAt: value %v(%T) couldn't be cast to type int64", value, value)
		}

	case "pendingType", "PendingType":
		v, ok := value.(SingularityPendingTaskIdPendingType)
		if ok {
			self.PendingType = v
			self.present["pendingType"] = true
			return nil
		} else {
			return fmt.Errorf("Field pendingType/PendingType: value %v(%T) couldn't be cast to type SingularityPendingTaskIdPendingType", value, value)
		}

	case "requestId", "RequestId":
		v, ok := value.(string)
		if ok {
			self.RequestId = v
			self.present["requestId"] = true
			return nil
		} else {
			return fmt.Errorf("Field requestId/RequestId: value %v(%T) couldn't be cast to type string", value, value)
		}

	}
}

func (self *SingularityPendingTaskId) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityPendingTaskId", name)

	case "createdAt", "CreatedAt":
		if self.present != nil {
			if _, ok := self.present["createdAt"]; ok {
				return self.CreatedAt, nil
			}
		}
		return nil, fmt.Errorf("Field CreatedAt no set on CreatedAt %+v", self)

	case "deployId", "DeployId":
		if self.present != nil {
			if _, ok := self.present["deployId"]; ok {
				return self.DeployId, nil
			}
		}
		return nil, fmt.Errorf("Field DeployId no set on DeployId %+v", self)

	case "id", "Id":
		if self.present != nil {
			if _, ok := self.present["id"]; ok {
				return self.Id, nil
			}
		}
		return nil, fmt.Errorf("Field Id no set on Id %+v", self)

	case "instanceNo", "InstanceNo":
		if self.present != nil {
			if _, ok := self.present["instanceNo"]; ok {
				return self.InstanceNo, nil
			}
		}
		return nil, fmt.Errorf("Field InstanceNo no set on InstanceNo %+v", self)

	case "nextRunAt", "NextRunAt":
		if self.present != nil {
			if _, ok := self.present["nextRunAt"]; ok {
				return self.NextRunAt, nil
			}
		}
		return nil, fmt.Errorf("Field NextRunAt no set on NextRunAt %+v", self)

	case "pendingType", "PendingType":
		if self.present != nil {
			if _, ok := self.present["pendingType"]; ok {
				return self.PendingType, nil
			}
		}
		return nil, fmt.Errorf("Field PendingType no set on PendingType %+v", self)

	case "requestId", "RequestId":
		if self.present != nil {
			if _, ok := self.present["requestId"]; ok {
				return self.RequestId, nil
			}
		}
		return nil, fmt.Errorf("Field RequestId no set on RequestId %+v", self)

	}
}

func (self *SingularityPendingTaskId) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityPendingTaskId", name)

	case "createdAt", "CreatedAt":
		self.present["createdAt"] = false

	case "deployId", "DeployId":
		self.present["deployId"] = false

	case "id", "Id":
		self.present["id"] = false

	case "instanceNo", "InstanceNo":
		self.present["instanceNo"] = false

	case "nextRunAt", "NextRunAt":
		self.present["nextRunAt"] = false

	case "pendingType", "PendingType":
		self.present["pendingType"] = false

	case "requestId", "RequestId":
		self.present["requestId"] = false

	}

	return nil
}

func (self *SingularityPendingTaskId) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type SingularityPendingTaskIdList []*SingularityPendingTaskId

func (self *SingularityPendingTaskIdList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityPendingTaskIdList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityPendingTaskIdList cannot copy the values from %#v", other)
}

func (list *SingularityPendingTaskIdList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *SingularityPendingTaskIdList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityPendingTaskIdList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
