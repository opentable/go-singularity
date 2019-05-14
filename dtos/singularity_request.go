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

type SingularityRequestScheduleType string

const (
	SingularityRequestScheduleTypeCRON    SingularityRequestScheduleType = "CRON"
	SingularityRequestScheduleTypeQUARTZ  SingularityRequestScheduleType = "QUARTZ"
	SingularityRequestScheduleTypeRFC5545 SingularityRequestScheduleType = "RFC5545"
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
	SlavePlacement                        *SingularityRequestSlavePlacement `json:"slavePlacement,omitempty"`
	LoadBalanced                          *bool                             `json:"loadBalanced,omitempty"`
	MaxTasksPerOffer                      *int32                            `json:"maxTasksPerOffer,omitempty"`
	RackSensitive                         *bool                             `json:"rackSensitive,omitempty"`
	ScheduleType                          *SingularityRequestScheduleType   `json:"scheduleType,omitempty"`
	KillOldNonLongRunningTasksAfterMillis *int64                            `json:"killOldNonLongRunningTasksAfterMillis,omitempty"`
	Instances                             *int32                            `json:"instances,omitempty"`
	RackAffinity                          *swaggering.StringList            `json:"rackAffinity,omitempty"`
	HideEvenNumberAcrossRacksHint         *bool                             `json:"hideEvenNumberAcrossRacksHint,omitempty"`
	AllowBounceToSameHost                 *bool                             `json:"allowBounceToSameHost,omitempty"`
	NumRetriesOnFailure                   *int32                            `json:"numRetriesOnFailure,omitempty"`
	TaskLogErrorRegexCaseSensitive        *bool                             `json:"taskLogErrorRegexCaseSensitive,omitempty"`
	RequestType                           *SingularityRequestRequestType    `json:"requestType,omitempty"`
	ReadOnlyGroups                        *swaggering.StringList            `json:"readOnlyGroups,omitempty"`
	ScheduleTimeZone                      *string                           `json:"scheduleTimeZone,omitempty"`
	SkipHealthchecks                      *bool                             `json:"skipHealthchecks,omitempty"`
	Group                                 *string                           `json:"group,omitempty"`
	// Invalid field: EmailConfigurationOverrides *notfound.Map[SingularityEmailType,List[SingularityEmailDestination]] `json:"emailConfigurationOverrides,omitempty"`
	Owners                                          *swaggering.StringList `json:"owners,omitempty"`
	AllowedSlaveAttributes                          *map[string]string     `json:"allowedSlaveAttributes,omitempty"`
	RequiredRole                                    *string                `json:"requiredRole,omitempty"`
	TaskExecutionTimeLimitMillis                    *int64                 `json:"taskExecutionTimeLimitMillis,omitempty"`
	QuartzSchedule                                  *string                `json:"quartzSchedule,omitempty"`
	ScheduledExpectedRuntimeMillis                  *int64                 `json:"scheduledExpectedRuntimeMillis,omitempty"`
	WaitAtLeastMillisAfterTaskFinishesForReschedule *int64                 `json:"waitAtLeastMillisAfterTaskFinishesForReschedule,omitempty"`
	RequiredSlaveAttributes                         *map[string]string     `json:"requiredSlaveAttributes,omitempty"`
	ReadWriteGroups                                 *swaggering.StringList `json:"readWriteGroups,omitempty"`
	TaskLogErrorRegex                               *string                `json:"taskLogErrorRegex,omitempty"`
	TaskPriorityLevel                               *float64               `json:"taskPriorityLevel,omitempty"`
	Schedule                                        *string                `json:"schedule,omitempty"`
	BounceAfterScale                                *bool                  `json:"bounceAfterScale,omitempty"`
	Id                                              *string                `json:"id,omitempty"`
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

func (self *SingularityRequest) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityRequest) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityRequest) SetField(name string, value interface{}) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityRequest", name)

	case "slavePlacement", "SlavePlacement":
		v, ok := value.(SingularityRequestSlavePlacement)
		if ok {
			self.SlavePlacement = &v
			return nil
		}
		return fmt.Errorf("Field slavePlacement/SlavePlacement: value %v(%T) couldn't be cast to type SingularityRequestSlavePlacement", value, value)

	case "loadBalanced", "LoadBalanced":
		v, ok := value.(bool)
		if ok {
			self.LoadBalanced = &v
			return nil
		}
		return fmt.Errorf("Field loadBalanced/LoadBalanced: value %v(%T) couldn't be cast to type bool", value, value)

	case "maxTasksPerOffer", "MaxTasksPerOffer":
		v, ok := value.(int32)
		if ok {
			self.MaxTasksPerOffer = &v
			return nil
		}
		return fmt.Errorf("Field maxTasksPerOffer/MaxTasksPerOffer: value %v(%T) couldn't be cast to type int32", value, value)

	case "rackSensitive", "RackSensitive":
		v, ok := value.(bool)
		if ok {
			self.RackSensitive = &v
			return nil
		}
		return fmt.Errorf("Field rackSensitive/RackSensitive: value %v(%T) couldn't be cast to type bool", value, value)

	case "scheduleType", "ScheduleType":
		v, ok := value.(SingularityRequestScheduleType)
		if ok {
			self.ScheduleType = &v
			return nil
		}
		return fmt.Errorf("Field scheduleType/ScheduleType: value %v(%T) couldn't be cast to type SingularityRequestScheduleType", value, value)

	case "killOldNonLongRunningTasksAfterMillis", "KillOldNonLongRunningTasksAfterMillis":
		v, ok := value.(int64)
		if ok {
			self.KillOldNonLongRunningTasksAfterMillis = &v
			return nil
		}
		return fmt.Errorf("Field killOldNonLongRunningTasksAfterMillis/KillOldNonLongRunningTasksAfterMillis: value %v(%T) couldn't be cast to type int64", value, value)

	case "instances", "Instances":
		v, ok := value.(int32)
		if ok {
			self.Instances = &v
			return nil
		}
		return fmt.Errorf("Field instances/Instances: value %v(%T) couldn't be cast to type int32", value, value)

	case "rackAffinity", "RackAffinity":
		v, ok := value.(swaggering.StringList)
		if ok {
			self.RackAffinity = &v
			return nil
		}
		return fmt.Errorf("Field rackAffinity/RackAffinity: value %v(%T) couldn't be cast to type swaggering.StringList", value, value)

	case "hideEvenNumberAcrossRacksHint", "HideEvenNumberAcrossRacksHint":
		v, ok := value.(bool)
		if ok {
			self.HideEvenNumberAcrossRacksHint = &v
			return nil
		}
		return fmt.Errorf("Field hideEvenNumberAcrossRacksHint/HideEvenNumberAcrossRacksHint: value %v(%T) couldn't be cast to type bool", value, value)

	case "allowBounceToSameHost", "AllowBounceToSameHost":
		v, ok := value.(bool)
		if ok {
			self.AllowBounceToSameHost = &v
			return nil
		}
		return fmt.Errorf("Field allowBounceToSameHost/AllowBounceToSameHost: value %v(%T) couldn't be cast to type bool", value, value)

	case "numRetriesOnFailure", "NumRetriesOnFailure":
		v, ok := value.(int32)
		if ok {
			self.NumRetriesOnFailure = &v
			return nil
		}
		return fmt.Errorf("Field numRetriesOnFailure/NumRetriesOnFailure: value %v(%T) couldn't be cast to type int32", value, value)

	case "taskLogErrorRegexCaseSensitive", "TaskLogErrorRegexCaseSensitive":
		v, ok := value.(bool)
		if ok {
			self.TaskLogErrorRegexCaseSensitive = &v
			return nil
		}
		return fmt.Errorf("Field taskLogErrorRegexCaseSensitive/TaskLogErrorRegexCaseSensitive: value %v(%T) couldn't be cast to type bool", value, value)

	case "requestType", "RequestType":
		v, ok := value.(SingularityRequestRequestType)
		if ok {
			self.RequestType = &v
			return nil
		}
		return fmt.Errorf("Field requestType/RequestType: value %v(%T) couldn't be cast to type SingularityRequestRequestType", value, value)

	case "readOnlyGroups", "ReadOnlyGroups":
		v, ok := value.(swaggering.StringList)
		if ok {
			self.ReadOnlyGroups = &v
			return nil
		}
		return fmt.Errorf("Field readOnlyGroups/ReadOnlyGroups: value %v(%T) couldn't be cast to type swaggering.StringList", value, value)

	case "scheduleTimeZone", "ScheduleTimeZone":
		v, ok := value.(string)
		if ok {
			self.ScheduleTimeZone = &v
			return nil
		}
		return fmt.Errorf("Field scheduleTimeZone/ScheduleTimeZone: value %v(%T) couldn't be cast to type string", value, value)

	case "skipHealthchecks", "SkipHealthchecks":
		v, ok := value.(bool)
		if ok {
			self.SkipHealthchecks = &v
			return nil
		}
		return fmt.Errorf("Field skipHealthchecks/SkipHealthchecks: value %v(%T) couldn't be cast to type bool", value, value)

	case "group", "Group":
		v, ok := value.(string)
		if ok {
			self.Group = &v
			return nil
		}
		return fmt.Errorf("Field group/Group: value %v(%T) couldn't be cast to type string", value, value)

	case "owners", "Owners":
		v, ok := value.(swaggering.StringList)
		if ok {
			self.Owners = &v
			return nil
		}
		return fmt.Errorf("Field owners/Owners: value %v(%T) couldn't be cast to type swaggering.StringList", value, value)

	case "allowedSlaveAttributes", "AllowedSlaveAttributes":
		v, ok := value.(map[string]string)
		if ok {
			self.AllowedSlaveAttributes = &v
			return nil
		}
		return fmt.Errorf("Field allowedSlaveAttributes/AllowedSlaveAttributes: value %v(%T) couldn't be cast to type map[string]string", value, value)

	case "requiredRole", "RequiredRole":
		v, ok := value.(string)
		if ok {
			self.RequiredRole = &v
			return nil
		}
		return fmt.Errorf("Field requiredRole/RequiredRole: value %v(%T) couldn't be cast to type string", value, value)

	case "taskExecutionTimeLimitMillis", "TaskExecutionTimeLimitMillis":
		v, ok := value.(int64)
		if ok {
			self.TaskExecutionTimeLimitMillis = &v
			return nil
		}
		return fmt.Errorf("Field taskExecutionTimeLimitMillis/TaskExecutionTimeLimitMillis: value %v(%T) couldn't be cast to type int64", value, value)

	case "quartzSchedule", "QuartzSchedule":
		v, ok := value.(string)
		if ok {
			self.QuartzSchedule = &v
			return nil
		}
		return fmt.Errorf("Field quartzSchedule/QuartzSchedule: value %v(%T) couldn't be cast to type string", value, value)

	case "scheduledExpectedRuntimeMillis", "ScheduledExpectedRuntimeMillis":
		v, ok := value.(int64)
		if ok {
			self.ScheduledExpectedRuntimeMillis = &v
			return nil
		}
		return fmt.Errorf("Field scheduledExpectedRuntimeMillis/ScheduledExpectedRuntimeMillis: value %v(%T) couldn't be cast to type int64", value, value)

	case "waitAtLeastMillisAfterTaskFinishesForReschedule", "WaitAtLeastMillisAfterTaskFinishesForReschedule":
		v, ok := value.(int64)
		if ok {
			self.WaitAtLeastMillisAfterTaskFinishesForReschedule = &v
			return nil
		}
		return fmt.Errorf("Field waitAtLeastMillisAfterTaskFinishesForReschedule/WaitAtLeastMillisAfterTaskFinishesForReschedule: value %v(%T) couldn't be cast to type int64", value, value)

	case "requiredSlaveAttributes", "RequiredSlaveAttributes":
		v, ok := value.(map[string]string)
		if ok {
			self.RequiredSlaveAttributes = &v
			return nil
		}
		return fmt.Errorf("Field requiredSlaveAttributes/RequiredSlaveAttributes: value %v(%T) couldn't be cast to type map[string]string", value, value)

	case "readWriteGroups", "ReadWriteGroups":
		v, ok := value.(swaggering.StringList)
		if ok {
			self.ReadWriteGroups = &v
			return nil
		}
		return fmt.Errorf("Field readWriteGroups/ReadWriteGroups: value %v(%T) couldn't be cast to type swaggering.StringList", value, value)

	case "taskLogErrorRegex", "TaskLogErrorRegex":
		v, ok := value.(string)
		if ok {
			self.TaskLogErrorRegex = &v
			return nil
		}
		return fmt.Errorf("Field taskLogErrorRegex/TaskLogErrorRegex: value %v(%T) couldn't be cast to type string", value, value)

	case "taskPriorityLevel", "TaskPriorityLevel":
		v, ok := value.(float64)
		if ok {
			self.TaskPriorityLevel = &v
			return nil
		}
		return fmt.Errorf("Field taskPriorityLevel/TaskPriorityLevel: value %v(%T) couldn't be cast to type float64", value, value)

	case "schedule", "Schedule":
		v, ok := value.(string)
		if ok {
			self.Schedule = &v
			return nil
		}
		return fmt.Errorf("Field schedule/Schedule: value %v(%T) couldn't be cast to type string", value, value)

	case "bounceAfterScale", "BounceAfterScale":
		v, ok := value.(bool)
		if ok {
			self.BounceAfterScale = &v
			return nil
		}
		return fmt.Errorf("Field bounceAfterScale/BounceAfterScale: value %v(%T) couldn't be cast to type bool", value, value)

	case "id", "Id":
		v, ok := value.(string)
		if ok {
			self.Id = &v
			return nil
		}
		return fmt.Errorf("Field id/Id: value %v(%T) couldn't be cast to type string", value, value)

	}
}

func (self *SingularityRequest) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityRequest", name)

	case "slavePlacement", "SlavePlacement":
		return *self.SlavePlacement, nil
		return nil, fmt.Errorf("Field SlavePlacement no set on SlavePlacement %+v", self)

	case "loadBalanced", "LoadBalanced":
		return *self.LoadBalanced, nil
		return nil, fmt.Errorf("Field LoadBalanced no set on LoadBalanced %+v", self)

	case "maxTasksPerOffer", "MaxTasksPerOffer":
		return *self.MaxTasksPerOffer, nil
		return nil, fmt.Errorf("Field MaxTasksPerOffer no set on MaxTasksPerOffer %+v", self)

	case "rackSensitive", "RackSensitive":
		return *self.RackSensitive, nil
		return nil, fmt.Errorf("Field RackSensitive no set on RackSensitive %+v", self)

	case "scheduleType", "ScheduleType":
		return *self.ScheduleType, nil
		return nil, fmt.Errorf("Field ScheduleType no set on ScheduleType %+v", self)

	case "killOldNonLongRunningTasksAfterMillis", "KillOldNonLongRunningTasksAfterMillis":
		return *self.KillOldNonLongRunningTasksAfterMillis, nil
		return nil, fmt.Errorf("Field KillOldNonLongRunningTasksAfterMillis no set on KillOldNonLongRunningTasksAfterMillis %+v", self)

	case "instances", "Instances":
		return *self.Instances, nil
		return nil, fmt.Errorf("Field Instances no set on Instances %+v", self)

	case "rackAffinity", "RackAffinity":
		return *self.RackAffinity, nil
		return nil, fmt.Errorf("Field RackAffinity no set on RackAffinity %+v", self)

	case "hideEvenNumberAcrossRacksHint", "HideEvenNumberAcrossRacksHint":
		return *self.HideEvenNumberAcrossRacksHint, nil
		return nil, fmt.Errorf("Field HideEvenNumberAcrossRacksHint no set on HideEvenNumberAcrossRacksHint %+v", self)

	case "allowBounceToSameHost", "AllowBounceToSameHost":
		return *self.AllowBounceToSameHost, nil
		return nil, fmt.Errorf("Field AllowBounceToSameHost no set on AllowBounceToSameHost %+v", self)

	case "numRetriesOnFailure", "NumRetriesOnFailure":
		return *self.NumRetriesOnFailure, nil
		return nil, fmt.Errorf("Field NumRetriesOnFailure no set on NumRetriesOnFailure %+v", self)

	case "taskLogErrorRegexCaseSensitive", "TaskLogErrorRegexCaseSensitive":
		return *self.TaskLogErrorRegexCaseSensitive, nil
		return nil, fmt.Errorf("Field TaskLogErrorRegexCaseSensitive no set on TaskLogErrorRegexCaseSensitive %+v", self)

	case "requestType", "RequestType":
		return *self.RequestType, nil
		return nil, fmt.Errorf("Field RequestType no set on RequestType %+v", self)

	case "readOnlyGroups", "ReadOnlyGroups":
		return *self.ReadOnlyGroups, nil
		return nil, fmt.Errorf("Field ReadOnlyGroups no set on ReadOnlyGroups %+v", self)

	case "scheduleTimeZone", "ScheduleTimeZone":
		return *self.ScheduleTimeZone, nil
		return nil, fmt.Errorf("Field ScheduleTimeZone no set on ScheduleTimeZone %+v", self)

	case "skipHealthchecks", "SkipHealthchecks":
		return *self.SkipHealthchecks, nil
		return nil, fmt.Errorf("Field SkipHealthchecks no set on SkipHealthchecks %+v", self)

	case "group", "Group":
		return *self.Group, nil
		return nil, fmt.Errorf("Field Group no set on Group %+v", self)

	case "owners", "Owners":
		return *self.Owners, nil
		return nil, fmt.Errorf("Field Owners no set on Owners %+v", self)

	case "allowedSlaveAttributes", "AllowedSlaveAttributes":
		return *self.AllowedSlaveAttributes, nil
		return nil, fmt.Errorf("Field AllowedSlaveAttributes no set on AllowedSlaveAttributes %+v", self)

	case "requiredRole", "RequiredRole":
		return *self.RequiredRole, nil
		return nil, fmt.Errorf("Field RequiredRole no set on RequiredRole %+v", self)

	case "taskExecutionTimeLimitMillis", "TaskExecutionTimeLimitMillis":
		return *self.TaskExecutionTimeLimitMillis, nil
		return nil, fmt.Errorf("Field TaskExecutionTimeLimitMillis no set on TaskExecutionTimeLimitMillis %+v", self)

	case "quartzSchedule", "QuartzSchedule":
		return *self.QuartzSchedule, nil
		return nil, fmt.Errorf("Field QuartzSchedule no set on QuartzSchedule %+v", self)

	case "scheduledExpectedRuntimeMillis", "ScheduledExpectedRuntimeMillis":
		return *self.ScheduledExpectedRuntimeMillis, nil
		return nil, fmt.Errorf("Field ScheduledExpectedRuntimeMillis no set on ScheduledExpectedRuntimeMillis %+v", self)

	case "waitAtLeastMillisAfterTaskFinishesForReschedule", "WaitAtLeastMillisAfterTaskFinishesForReschedule":
		return *self.WaitAtLeastMillisAfterTaskFinishesForReschedule, nil
		return nil, fmt.Errorf("Field WaitAtLeastMillisAfterTaskFinishesForReschedule no set on WaitAtLeastMillisAfterTaskFinishesForReschedule %+v", self)

	case "requiredSlaveAttributes", "RequiredSlaveAttributes":
		return *self.RequiredSlaveAttributes, nil
		return nil, fmt.Errorf("Field RequiredSlaveAttributes no set on RequiredSlaveAttributes %+v", self)

	case "readWriteGroups", "ReadWriteGroups":
		return *self.ReadWriteGroups, nil
		return nil, fmt.Errorf("Field ReadWriteGroups no set on ReadWriteGroups %+v", self)

	case "taskLogErrorRegex", "TaskLogErrorRegex":
		return *self.TaskLogErrorRegex, nil
		return nil, fmt.Errorf("Field TaskLogErrorRegex no set on TaskLogErrorRegex %+v", self)

	case "taskPriorityLevel", "TaskPriorityLevel":
		return *self.TaskPriorityLevel, nil
		return nil, fmt.Errorf("Field TaskPriorityLevel no set on TaskPriorityLevel %+v", self)

	case "schedule", "Schedule":
		return *self.Schedule, nil
		return nil, fmt.Errorf("Field Schedule no set on Schedule %+v", self)

	case "bounceAfterScale", "BounceAfterScale":
		return *self.BounceAfterScale, nil
		return nil, fmt.Errorf("Field BounceAfterScale no set on BounceAfterScale %+v", self)

	case "id", "Id":
		return *self.Id, nil
		return nil, fmt.Errorf("Field Id no set on Id %+v", self)

	}
}

func (self *SingularityRequest) ClearField(name string) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityRequest", name)

	case "slavePlacement", "SlavePlacement":
		self.SlavePlacement = nil

	case "loadBalanced", "LoadBalanced":
		self.LoadBalanced = nil

	case "maxTasksPerOffer", "MaxTasksPerOffer":
		self.MaxTasksPerOffer = nil

	case "rackSensitive", "RackSensitive":
		self.RackSensitive = nil

	case "scheduleType", "ScheduleType":
		self.ScheduleType = nil

	case "killOldNonLongRunningTasksAfterMillis", "KillOldNonLongRunningTasksAfterMillis":
		self.KillOldNonLongRunningTasksAfterMillis = nil

	case "instances", "Instances":
		self.Instances = nil

	case "rackAffinity", "RackAffinity":
		self.RackAffinity = nil

	case "hideEvenNumberAcrossRacksHint", "HideEvenNumberAcrossRacksHint":
		self.HideEvenNumberAcrossRacksHint = nil

	case "allowBounceToSameHost", "AllowBounceToSameHost":
		self.AllowBounceToSameHost = nil

	case "numRetriesOnFailure", "NumRetriesOnFailure":
		self.NumRetriesOnFailure = nil

	case "taskLogErrorRegexCaseSensitive", "TaskLogErrorRegexCaseSensitive":
		self.TaskLogErrorRegexCaseSensitive = nil

	case "requestType", "RequestType":
		self.RequestType = nil

	case "readOnlyGroups", "ReadOnlyGroups":
		self.ReadOnlyGroups = nil

	case "scheduleTimeZone", "ScheduleTimeZone":
		self.ScheduleTimeZone = nil

	case "skipHealthchecks", "SkipHealthchecks":
		self.SkipHealthchecks = nil

	case "group", "Group":
		self.Group = nil

	case "owners", "Owners":
		self.Owners = nil

	case "allowedSlaveAttributes", "AllowedSlaveAttributes":
		self.AllowedSlaveAttributes = nil

	case "requiredRole", "RequiredRole":
		self.RequiredRole = nil

	case "taskExecutionTimeLimitMillis", "TaskExecutionTimeLimitMillis":
		self.TaskExecutionTimeLimitMillis = nil

	case "quartzSchedule", "QuartzSchedule":
		self.QuartzSchedule = nil

	case "scheduledExpectedRuntimeMillis", "ScheduledExpectedRuntimeMillis":
		self.ScheduledExpectedRuntimeMillis = nil

	case "waitAtLeastMillisAfterTaskFinishesForReschedule", "WaitAtLeastMillisAfterTaskFinishesForReschedule":
		self.WaitAtLeastMillisAfterTaskFinishesForReschedule = nil

	case "requiredSlaveAttributes", "RequiredSlaveAttributes":
		self.RequiredSlaveAttributes = nil

	case "readWriteGroups", "ReadWriteGroups":
		self.ReadWriteGroups = nil

	case "taskLogErrorRegex", "TaskLogErrorRegex":
		self.TaskLogErrorRegex = nil

	case "taskPriorityLevel", "TaskPriorityLevel":
		self.TaskPriorityLevel = nil

	case "schedule", "Schedule":
		self.Schedule = nil

	case "bounceAfterScale", "BounceAfterScale":
		self.BounceAfterScale = nil

	case "id", "Id":
		self.Id = nil

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
