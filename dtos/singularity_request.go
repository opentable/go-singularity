package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityRequestSlavePlacement string

const (
	SingularityRequestSlavePlacementSEPARATE            SingularityRequestSlavePlacement = "SEPARATE"
	SingularityRequestSlavePlacementOPTIMISTIC          SingularityRequestSlavePlacement = "OPTIMISTIC"
	SingularityRequestSlavePlacementGREEDY              SingularityRequestSlavePlacement = "GREEDY"
	SingularityRequestSlavePlacementSEPARATE_BY_DEPLOY  SingularityRequestSlavePlacement = "SEPARATE_BY_DEPLOY"
	SingularityRequestSlavePlacementSEPARATE_BY_REQUEST SingularityRequestSlavePlacement = "SEPARATE_BY_REQUEST"
	SingularityRequestSlavePlacementSPREAD_ALL_SLAVES   SingularityRequestSlavePlacement = "SPREAD_ALL_SLAVES"
)

type SingularityRequestRequestType string

const (
	SingularityRequestRequestTypeSERVICE   SingularityRequestRequestType = "SERVICE"
	SingularityRequestRequestTypeWORKER    SingularityRequestRequestType = "WORKER"
	SingularityRequestRequestTypeSCHEDULED SingularityRequestRequestType = "SCHEDULED"
	SingularityRequestRequestTypeON_DEMAND SingularityRequestRequestType = "ON_DEMAND"
	SingularityRequestRequestTypeRUN_ONCE  SingularityRequestRequestType = "RUN_ONCE"
)

type SingularityRequestScheduleType string

const (
	SingularityRequestScheduleTypeCRON    SingularityRequestScheduleType = "CRON"
	SingularityRequestScheduleTypeQUARTZ  SingularityRequestScheduleType = "QUARTZ"
	SingularityRequestScheduleTypeRFC5545 SingularityRequestScheduleType = "RFC5545"
)

type SingularityRequest struct {
	present map[string]bool

	QuartzSchedule string `json:"quartzSchedule,omitempty"`

	ScheduleTimeZone string `json:"scheduleTimeZone,omitempty"`

	TaskExecutionTimeLimitMillis int64 `json:"taskExecutionTimeLimitMillis"`

	SlavePlacement SingularityRequestSlavePlacement `json:"slavePlacement"`

	AllowBounceToSameHost bool `json:"allowBounceToSameHost"`

	ScheduledExpectedRuntimeMillis int64 `json:"scheduledExpectedRuntimeMillis"`

	LoadBalanced bool `json:"loadBalanced"`

	ReadWriteGroups swaggering.StringList `json:"readWriteGroups"`

	RequestType SingularityRequestRequestType `json:"requestType"`

	TaskLogErrorRegexCaseSensitive bool `json:"taskLogErrorRegexCaseSensitive"`

	MaxTasksPerOffer int32 `json:"maxTasksPerOffer"`

	Schedule string `json:"schedule,omitempty"`

	RackAffinity swaggering.StringList `json:"rackAffinity"`

	AllowedSlaveAttributes map[string]string `json:"allowedSlaveAttributes"`

	Group string `json:"group,omitempty"`

	HideEvenNumberAcrossRacksHint bool `json:"hideEvenNumberAcrossRacksHint"`

	Owners swaggering.StringList `json:"owners"`

	ScheduleType SingularityRequestScheduleType `json:"scheduleType"`

	KillOldNonLongRunningTasksAfterMillis int64 `json:"killOldNonLongRunningTasksAfterMillis"`

	RequiredRole string `json:"requiredRole,omitempty"`

	ReadOnlyGroups swaggering.StringList `json:"readOnlyGroups"`

	TaskLogErrorRegex string `json:"taskLogErrorRegex,omitempty"`

	NumRetriesOnFailure int32 `json:"numRetriesOnFailure"`

	WaitAtLeastMillisAfterTaskFinishesForReschedule int64 `json:"waitAtLeastMillisAfterTaskFinishesForReschedule"`

	RackSensitive bool `json:"rackSensitive"`

	RequiredSlaveAttributes map[string]string `json:"requiredSlaveAttributes"`

	BounceAfterScale bool `json:"bounceAfterScale"`

	TaskPriorityLevel float64 `json:"taskPriorityLevel"`

	Id string `json:"id,omitempty"`

	Instances int32 `json:"instances"`

	SkipHealthchecks bool `json:"skipHealthchecks"`

	// Invalid field: EmailConfigurationOverrides *notfound.Map[SingularityEmailType,List[SingularityEmailDestination]] `json:"emailConfigurationOverrides"`

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

	case "quartzSchedule", "QuartzSchedule":
		v, ok := value.(string)
		if ok {
			self.QuartzSchedule = v
			self.present["quartzSchedule"] = true
			return nil
		} else {
			return fmt.Errorf("Field quartzSchedule/QuartzSchedule: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "scheduleTimeZone", "ScheduleTimeZone":
		v, ok := value.(string)
		if ok {
			self.ScheduleTimeZone = v
			self.present["scheduleTimeZone"] = true
			return nil
		} else {
			return fmt.Errorf("Field scheduleTimeZone/ScheduleTimeZone: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "taskExecutionTimeLimitMillis", "TaskExecutionTimeLimitMillis":
		v, ok := value.(int64)
		if ok {
			self.TaskExecutionTimeLimitMillis = v
			self.present["taskExecutionTimeLimitMillis"] = true
			return nil
		} else {
			return fmt.Errorf("Field taskExecutionTimeLimitMillis/TaskExecutionTimeLimitMillis: value %v(%T) couldn't be cast to type int64", value, value)
		}

	case "slavePlacement", "SlavePlacement":
		v, ok := value.(SingularityRequestSlavePlacement)
		if ok {
			self.SlavePlacement = v
			self.present["slavePlacement"] = true
			return nil
		} else {
			return fmt.Errorf("Field slavePlacement/SlavePlacement: value %v(%T) couldn't be cast to type SingularityRequestSlavePlacement", value, value)
		}

	case "allowBounceToSameHost", "AllowBounceToSameHost":
		v, ok := value.(bool)
		if ok {
			self.AllowBounceToSameHost = v
			self.present["allowBounceToSameHost"] = true
			return nil
		} else {
			return fmt.Errorf("Field allowBounceToSameHost/AllowBounceToSameHost: value %v(%T) couldn't be cast to type bool", value, value)
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

	case "loadBalanced", "LoadBalanced":
		v, ok := value.(bool)
		if ok {
			self.LoadBalanced = v
			self.present["loadBalanced"] = true
			return nil
		} else {
			return fmt.Errorf("Field loadBalanced/LoadBalanced: value %v(%T) couldn't be cast to type bool", value, value)
		}

	case "readWriteGroups", "ReadWriteGroups":
		v, ok := value.(swaggering.StringList)
		if ok {
			self.ReadWriteGroups = v
			self.present["readWriteGroups"] = true
			return nil
		} else {
			return fmt.Errorf("Field readWriteGroups/ReadWriteGroups: value %v(%T) couldn't be cast to type swaggering.StringList", value, value)
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

	case "taskLogErrorRegexCaseSensitive", "TaskLogErrorRegexCaseSensitive":
		v, ok := value.(bool)
		if ok {
			self.TaskLogErrorRegexCaseSensitive = v
			self.present["taskLogErrorRegexCaseSensitive"] = true
			return nil
		} else {
			return fmt.Errorf("Field taskLogErrorRegexCaseSensitive/TaskLogErrorRegexCaseSensitive: value %v(%T) couldn't be cast to type bool", value, value)
		}

	case "maxTasksPerOffer", "MaxTasksPerOffer":
		v, ok := value.(int32)
		if ok {
			self.MaxTasksPerOffer = v
			self.present["maxTasksPerOffer"] = true
			return nil
		} else {
			return fmt.Errorf("Field maxTasksPerOffer/MaxTasksPerOffer: value %v(%T) couldn't be cast to type int32", value, value)
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

	case "rackAffinity", "RackAffinity":
		v, ok := value.(swaggering.StringList)
		if ok {
			self.RackAffinity = v
			self.present["rackAffinity"] = true
			return nil
		} else {
			return fmt.Errorf("Field rackAffinity/RackAffinity: value %v(%T) couldn't be cast to type swaggering.StringList", value, value)
		}

	case "allowedSlaveAttributes", "AllowedSlaveAttributes":
		v, ok := value.(map[string]string)
		if ok {
			self.AllowedSlaveAttributes = v
			self.present["allowedSlaveAttributes"] = true
			return nil
		} else {
			return fmt.Errorf("Field allowedSlaveAttributes/AllowedSlaveAttributes: value %v(%T) couldn't be cast to type map[string]string", value, value)
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

	case "hideEvenNumberAcrossRacksHint", "HideEvenNumberAcrossRacksHint":
		v, ok := value.(bool)
		if ok {
			self.HideEvenNumberAcrossRacksHint = v
			self.present["hideEvenNumberAcrossRacksHint"] = true
			return nil
		} else {
			return fmt.Errorf("Field hideEvenNumberAcrossRacksHint/HideEvenNumberAcrossRacksHint: value %v(%T) couldn't be cast to type bool", value, value)
		}

	case "owners", "Owners":
		v, ok := value.(swaggering.StringList)
		if ok {
			self.Owners = v
			self.present["owners"] = true
			return nil
		} else {
			return fmt.Errorf("Field owners/Owners: value %v(%T) couldn't be cast to type swaggering.StringList", value, value)
		}

	case "scheduleType", "ScheduleType":
		v, ok := value.(SingularityRequestScheduleType)
		if ok {
			self.ScheduleType = v
			self.present["scheduleType"] = true
			return nil
		} else {
			return fmt.Errorf("Field scheduleType/ScheduleType: value %v(%T) couldn't be cast to type SingularityRequestScheduleType", value, value)
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

	case "requiredRole", "RequiredRole":
		v, ok := value.(string)
		if ok {
			self.RequiredRole = v
			self.present["requiredRole"] = true
			return nil
		} else {
			return fmt.Errorf("Field requiredRole/RequiredRole: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "readOnlyGroups", "ReadOnlyGroups":
		v, ok := value.(swaggering.StringList)
		if ok {
			self.ReadOnlyGroups = v
			self.present["readOnlyGroups"] = true
			return nil
		} else {
			return fmt.Errorf("Field readOnlyGroups/ReadOnlyGroups: value %v(%T) couldn't be cast to type swaggering.StringList", value, value)
		}

	case "taskLogErrorRegex", "TaskLogErrorRegex":
		v, ok := value.(string)
		if ok {
			self.TaskLogErrorRegex = v
			self.present["taskLogErrorRegex"] = true
			return nil
		} else {
			return fmt.Errorf("Field taskLogErrorRegex/TaskLogErrorRegex: value %v(%T) couldn't be cast to type string", value, value)
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

	case "waitAtLeastMillisAfterTaskFinishesForReschedule", "WaitAtLeastMillisAfterTaskFinishesForReschedule":
		v, ok := value.(int64)
		if ok {
			self.WaitAtLeastMillisAfterTaskFinishesForReschedule = v
			self.present["waitAtLeastMillisAfterTaskFinishesForReschedule"] = true
			return nil
		} else {
			return fmt.Errorf("Field waitAtLeastMillisAfterTaskFinishesForReschedule/WaitAtLeastMillisAfterTaskFinishesForReschedule: value %v(%T) couldn't be cast to type int64", value, value)
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

	case "requiredSlaveAttributes", "RequiredSlaveAttributes":
		v, ok := value.(map[string]string)
		if ok {
			self.RequiredSlaveAttributes = v
			self.present["requiredSlaveAttributes"] = true
			return nil
		} else {
			return fmt.Errorf("Field requiredSlaveAttributes/RequiredSlaveAttributes: value %v(%T) couldn't be cast to type map[string]string", value, value)
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

	case "taskPriorityLevel", "TaskPriorityLevel":
		v, ok := value.(float64)
		if ok {
			self.TaskPriorityLevel = v
			self.present["taskPriorityLevel"] = true
			return nil
		} else {
			return fmt.Errorf("Field taskPriorityLevel/TaskPriorityLevel: value %v(%T) couldn't be cast to type float64", value, value)
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

	case "skipHealthchecks", "SkipHealthchecks":
		v, ok := value.(bool)
		if ok {
			self.SkipHealthchecks = v
			self.present["skipHealthchecks"] = true
			return nil
		} else {
			return fmt.Errorf("Field skipHealthchecks/SkipHealthchecks: value %v(%T) couldn't be cast to type bool", value, value)
		}

	}
}

func (self *SingularityRequest) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityRequest", name)

	case "quartzSchedule", "QuartzSchedule":
		if self.present != nil {
			if _, ok := self.present["quartzSchedule"]; ok {
				return self.QuartzSchedule, nil
			}
		}
		return nil, fmt.Errorf("Field QuartzSchedule no set on QuartzSchedule %+v", self)

	case "scheduleTimeZone", "ScheduleTimeZone":
		if self.present != nil {
			if _, ok := self.present["scheduleTimeZone"]; ok {
				return self.ScheduleTimeZone, nil
			}
		}
		return nil, fmt.Errorf("Field ScheduleTimeZone no set on ScheduleTimeZone %+v", self)

	case "taskExecutionTimeLimitMillis", "TaskExecutionTimeLimitMillis":
		if self.present != nil {
			if _, ok := self.present["taskExecutionTimeLimitMillis"]; ok {
				return self.TaskExecutionTimeLimitMillis, nil
			}
		}
		return nil, fmt.Errorf("Field TaskExecutionTimeLimitMillis no set on TaskExecutionTimeLimitMillis %+v", self)

	case "slavePlacement", "SlavePlacement":
		if self.present != nil {
			if _, ok := self.present["slavePlacement"]; ok {
				return self.SlavePlacement, nil
			}
		}
		return nil, fmt.Errorf("Field SlavePlacement no set on SlavePlacement %+v", self)

	case "allowBounceToSameHost", "AllowBounceToSameHost":
		if self.present != nil {
			if _, ok := self.present["allowBounceToSameHost"]; ok {
				return self.AllowBounceToSameHost, nil
			}
		}
		return nil, fmt.Errorf("Field AllowBounceToSameHost no set on AllowBounceToSameHost %+v", self)

	case "scheduledExpectedRuntimeMillis", "ScheduledExpectedRuntimeMillis":
		if self.present != nil {
			if _, ok := self.present["scheduledExpectedRuntimeMillis"]; ok {
				return self.ScheduledExpectedRuntimeMillis, nil
			}
		}
		return nil, fmt.Errorf("Field ScheduledExpectedRuntimeMillis no set on ScheduledExpectedRuntimeMillis %+v", self)

	case "loadBalanced", "LoadBalanced":
		if self.present != nil {
			if _, ok := self.present["loadBalanced"]; ok {
				return self.LoadBalanced, nil
			}
		}
		return nil, fmt.Errorf("Field LoadBalanced no set on LoadBalanced %+v", self)

	case "readWriteGroups", "ReadWriteGroups":
		if self.present != nil {
			if _, ok := self.present["readWriteGroups"]; ok {
				return self.ReadWriteGroups, nil
			}
		}
		return nil, fmt.Errorf("Field ReadWriteGroups no set on ReadWriteGroups %+v", self)

	case "requestType", "RequestType":
		if self.present != nil {
			if _, ok := self.present["requestType"]; ok {
				return self.RequestType, nil
			}
		}
		return nil, fmt.Errorf("Field RequestType no set on RequestType %+v", self)

	case "taskLogErrorRegexCaseSensitive", "TaskLogErrorRegexCaseSensitive":
		if self.present != nil {
			if _, ok := self.present["taskLogErrorRegexCaseSensitive"]; ok {
				return self.TaskLogErrorRegexCaseSensitive, nil
			}
		}
		return nil, fmt.Errorf("Field TaskLogErrorRegexCaseSensitive no set on TaskLogErrorRegexCaseSensitive %+v", self)

	case "maxTasksPerOffer", "MaxTasksPerOffer":
		if self.present != nil {
			if _, ok := self.present["maxTasksPerOffer"]; ok {
				return self.MaxTasksPerOffer, nil
			}
		}
		return nil, fmt.Errorf("Field MaxTasksPerOffer no set on MaxTasksPerOffer %+v", self)

	case "schedule", "Schedule":
		if self.present != nil {
			if _, ok := self.present["schedule"]; ok {
				return self.Schedule, nil
			}
		}
		return nil, fmt.Errorf("Field Schedule no set on Schedule %+v", self)

	case "rackAffinity", "RackAffinity":
		if self.present != nil {
			if _, ok := self.present["rackAffinity"]; ok {
				return self.RackAffinity, nil
			}
		}
		return nil, fmt.Errorf("Field RackAffinity no set on RackAffinity %+v", self)

	case "allowedSlaveAttributes", "AllowedSlaveAttributes":
		if self.present != nil {
			if _, ok := self.present["allowedSlaveAttributes"]; ok {
				return self.AllowedSlaveAttributes, nil
			}
		}
		return nil, fmt.Errorf("Field AllowedSlaveAttributes no set on AllowedSlaveAttributes %+v", self)

	case "group", "Group":
		if self.present != nil {
			if _, ok := self.present["group"]; ok {
				return self.Group, nil
			}
		}
		return nil, fmt.Errorf("Field Group no set on Group %+v", self)

	case "hideEvenNumberAcrossRacksHint", "HideEvenNumberAcrossRacksHint":
		if self.present != nil {
			if _, ok := self.present["hideEvenNumberAcrossRacksHint"]; ok {
				return self.HideEvenNumberAcrossRacksHint, nil
			}
		}
		return nil, fmt.Errorf("Field HideEvenNumberAcrossRacksHint no set on HideEvenNumberAcrossRacksHint %+v", self)

	case "owners", "Owners":
		if self.present != nil {
			if _, ok := self.present["owners"]; ok {
				return self.Owners, nil
			}
		}
		return nil, fmt.Errorf("Field Owners no set on Owners %+v", self)

	case "scheduleType", "ScheduleType":
		if self.present != nil {
			if _, ok := self.present["scheduleType"]; ok {
				return self.ScheduleType, nil
			}
		}
		return nil, fmt.Errorf("Field ScheduleType no set on ScheduleType %+v", self)

	case "killOldNonLongRunningTasksAfterMillis", "KillOldNonLongRunningTasksAfterMillis":
		if self.present != nil {
			if _, ok := self.present["killOldNonLongRunningTasksAfterMillis"]; ok {
				return self.KillOldNonLongRunningTasksAfterMillis, nil
			}
		}
		return nil, fmt.Errorf("Field KillOldNonLongRunningTasksAfterMillis no set on KillOldNonLongRunningTasksAfterMillis %+v", self)

	case "requiredRole", "RequiredRole":
		if self.present != nil {
			if _, ok := self.present["requiredRole"]; ok {
				return self.RequiredRole, nil
			}
		}
		return nil, fmt.Errorf("Field RequiredRole no set on RequiredRole %+v", self)

	case "readOnlyGroups", "ReadOnlyGroups":
		if self.present != nil {
			if _, ok := self.present["readOnlyGroups"]; ok {
				return self.ReadOnlyGroups, nil
			}
		}
		return nil, fmt.Errorf("Field ReadOnlyGroups no set on ReadOnlyGroups %+v", self)

	case "taskLogErrorRegex", "TaskLogErrorRegex":
		if self.present != nil {
			if _, ok := self.present["taskLogErrorRegex"]; ok {
				return self.TaskLogErrorRegex, nil
			}
		}
		return nil, fmt.Errorf("Field TaskLogErrorRegex no set on TaskLogErrorRegex %+v", self)

	case "numRetriesOnFailure", "NumRetriesOnFailure":
		if self.present != nil {
			if _, ok := self.present["numRetriesOnFailure"]; ok {
				return self.NumRetriesOnFailure, nil
			}
		}
		return nil, fmt.Errorf("Field NumRetriesOnFailure no set on NumRetriesOnFailure %+v", self)

	case "waitAtLeastMillisAfterTaskFinishesForReschedule", "WaitAtLeastMillisAfterTaskFinishesForReschedule":
		if self.present != nil {
			if _, ok := self.present["waitAtLeastMillisAfterTaskFinishesForReschedule"]; ok {
				return self.WaitAtLeastMillisAfterTaskFinishesForReschedule, nil
			}
		}
		return nil, fmt.Errorf("Field WaitAtLeastMillisAfterTaskFinishesForReschedule no set on WaitAtLeastMillisAfterTaskFinishesForReschedule %+v", self)

	case "rackSensitive", "RackSensitive":
		if self.present != nil {
			if _, ok := self.present["rackSensitive"]; ok {
				return self.RackSensitive, nil
			}
		}
		return nil, fmt.Errorf("Field RackSensitive no set on RackSensitive %+v", self)

	case "requiredSlaveAttributes", "RequiredSlaveAttributes":
		if self.present != nil {
			if _, ok := self.present["requiredSlaveAttributes"]; ok {
				return self.RequiredSlaveAttributes, nil
			}
		}
		return nil, fmt.Errorf("Field RequiredSlaveAttributes no set on RequiredSlaveAttributes %+v", self)

	case "bounceAfterScale", "BounceAfterScale":
		if self.present != nil {
			if _, ok := self.present["bounceAfterScale"]; ok {
				return self.BounceAfterScale, nil
			}
		}
		return nil, fmt.Errorf("Field BounceAfterScale no set on BounceAfterScale %+v", self)

	case "taskPriorityLevel", "TaskPriorityLevel":
		if self.present != nil {
			if _, ok := self.present["taskPriorityLevel"]; ok {
				return self.TaskPriorityLevel, nil
			}
		}
		return nil, fmt.Errorf("Field TaskPriorityLevel no set on TaskPriorityLevel %+v", self)

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

	case "skipHealthchecks", "SkipHealthchecks":
		if self.present != nil {
			if _, ok := self.present["skipHealthchecks"]; ok {
				return self.SkipHealthchecks, nil
			}
		}
		return nil, fmt.Errorf("Field SkipHealthchecks no set on SkipHealthchecks %+v", self)

	}
}

func (self *SingularityRequest) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityRequest", name)

	case "quartzSchedule", "QuartzSchedule":
		self.present["quartzSchedule"] = false

	case "scheduleTimeZone", "ScheduleTimeZone":
		self.present["scheduleTimeZone"] = false

	case "taskExecutionTimeLimitMillis", "TaskExecutionTimeLimitMillis":
		self.present["taskExecutionTimeLimitMillis"] = false

	case "slavePlacement", "SlavePlacement":
		self.present["slavePlacement"] = false

	case "allowBounceToSameHost", "AllowBounceToSameHost":
		self.present["allowBounceToSameHost"] = false

	case "scheduledExpectedRuntimeMillis", "ScheduledExpectedRuntimeMillis":
		self.present["scheduledExpectedRuntimeMillis"] = false

	case "loadBalanced", "LoadBalanced":
		self.present["loadBalanced"] = false

	case "readWriteGroups", "ReadWriteGroups":
		self.present["readWriteGroups"] = false

	case "requestType", "RequestType":
		self.present["requestType"] = false

	case "taskLogErrorRegexCaseSensitive", "TaskLogErrorRegexCaseSensitive":
		self.present["taskLogErrorRegexCaseSensitive"] = false

	case "maxTasksPerOffer", "MaxTasksPerOffer":
		self.present["maxTasksPerOffer"] = false

	case "schedule", "Schedule":
		self.present["schedule"] = false

	case "rackAffinity", "RackAffinity":
		self.present["rackAffinity"] = false

	case "allowedSlaveAttributes", "AllowedSlaveAttributes":
		self.present["allowedSlaveAttributes"] = false

	case "group", "Group":
		self.present["group"] = false

	case "hideEvenNumberAcrossRacksHint", "HideEvenNumberAcrossRacksHint":
		self.present["hideEvenNumberAcrossRacksHint"] = false

	case "owners", "Owners":
		self.present["owners"] = false

	case "scheduleType", "ScheduleType":
		self.present["scheduleType"] = false

	case "killOldNonLongRunningTasksAfterMillis", "KillOldNonLongRunningTasksAfterMillis":
		self.present["killOldNonLongRunningTasksAfterMillis"] = false

	case "requiredRole", "RequiredRole":
		self.present["requiredRole"] = false

	case "readOnlyGroups", "ReadOnlyGroups":
		self.present["readOnlyGroups"] = false

	case "taskLogErrorRegex", "TaskLogErrorRegex":
		self.present["taskLogErrorRegex"] = false

	case "numRetriesOnFailure", "NumRetriesOnFailure":
		self.present["numRetriesOnFailure"] = false

	case "waitAtLeastMillisAfterTaskFinishesForReschedule", "WaitAtLeastMillisAfterTaskFinishesForReschedule":
		self.present["waitAtLeastMillisAfterTaskFinishesForReschedule"] = false

	case "rackSensitive", "RackSensitive":
		self.present["rackSensitive"] = false

	case "requiredSlaveAttributes", "RequiredSlaveAttributes":
		self.present["requiredSlaveAttributes"] = false

	case "bounceAfterScale", "BounceAfterScale":
		self.present["bounceAfterScale"] = false

	case "taskPriorityLevel", "TaskPriorityLevel":
		self.present["taskPriorityLevel"] = false

	case "id", "Id":
		self.present["id"] = false

	case "instances", "Instances":
		self.present["instances"] = false

	case "skipHealthchecks", "SkipHealthchecks":
		self.present["skipHealthchecks"] = false

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
