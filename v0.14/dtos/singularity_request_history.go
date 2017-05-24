package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityRequestHistoryRequestHistoryType string

const (
	SingularityRequestHistoryRequestHistoryTypeCREATED             SingularityRequestHistoryRequestHistoryType = "CREATED"
	SingularityRequestHistoryRequestHistoryTypeUPDATED             SingularityRequestHistoryRequestHistoryType = "UPDATED"
	SingularityRequestHistoryRequestHistoryTypeDELETED             SingularityRequestHistoryRequestHistoryType = "DELETED"
	SingularityRequestHistoryRequestHistoryTypePAUSED              SingularityRequestHistoryRequestHistoryType = "PAUSED"
	SingularityRequestHistoryRequestHistoryTypeUNPAUSED            SingularityRequestHistoryRequestHistoryType = "UNPAUSED"
	SingularityRequestHistoryRequestHistoryTypeENTERED_COOLDOWN    SingularityRequestHistoryRequestHistoryType = "ENTERED_COOLDOWN"
	SingularityRequestHistoryRequestHistoryTypeEXITED_COOLDOWN     SingularityRequestHistoryRequestHistoryType = "EXITED_COOLDOWN"
	SingularityRequestHistoryRequestHistoryTypeFINISHED            SingularityRequestHistoryRequestHistoryType = "FINISHED"
	SingularityRequestHistoryRequestHistoryTypeDEPLOYED_TO_UNPAUSE SingularityRequestHistoryRequestHistoryType = "DEPLOYED_TO_UNPAUSE"
	SingularityRequestHistoryRequestHistoryTypeBOUNCED             SingularityRequestHistoryRequestHistoryType = "BOUNCED"
	SingularityRequestHistoryRequestHistoryTypeSCALED              SingularityRequestHistoryRequestHistoryType = "SCALED"
	SingularityRequestHistoryRequestHistoryTypeSCALE_REVERTED      SingularityRequestHistoryRequestHistoryType = "SCALE_REVERTED"
)

type SingularityRequestHistory struct {
	present map[string]bool

	CreatedAt int64 `json:"createdAt"`

	EventType SingularityRequestHistoryRequestHistoryType `json:"eventType"`

	Message string `json:"message,omitempty"`

	Request *SingularityRequest `json:"request"`

	User string `json:"user,omitempty"`
}

func (self *SingularityRequestHistory) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *SingularityRequestHistory) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityRequestHistory); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityRequestHistory cannot copy the values from %#v", other)
}

func (self *SingularityRequestHistory) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *SingularityRequestHistory) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityRequestHistory) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityRequestHistory) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *SingularityRequestHistory) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityRequestHistory", name)

	case "createdAt", "CreatedAt":
		v, ok := value.(int64)
		if ok {
			self.CreatedAt = v
			self.present["createdAt"] = true
			return nil
		} else {
			return fmt.Errorf("Field createdAt/CreatedAt: value %v(%T) couldn't be cast to type int64", value, value)
		}

	case "eventType", "EventType":
		v, ok := value.(SingularityRequestHistoryRequestHistoryType)
		if ok {
			self.EventType = v
			self.present["eventType"] = true
			return nil
		} else {
			return fmt.Errorf("Field eventType/EventType: value %v(%T) couldn't be cast to type SingularityRequestHistoryRequestHistoryType", value, value)
		}

	case "message", "Message":
		v, ok := value.(string)
		if ok {
			self.Message = v
			self.present["message"] = true
			return nil
		} else {
			return fmt.Errorf("Field message/Message: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "request", "Request":
		v, ok := value.(*SingularityRequest)
		if ok {
			self.Request = v
			self.present["request"] = true
			return nil
		} else {
			return fmt.Errorf("Field request/Request: value %v(%T) couldn't be cast to type *SingularityRequest", value, value)
		}

	case "user", "User":
		v, ok := value.(string)
		if ok {
			self.User = v
			self.present["user"] = true
			return nil
		} else {
			return fmt.Errorf("Field user/User: value %v(%T) couldn't be cast to type string", value, value)
		}

	}
}

func (self *SingularityRequestHistory) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityRequestHistory", name)

	case "createdAt", "CreatedAt":
		if self.present != nil {
			if _, ok := self.present["createdAt"]; ok {
				return self.CreatedAt, nil
			}
		}
		return nil, fmt.Errorf("Field CreatedAt no set on CreatedAt %+v", self)

	case "eventType", "EventType":
		if self.present != nil {
			if _, ok := self.present["eventType"]; ok {
				return self.EventType, nil
			}
		}
		return nil, fmt.Errorf("Field EventType no set on EventType %+v", self)

	case "message", "Message":
		if self.present != nil {
			if _, ok := self.present["message"]; ok {
				return self.Message, nil
			}
		}
		return nil, fmt.Errorf("Field Message no set on Message %+v", self)

	case "request", "Request":
		if self.present != nil {
			if _, ok := self.present["request"]; ok {
				return self.Request, nil
			}
		}
		return nil, fmt.Errorf("Field Request no set on Request %+v", self)

	case "user", "User":
		if self.present != nil {
			if _, ok := self.present["user"]; ok {
				return self.User, nil
			}
		}
		return nil, fmt.Errorf("Field User no set on User %+v", self)

	}
}

func (self *SingularityRequestHistory) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityRequestHistory", name)

	case "createdAt", "CreatedAt":
		self.present["createdAt"] = false

	case "eventType", "EventType":
		self.present["eventType"] = false

	case "message", "Message":
		self.present["message"] = false

	case "request", "Request":
		self.present["request"] = false

	case "user", "User":
		self.present["user"] = false

	}

	return nil
}

func (self *SingularityRequestHistory) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type SingularityRequestHistoryList []*SingularityRequestHistory

func (self *SingularityRequestHistoryList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityRequestHistoryList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityRequestHistoryList cannot copy the values from %#v", other)
}

func (list *SingularityRequestHistoryList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *SingularityRequestHistoryList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityRequestHistoryList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
