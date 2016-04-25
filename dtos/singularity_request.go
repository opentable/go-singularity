package dtos

import "io"

type SingularityRequestRequestType string

const (
	SingularityRequestRequestTypeSERVICE   SingularityRequestRequestType = "SERVICE"
	SingularityRequestRequestTypeWORKER    SingularityRequestRequestType = "WORKER"
	SingularityRequestRequestTypeSCHEDULED SingularityRequestRequestType = "SCHEDULED"
	SingularityRequestRequestTypeON_DEMAND SingularityRequestRequestType = "ON_DEMAND"
	SingularityRequestRequestTypeRUN_ONCE  SingularityRequestRequestType = "RUN_ONCE"
)

type SingularityRequest struct {
	AllowedSlaveAttributes map[string]string `json:"allowedSlaveAttributes"`
	BounceAfterScale       bool              `json:"bounceAfterScale"`
	//	EmailConfigurationOverrides *Map[SingularityEmailType,List[SingularityEmailDestination]] `json:"emailConfigurationOverrides"`
	Group                                 string                        `json:"group"`
	Id                                    string                        `json:"id"`
	Instances                             int32                         `json:"instances"`
	KillOldNonLongRunningTasksAfterMillis int64                         `json:"killOldNonLongRunningTasksAfterMillis"`
	LoadBalanced                          bool                          `json:"loadBalanced"`
	NumRetriesOnFailure                   int32                         `json:"numRetriesOnFailure"`
	Owners                                StringList                    `json:"owners"`
	QuartzSchedule                        string                        `json:"quartzSchedule"`
	RackAffinity                          StringList                    `json:"rackAffinity"`
	RackSensitive                         bool                          `json:"rackSensitive"`
	ReadOnlyGroups                        StringList                    `json:"readOnlyGroups"`
	RequestType                           SingularityRequestRequestType `json:"requestType"`
	RequiredSlaveAttributes               map[string]string             `json:"requiredSlaveAttributes"`
	Schedule                              string                        `json:"schedule"`
	//	ScheduleType *ScheduleType `json:"scheduleType"`
	ScheduledExpectedRuntimeMillis int64 `json:"scheduledExpectedRuntimeMillis"`
	SkipHealthchecks               bool  `json:"skipHealthchecks"`
	//	SlavePlacement *SlavePlacement `json:"slavePlacement"`
	WaitAtLeastMillisAfterTaskFinishesForReschedule int64 `json:"waitAtLeastMillisAfterTaskFinishesForReschedule"`
}

func (self *SingularityRequest) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SingularityRequest) FormatText() string {
	return FormatText(self)
}

func (self *SingularityRequest) FormatJSON() string {
	return FormatJSON(self)
}

type SingularityRequestList []*SingularityRequest

func (list *SingularityRequestList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
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
	return FormatJSON(list)
}
