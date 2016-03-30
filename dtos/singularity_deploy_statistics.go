package dtos

import "io"

type SingularityDeployStatistics struct {
	AverageRuntimeMillis int64
	DeployId             string
	//	InstanceSequentialFailureTimestamps *com.google.common.collect.ListMultimap&lt;java.lang.Integer, java.lang.Long&gt;
	LastFinishAt int64
	//	LastTaskState *ExtendedTaskState
	NumFailures          int32
	NumSequentialRetries int32
	NumSuccess           int32
	NumTasks             int32
	RequestId            string
}

func (self *SingularityDeployStatistics) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SingularityDeployStatistics) FormatText() string {
	return FormatText(self)
}

func (self *SingularityDeployStatistics) FormatJSON() string {
	return FormatJSON(self)
}
