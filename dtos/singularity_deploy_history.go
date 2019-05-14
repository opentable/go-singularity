package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityDeployHistory struct {
	DeployResult     *SingularityDeployResult     `json:"deployResult,omitempty"`
	DeployMarker     *SingularityDeployMarker     `json:"deployMarker,omitempty"`
	Deploy           *SingularityDeploy           `json:"deploy,omitempty"`
	DeployStatistics *SingularityDeployStatistics `json:"deployStatistics,omitempty"`
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

func (self *SingularityDeployHistory) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityDeployHistory) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityDeployHistory) SetField(name string, value interface{}) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityDeployHistory", name)

	case "deployResult", "DeployResult":
		v, ok := value.(*SingularityDeployResult)
		if ok {
			self.DeployResult = v
			return nil
		}
		return fmt.Errorf("Field deployResult/DeployResult: value %v(%T) couldn't be cast to type *SingularityDeployResult", value, value)

	case "deployMarker", "DeployMarker":
		v, ok := value.(*SingularityDeployMarker)
		if ok {
			self.DeployMarker = v
			return nil
		}
		return fmt.Errorf("Field deployMarker/DeployMarker: value %v(%T) couldn't be cast to type *SingularityDeployMarker", value, value)

	case "deploy", "Deploy":
		v, ok := value.(*SingularityDeploy)
		if ok {
			self.Deploy = v
			return nil
		}
		return fmt.Errorf("Field deploy/Deploy: value %v(%T) couldn't be cast to type *SingularityDeploy", value, value)

	case "deployStatistics", "DeployStatistics":
		v, ok := value.(*SingularityDeployStatistics)
		if ok {
			self.DeployStatistics = v
			return nil
		}
		return fmt.Errorf("Field deployStatistics/DeployStatistics: value %v(%T) couldn't be cast to type *SingularityDeployStatistics", value, value)

	}
}

func (self *SingularityDeployHistory) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityDeployHistory", name)

	case "deployResult", "DeployResult":
		return self.DeployResult, nil
		return nil, fmt.Errorf("Field DeployResult no set on DeployResult %+v", self)

	case "deployMarker", "DeployMarker":
		return self.DeployMarker, nil
		return nil, fmt.Errorf("Field DeployMarker no set on DeployMarker %+v", self)

	case "deploy", "Deploy":
		return self.Deploy, nil
		return nil, fmt.Errorf("Field Deploy no set on Deploy %+v", self)

	case "deployStatistics", "DeployStatistics":
		return self.DeployStatistics, nil
		return nil, fmt.Errorf("Field DeployStatistics no set on DeployStatistics %+v", self)

	}
}

func (self *SingularityDeployHistory) ClearField(name string) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityDeployHistory", name)

	case "deployResult", "DeployResult":
		self.DeployResult = nil

	case "deployMarker", "DeployMarker":
		self.DeployMarker = nil

	case "deploy", "Deploy":
		self.Deploy = nil

	case "deployStatistics", "DeployStatistics":
		self.DeployStatistics = nil

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
