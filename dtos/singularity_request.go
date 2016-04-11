package dtos

import "io"

type SingularityRequest struct {
	//	AllowedSlaveAttributes *Map[string,string]
	BounceAfterScale bool
	//	EmailConfigurationOverrides *Map[SingularityEmailType,List[SingularityEmailDestination]]
	Group                                 string
	Id                                    string
	Instances                             int32
	KillOldNonLongRunningTasksAfterMillis int64
	LoadBalanced                          bool
	NumRetriesOnFailure                   int32
	Owners                                StringList
	QuartzSchedule                        string
	RackAffinity                          StringList
	RackSensitive                         bool
	ReadOnlyGroups                        StringList
	//	RequestType *RequestType
	//	RequiredSlaveAttributes *Map[string,string]
	Schedule string
	//	ScheduleType *ScheduleType
	ScheduledExpectedRuntimeMillis int64
	SkipHealthchecks               bool
	//	SlavePlacement *SlavePlacement
	WaitAtLeastMillisAfterTaskFinishesForReschedule int64
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
