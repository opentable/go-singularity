package dtos

import "io"

type SingularityState struct {
	ActiveRacks                int32
	ActiveRequests             int32
	ActiveSlaves               int32
	ActiveTasks                int32
	AllRequests                int32
	AuthDatastoreHealthy       bool
	CleaningRequests           int32
	CleaningTasks              int32
	CooldownRequests           int32
	DeadRacks                  int32
	DeadSlaves                 int32
	DecomissioningRacks        int32
	DecomissioningSlaves       int32
	DecommissioningRacks       int32
	DecommissioningSlaves      int32
	FinishedRequests           int32
	FutureTasks                int32
	GeneratedAt                int64
	HostStates                 SingularityHostStateList
	LateTasks                  int32
	LbCleanupRequests          int32
	LbCleanupTasks             int32
	MaxTaskLag                 int64
	NumDeploys                 int32
	OldestDeploy               int64
	OverProvisionedRequestIds  StringList
	OverProvisionedRequests    int32
	PausedRequests             int32
	PendingRequests            int32
	ScheduledTasks             int32
	UnderProvisionedRequestIds StringList
	UnderProvisionedRequests   int32
	UnknownRacks               int32
	UnknownSlaves              int32
}

func (self *SingularityState) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SingularityState) FormatText() string {
	return FormatText(self)
}

func (self *SingularityState) FormatJSON() string {
	return FormatJSON(self)
}

type SingularityStateList []*SingularityState

func (list SingularityStateList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list SingularityStateList) FormatText() string {
	text := []byte{}
	for _, dto := range list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list SingularityStateList) FormatJSON() string {
	return FormatJSON(list)
}
