package dtos

import "io"

type Request struct {
	//ReadOnlyGroups [Set]
	//SlavePlacement [SlavePlacement]
	//EmailConfigurationOverrides [Map[SingularityEmailType,List[SingularityEmailDestination]]]
	//RequestType [RequestType]
	//ScheduleType [ScheduleType]
	Id        string
	Owners    []string
	Instances int

	Schedule                                        string
	QuartzSchedule                                  string
	WaitAtLeastMillisAfterTaskFinishesForReschedule uint64

	LoadBalanced                   bool
	SkipHealthchecks               bool
	ScheduledExpectedRuntimeMillis uint64

	Group         string
	RackAffinity  []string
	RackSensitive bool

	BounceAfterScale                      bool
	NumRetriesOnFailure                   int
	KillOldNonLongRunningTasksAfterMillis uint64

	AllowedSlaveAttributes  map[string]string
	RequiredSlaveAttributes map[string]string
}

type Requests []Request

func (reqs *Requests) Get(client APIClient) (err error) {
	return client.APIGet("requests", reqs)
}

func (reqs *Requests) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, reqs)
	return
}

func (reqs *Requests) FormatText() string {
	return FormatText(reqs)
}

func (reqs *Requests) FormatJSON() string {
	return FormatJSON(reqs)
}
