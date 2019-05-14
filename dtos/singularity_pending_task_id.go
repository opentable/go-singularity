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
	SingularityPendingTaskIdPendingTypeDEPLOY_FAILED               SingularityPendingTaskIdPendingType = "DEPLOY_FAILED"
)

type SingularityPendingTaskId struct {
	InstanceNo  *int32                               `json:"instanceNo,omitempty"`
	PendingType *SingularityPendingTaskIdPendingType `json:"pendingType,omitempty"`
	Id          *string                              `json:"id,omitempty"`
	RequestId   *string                              `json:"requestId,omitempty"`
	DeployId    *string                              `json:"deployId,omitempty"`
	NextRunAt   *int64                               `json:"nextRunAt,omitempty"`
	CreatedAt   *int64                               `json:"createdAt,omitempty"`
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

func (self *SingularityPendingTaskId) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityPendingTaskId) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityPendingTaskId) SetField(name string, value interface{}) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityPendingTaskId", name)

	case "instanceNo", "InstanceNo":
		v, ok := value.(int32)
		if ok {
			self.InstanceNo = &v
			return nil
		}
		return fmt.Errorf("Field instanceNo/InstanceNo: value %v(%T) couldn't be cast to type int32", value, value)

	case "pendingType", "PendingType":
		v, ok := value.(SingularityPendingTaskIdPendingType)
		if ok {
			self.PendingType = &v
			return nil
		}
		return fmt.Errorf("Field pendingType/PendingType: value %v(%T) couldn't be cast to type SingularityPendingTaskIdPendingType", value, value)

	case "id", "Id":
		v, ok := value.(string)
		if ok {
			self.Id = &v
			return nil
		}
		return fmt.Errorf("Field id/Id: value %v(%T) couldn't be cast to type string", value, value)

	case "requestId", "RequestId":
		v, ok := value.(string)
		if ok {
			self.RequestId = &v
			return nil
		}
		return fmt.Errorf("Field requestId/RequestId: value %v(%T) couldn't be cast to type string", value, value)

	case "deployId", "DeployId":
		v, ok := value.(string)
		if ok {
			self.DeployId = &v
			return nil
		}
		return fmt.Errorf("Field deployId/DeployId: value %v(%T) couldn't be cast to type string", value, value)

	case "nextRunAt", "NextRunAt":
		v, ok := value.(int64)
		if ok {
			self.NextRunAt = &v
			return nil
		}
		return fmt.Errorf("Field nextRunAt/NextRunAt: value %v(%T) couldn't be cast to type int64", value, value)

	case "createdAt", "CreatedAt":
		v, ok := value.(int64)
		if ok {
			self.CreatedAt = &v
			return nil
		}
		return fmt.Errorf("Field createdAt/CreatedAt: value %v(%T) couldn't be cast to type int64", value, value)

	}
}

func (self *SingularityPendingTaskId) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityPendingTaskId", name)

	case "instanceNo", "InstanceNo":
		return *self.InstanceNo, nil
		return nil, fmt.Errorf("Field InstanceNo no set on InstanceNo %+v", self)

	case "pendingType", "PendingType":
		return *self.PendingType, nil
		return nil, fmt.Errorf("Field PendingType no set on PendingType %+v", self)

	case "id", "Id":
		return *self.Id, nil
		return nil, fmt.Errorf("Field Id no set on Id %+v", self)

	case "requestId", "RequestId":
		return *self.RequestId, nil
		return nil, fmt.Errorf("Field RequestId no set on RequestId %+v", self)

	case "deployId", "DeployId":
		return *self.DeployId, nil
		return nil, fmt.Errorf("Field DeployId no set on DeployId %+v", self)

	case "nextRunAt", "NextRunAt":
		return *self.NextRunAt, nil
		return nil, fmt.Errorf("Field NextRunAt no set on NextRunAt %+v", self)

	case "createdAt", "CreatedAt":
		return *self.CreatedAt, nil
		return nil, fmt.Errorf("Field CreatedAt no set on CreatedAt %+v", self)

	}
}

func (self *SingularityPendingTaskId) ClearField(name string) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityPendingTaskId", name)

	case "instanceNo", "InstanceNo":
		self.InstanceNo = nil

	case "pendingType", "PendingType":
		self.PendingType = nil

	case "id", "Id":
		self.Id = nil

	case "requestId", "RequestId":
		self.RequestId = nil

	case "deployId", "DeployId":
		self.DeployId = nil

	case "nextRunAt", "NextRunAt":
		self.NextRunAt = nil

	case "createdAt", "CreatedAt":
		self.CreatedAt = nil

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
