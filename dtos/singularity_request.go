package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityRequestRequestType string

const (
	SingularityRequestRequestTypeSERVICE   SingularityRequestRequestType = "SERVICE"
	SingularityRequestRequestTypeWORKER    SingularityRequestRequestType = "WORKER"
	SingularityRequestRequestTypeSCHEDULED SingularityRequestRequestType = "SCHEDULED"
	SingularityRequestRequestTypeON_DEMAND SingularityRequestRequestType = "ON_DEMAND"
	SingularityRequestRequestTypeRUN_ONCE  SingularityRequestRequestType = "RUN_ONCE"
)

type SingularityRequest struct {
	present map[string]bool

	AllowedSlaveAttributes map[string]string `json:"allowedSlaveAttributes"`

	BounceAfterScale bool `json:"bounceAfterScale"`

	// EmailConfigurationOverrides *Map[SingularityEmailType,List[SingularityEmailDestination]] `json:"emailConfigurationOverrides"`

	Group string `json:"group,omitempty"`

	Id string `json:"id,omitempty"`

	Instances int32 `json:"instances"`

	KillOldNonLongRunningTasksAfterMillis int64 `json:"killOldNonLongRunningTasksAfterMillis"`

	LoadBalanced bool `json:"loadBalanced"`

	NumRetriesOnFailure int32 `json:"numRetriesOnFailure"`

	Owners swaggering.StringList `json:"owners"`

	QuartzSchedule string `json:"quartzSchedule,omitempty"`

	RackAffinity swaggering.StringList `json:"rackAffinity"`

	RackSensitive bool `json:"rackSensitive"`

	ReadOnlyGroups swaggering.StringList `json:"readOnlyGroups"`

	RequestType SingularityRequestRequestType `json:"requestType"`

	RequiredSlaveAttributes map[string]string `json:"requiredSlaveAttributes"`

	Schedule string `json:"schedule,omitempty"`

	// ScheduleType *ScheduleType `json:"scheduleType"`

	ScheduledExpectedRuntimeMillis int64 `json:"scheduledExpectedRuntimeMillis"`

	SkipHealthchecks bool `json:"skipHealthchecks"`

	// SlavePlacement *SlavePlacement `json:"slavePlacement"`

	WaitAtLeastMillisAfterTaskFinishesForReschedule int64 `json:"waitAtLeastMillisAfterTaskFinishesForReschedule"`
}

func (self *SingularityRequest) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *SingularityRequest) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityRequest); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityRequest cannot copy the values from %#v", other)
}

func (self *SingularityRequest) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *SingularityRequest) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityRequest) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityRequest) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *SingularityRequest) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityRequest", name)

	case "allowedSlaveAttributes", "AllowedSlaveAttributes":
		v, ok := value.(map[string]string)
		if ok {
			self.AllowedSlaveAttributes = v
			self.present["allowedSlaveAttributes"] = true
			return nil
		} else {
			return fmt.Errorf("Field allowedSlaveAttributes/AllowedSlaveAttributes: value %v(%T) couldn't be cast to type map[string]string", value, value)
		}

	case "bounceAfterScale", "BounceAfterScale":
		v, ok := value.(bool)
		if ok {
			self.BounceAfterScale = v
			self.present["bounceAfterScale"] = true
			return nil
		} else {
			return fmt.Errorf("Field bounceAfterScale/BounceAfterScale: value %v(%T) couldn't be cast to type bool", value, value)
		}

	case "group", "Group":
		v, ok := value.(string)
		if ok {
			self.Group = v
			self.present["group"] = true
			return nil
		} else {
			return fmt.Errorf("Field group/Group: value %v(%T) couldn't be cast to type string", value, value)
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

	case "instances", "Instances":
		v, ok := value.(int32)
		if ok {
			self.Instances = v
			self.present["instances"] = true
			return nil
		} else {
			return fmt.Errorf("Field instances/Instances: value %v(%T) couldn't be cast to type int32", value, value)
		}

	case "killOldNonLongRunningTasksAfterMillis", "KillOldNonLongRunningTasksAfterMillis":
		v, ok := value.(int64)
		if ok {
			self.KillOldNonLongRunningTasksAfterMillis = v
			self.present["killOldNonLongRunningTasksAfterMillis"] = true
			return nil
		} else {
			return fmt.Errorf("Field killOldNonLongRunningTasksAfterMillis/KillOldNonLongRunningTasksAfterMillis: value %v(%T) couldn't be cast to type int64", value, value)
		}

	case "loadBalanced", "LoadBalanced":
		v, ok := value.(bool)
		if ok {
			self.LoadBalanced = v
			self.present["loadBalanced"] = true
			return nil
		} else {
			return fmt.Errorf("Field loadBalanced/LoadBalanced: value %v(%T) couldn't be cast to type bool", value, value)
		}

	case "numRetriesOnFailure", "NumRetriesOnFailure":
		v, ok := value.(int32)
		if ok {
			self.NumRetriesOnFailure = v
			self.present["numRetriesOnFailure"] = true
			return nil
		} else {
			return fmt.Errorf("Field numRetriesOnFailure/NumRetriesOnFailure: value %v(%T) couldn't be cast to type int32", value, value)
		}

	case "owners", "Owners":
		v, ok := value.(swaggering.StringList)
		if ok {
			self.Owners = v
			self.present["owners"] = true
			return nil
		} else {
			return fmt.Errorf("Field owners/Owners: value %v(%T) couldn't be cast to type StringList", value, value)
		}

	case "quartzSchedule", "QuartzSchedule":
		v, ok := value.(string)
		if ok {
			self.QuartzSchedule = v
			self.present["quartzSchedule"] = true
			return nil
		} else {
			return fmt.Errorf("Field quartzSchedule/QuartzSchedule: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "rackAffinity", "RackAffinity":
		v, ok := value.(swaggering.StringList)
		if ok {
			self.RackAffinity = v
			self.present["rackAffinity"] = true
			return nil
		} else {
			return fmt.Errorf("Field rackAffinity/RackAffinity: value %v(%T) couldn't be cast to type StringList", value, value)
		}

	case "rackSensitive", "RackSensitive":
		v, ok := value.(bool)
		if ok {
			self.RackSensitive = v
			self.present["rackSensitive"] = true
			return nil
		} else {
			return fmt.Errorf("Field rackSensitive/RackSensitive: value %v(%T) couldn't be cast to type bool", value, value)
		}

	case "readOnlyGroups", "ReadOnlyGroups":
		v, ok := value.(swaggering.StringList)
		if ok {
			self.ReadOnlyGroups = v
			self.present["readOnlyGroups"] = true
			return nil
		} else {
			return fmt.Errorf("Field readOnlyGroups/ReadOnlyGroups: value %v(%T) couldn't be cast to type StringList", value, value)
		}

	case "requestType", "RequestType":
		v, ok := value.(SingularityRequestRequestType)
		if ok {
			self.RequestType = v
			self.present["requestType"] = true
			return nil
		} else {
			return fmt.Errorf("Field requestType/RequestType: value %v(%T) couldn't be cast to type SingularityRequestRequestType", value, value)
		}

	case "requiredSlaveAttributes", "RequiredSlaveAttributes":
		v, ok := value.(map[string]string)
		if ok {
			self.RequiredSlaveAttributes = v
			self.present["requiredSlaveAttributes"] = true
			return nil
		} else {
			return fmt.Errorf("Field requiredSlaveAttributes/RequiredSlaveAttributes: value %v(%T) couldn't be cast to type map[string]string", value, value)
		}

	case "schedule", "Schedule":
		v, ok := value.(string)
		if ok {
			self.Schedule = v
			self.present["schedule"] = true
			return nil
		} else {
			return fmt.Errorf("Field schedule/Schedule: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "scheduledExpectedRuntimeMillis", "ScheduledExpectedRuntimeMillis":
		v, ok := value.(int64)
		if ok {
			self.ScheduledExpectedRuntimeMillis = v
			self.present["scheduledExpectedRuntimeMillis"] = true
			return nil
		} else {
			return fmt.Errorf("Field scheduledExpectedRuntimeMillis/ScheduledExpectedRuntimeMillis: value %v(%T) couldn't be cast to type int64", value, value)
		}

	case "skipHealthchecks", "SkipHealthchecks":
		v, ok := value.(bool)
		if ok {
			self.SkipHealthchecks = v
			self.present["skipHealthchecks"] = true
			return nil
		} else {
			return fmt.Errorf("Field skipHealthchecks/SkipHealthchecks: value %v(%T) couldn't be cast to type bool", value, value)
		}

	case "waitAtLeastMillisAfterTaskFinishesForReschedule", "WaitAtLeastMillisAfterTaskFinishesForReschedule":
		v, ok := value.(int64)
		if ok {
			self.WaitAtLeastMillisAfterTaskFinishesForReschedule = v
			self.present["waitAtLeastMillisAfterTaskFinishesForReschedule"] = true
			return nil
		} else {
			return fmt.Errorf("Field waitAtLeastMillisAfterTaskFinishesForReschedule/WaitAtLeastMillisAfterTaskFinishesForReschedule: value %v(%T) couldn't be cast to type int64", value, value)
		}

	}
}

func (self *SingularityRequest) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityRequest", name)

	case "allowedSlaveAttributes", "AllowedSlaveAttributes":
		if self.present != nil {
			if _, ok := self.present["allowedSlaveAttributes"]; ok {
				return self.AllowedSlaveAttributes, nil
			}
		}
		return nil, fmt.Errorf("Field AllowedSlaveAttributes no set on AllowedSlaveAttributes %+v", self)

	case "bounceAfterScale", "BounceAfterScale":
		if self.present != nil {
			if _, ok := self.present["bounceAfterScale"]; ok {
				return self.BounceAfterScale, nil
			}
		}
		return nil, fmt.Errorf("Field BounceAfterScale no set on BounceAfterScale %+v", self)

	case "group", "Group":
		if self.present != nil {
			if _, ok := self.present["group"]; ok {
				return self.Group, nil
			}
		}
		return nil, fmt.Errorf("Field Group no set on Group %+v", self)

	case "id", "Id":
		if self.present != nil {
			if _, ok := self.present["id"]; ok {
				return self.Id, nil
			}
		}
		return nil, fmt.Errorf("Field Id no set on Id %+v", self)

	case "instances", "Instances":
		if self.present != nil {
			if _, ok := self.present["instances"]; ok {
				return self.Instances, nil
			}
		}
		return nil, fmt.Errorf("Field Instances no set on Instances %+v", self)

	case "killOldNonLongRunningTasksAfterMillis", "KillOldNonLongRunningTasksAfterMillis":
		if self.present != nil {
			if _, ok := self.present["killOldNonLongRunningTasksAfterMillis"]; ok {
				return self.KillOldNonLongRunningTasksAfterMillis, nil
			}
		}
		return nil, fmt.Errorf("Field KillOldNonLongRunningTasksAfterMillis no set on KillOldNonLongRunningTasksAfterMillis %+v", self)

	case "loadBalanced", "LoadBalanced":
		if self.present != nil {
			if _, ok := self.present["loadBalanced"]; ok {
				return self.LoadBalanced, nil
			}
		}
		return nil, fmt.Errorf("Field LoadBalanced no set on LoadBalanced %+v", self)

	case "numRetriesOnFailure", "NumRetriesOnFailure":
		if self.present != nil {
			if _, ok := self.present["numRetriesOnFailure"]; ok {
				return self.NumRetriesOnFailure, nil
			}
		}
		return nil, fmt.Errorf("Field NumRetriesOnFailure no set on NumRetriesOnFailure %+v", self)

	case "owners", "Owners":
		if self.present != nil {
			if _, ok := self.present["owners"]; ok {
				return self.Owners, nil
			}
		}
		return nil, fmt.Errorf("Field Owners no set on Owners %+v", self)

	case "quartzSchedule", "QuartzSchedule":
		if self.present != nil {
			if _, ok := self.present["quartzSchedule"]; ok {
				return self.QuartzSchedule, nil
			}
		}
		return nil, fmt.Errorf("Field QuartzSchedule no set on QuartzSchedule %+v", self)

	case "rackAffinity", "RackAffinity":
		if self.present != nil {
			if _, ok := self.present["rackAffinity"]; ok {
				return self.RackAffinity, nil
			}
		}
		return nil, fmt.Errorf("Field RackAffinity no set on RackAffinity %+v", self)

	case "rackSensitive", "RackSensitive":
		if self.present != nil {
			if _, ok := self.present["rackSensitive"]; ok {
				return self.RackSensitive, nil
			}
		}
		return nil, fmt.Errorf("Field RackSensitive no set on RackSensitive %+v", self)

	case "readOnlyGroups", "ReadOnlyGroups":
		if self.present != nil {
			if _, ok := self.present["readOnlyGroups"]; ok {
				return self.ReadOnlyGroups, nil
			}
		}
		return nil, fmt.Errorf("Field ReadOnlyGroups no set on ReadOnlyGroups %+v", self)

	case "requestType", "RequestType":
		if self.present != nil {
			if _, ok := self.present["requestType"]; ok {
				return self.RequestType, nil
			}
		}
		return nil, fmt.Errorf("Field RequestType no set on RequestType %+v", self)

	case "requiredSlaveAttributes", "RequiredSlaveAttributes":
		if self.present != nil {
			if _, ok := self.present["requiredSlaveAttributes"]; ok {
				return self.RequiredSlaveAttributes, nil
			}
		}
		return nil, fmt.Errorf("Field RequiredSlaveAttributes no set on RequiredSlaveAttributes %+v", self)

	case "schedule", "Schedule":
		if self.present != nil {
			if _, ok := self.present["schedule"]; ok {
				return self.Schedule, nil
			}
		}
		return nil, fmt.Errorf("Field Schedule no set on Schedule %+v", self)

	case "scheduledExpectedRuntimeMillis", "ScheduledExpectedRuntimeMillis":
		if self.present != nil {
			if _, ok := self.present["scheduledExpectedRuntimeMillis"]; ok {
				return self.ScheduledExpectedRuntimeMillis, nil
			}
		}
		return nil, fmt.Errorf("Field ScheduledExpectedRuntimeMillis no set on ScheduledExpectedRuntimeMillis %+v", self)

	case "skipHealthchecks", "SkipHealthchecks":
		if self.present != nil {
			if _, ok := self.present["skipHealthchecks"]; ok {
				return self.SkipHealthchecks, nil
			}
		}
		return nil, fmt.Errorf("Field SkipHealthchecks no set on SkipHealthchecks %+v", self)

	case "waitAtLeastMillisAfterTaskFinishesForReschedule", "WaitAtLeastMillisAfterTaskFinishesForReschedule":
		if self.present != nil {
			if _, ok := self.present["waitAtLeastMillisAfterTaskFinishesForReschedule"]; ok {
				return self.WaitAtLeastMillisAfterTaskFinishesForReschedule, nil
			}
		}
		return nil, fmt.Errorf("Field WaitAtLeastMillisAfterTaskFinishesForReschedule no set on WaitAtLeastMillisAfterTaskFinishesForReschedule %+v", self)

	}
}

func (self *SingularityRequest) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityRequest", name)

	case "allowedSlaveAttributes", "AllowedSlaveAttributes":
		self.present["allowedSlaveAttributes"] = false

	case "bounceAfterScale", "BounceAfterScale":
		self.present["bounceAfterScale"] = false

	case "group", "Group":
		self.present["group"] = false

	case "id", "Id":
		self.present["id"] = false

	case "instances", "Instances":
		self.present["instances"] = false

	case "killOldNonLongRunningTasksAfterMillis", "KillOldNonLongRunningTasksAfterMillis":
		self.present["killOldNonLongRunningTasksAfterMillis"] = false

	case "loadBalanced", "LoadBalanced":
		self.present["loadBalanced"] = false

	case "numRetriesOnFailure", "NumRetriesOnFailure":
		self.present["numRetriesOnFailure"] = false

	case "owners", "Owners":
		self.present["owners"] = false

	case "quartzSchedule", "QuartzSchedule":
		self.present["quartzSchedule"] = false

	case "rackAffinity", "RackAffinity":
		self.present["rackAffinity"] = false

	case "rackSensitive", "RackSensitive":
		self.present["rackSensitive"] = false

	case "readOnlyGroups", "ReadOnlyGroups":
		self.present["readOnlyGroups"] = false

	case "requestType", "RequestType":
		self.present["requestType"] = false

	case "requiredSlaveAttributes", "RequiredSlaveAttributes":
		self.present["requiredSlaveAttributes"] = false

	case "schedule", "Schedule":
		self.present["schedule"] = false

	case "scheduledExpectedRuntimeMillis", "ScheduledExpectedRuntimeMillis":
		self.present["scheduledExpectedRuntimeMillis"] = false

	case "skipHealthchecks", "SkipHealthchecks":
		self.present["skipHealthchecks"] = false

	case "waitAtLeastMillisAfterTaskFinishesForReschedule", "WaitAtLeastMillisAfterTaskFinishesForReschedule":
		self.present["waitAtLeastMillisAfterTaskFinishesForReschedule"] = false

	}

	return nil
}

func (self *SingularityRequest) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type SingularityRequestList []*SingularityRequest

func (self *SingularityRequestList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityRequestList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityRequestList cannot copy the values from %#v", other)
}

func (list *SingularityRequestList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *SingularityRequestList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityRequestList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
