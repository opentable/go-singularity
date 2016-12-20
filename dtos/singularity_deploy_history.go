package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityDeployHistory struct {
	present map[string]bool

	Deploy *SingularityDeploy `json:"deploy"`

	DeployMarker *SingularityDeployMarker `json:"deployMarker"`

	DeployResult *SingularityDeployResult `json:"deployResult"`

	DeployStatistics *SingularityDeployStatistics `json:"deployStatistics"`
}

func (self *SingularityDeployHistory) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *SingularityDeployHistory) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityDeployHistory); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityDeployHistory cannot copy the values from %#v", other)
}

func (self *SingularityDeployHistory) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *SingularityDeployHistory) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityDeployHistory) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityDeployHistory) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *SingularityDeployHistory) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityDeployHistory", name)

	case "deploy", "Deploy":
		v, ok := value.(*SingularityDeploy)
		if ok {
			self.Deploy = v
			self.present["deploy"] = true
			return nil
		} else {
			return fmt.Errorf("Field deploy/Deploy: value %v(%T) couldn't be cast to type *SingularityDeploy", value, value)
		}

	case "deployMarker", "DeployMarker":
		v, ok := value.(*SingularityDeployMarker)
		if ok {
			self.DeployMarker = v
			self.present["deployMarker"] = true
			return nil
		} else {
			return fmt.Errorf("Field deployMarker/DeployMarker: value %v(%T) couldn't be cast to type *SingularityDeployMarker", value, value)
		}

	case "deployResult", "DeployResult":
		v, ok := value.(*SingularityDeployResult)
		if ok {
			self.DeployResult = v
			self.present["deployResult"] = true
			return nil
		} else {
			return fmt.Errorf("Field deployResult/DeployResult: value %v(%T) couldn't be cast to type *SingularityDeployResult", value, value)
		}

	case "deployStatistics", "DeployStatistics":
		v, ok := value.(*SingularityDeployStatistics)
		if ok {
			self.DeployStatistics = v
			self.present["deployStatistics"] = true
			return nil
		} else {
			return fmt.Errorf("Field deployStatistics/DeployStatistics: value %v(%T) couldn't be cast to type *SingularityDeployStatistics", value, value)
		}

	}
}

func (self *SingularityDeployHistory) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityDeployHistory", name)

	case "deploy", "Deploy":
		if self.present != nil {
			if _, ok := self.present["deploy"]; ok {
				return self.Deploy, nil
			}
		}
		return nil, fmt.Errorf("Field Deploy no set on Deploy %+v", self)

	case "deployMarker", "DeployMarker":
		if self.present != nil {
			if _, ok := self.present["deployMarker"]; ok {
				return self.DeployMarker, nil
			}
		}
		return nil, fmt.Errorf("Field DeployMarker no set on DeployMarker %+v", self)

	case "deployResult", "DeployResult":
		if self.present != nil {
			if _, ok := self.present["deployResult"]; ok {
				return self.DeployResult, nil
			}
		}
		return nil, fmt.Errorf("Field DeployResult no set on DeployResult %+v", self)

	case "deployStatistics", "DeployStatistics":
		if self.present != nil {
			if _, ok := self.present["deployStatistics"]; ok {
				return self.DeployStatistics, nil
			}
		}
		return nil, fmt.Errorf("Field DeployStatistics no set on DeployStatistics %+v", self)

	}
}

func (self *SingularityDeployHistory) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityDeployHistory", name)

	case "deploy", "Deploy":
		self.present["deploy"] = false

	case "deployMarker", "DeployMarker":
		self.present["deployMarker"] = false

	case "deployResult", "DeployResult":
		self.present["deployResult"] = false

	case "deployStatistics", "DeployStatistics":
		self.present["deployStatistics"] = false

	}

	return nil
}

func (self *SingularityDeployHistory) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type SingularityDeployHistoryList []*SingularityDeployHistory

func (self *SingularityDeployHistoryList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityDeployHistoryList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityDeployHistoryList cannot copy the values from %#v", other)
}

func (list *SingularityDeployHistoryList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *SingularityDeployHistoryList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityDeployHistoryList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
