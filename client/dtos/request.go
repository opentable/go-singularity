package dtos

import "io"

type Request struct {
//	AllowedSlaveAttributes Map[string,string]
	BounceAfterScale bool
//	EmailConfigurationOverrides Map[SingularityEmailType,List[SingularityEmailDestination]]
	Group string
	Id string
	Instances int32
	KillOldNonLongRunningTasksAfterMillis int64
	LoadBalanced bool
	NumRetriesOnFailure int32
	Owners []string
	QuartzSchedule string
	RackAffinity []string
	RackSensitive bool
	ReadOnlyGroups []string
//	RequestType RequestType
//	RequiredSlaveAttributes Map[string,string]
	Schedule string
//	ScheduleType ScheduleType
	ScheduledExpectedRuntimeMillis int64
	SkipHealthchecks bool
//	SlavePlacement SlavePlacement
	WaitAtLeastMillisAfterTaskFinishesForReschedule int64
}

func (self *Request) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *Request) FormatText() string {
	return FormatText(self)
}

func (self *Request) FormatJSON() string {
	return FormatJSON(self)
}
