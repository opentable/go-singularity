package dtos

import "io"

type SingularityState struct {
	ActiveRacks                int32                    `json:"activeRacks"`
	ActiveRequests             int32                    `json:"activeRequests"`
	ActiveSlaves               int32                    `json:"activeSlaves"`
	ActiveTasks                int32                    `json:"activeTasks"`
	AllRequests                int32                    `json:"allRequests"`
	AuthDatastoreHealthy       bool                     `json:"authDatastoreHealthy"`
	CleaningRequests           int32                    `json:"cleaningRequests"`
	CleaningTasks              int32                    `json:"cleaningTasks"`
	CooldownRequests           int32                    `json:"cooldownRequests"`
	DeadRacks                  int32                    `json:"deadRacks"`
	DeadSlaves                 int32                    `json:"deadSlaves"`
	DecomissioningRacks        int32                    `json:"decomissioningRacks"`
	DecomissioningSlaves       int32                    `json:"decomissioningSlaves"`
	DecommissioningRacks       int32                    `json:"decommissioningRacks"`
	DecommissioningSlaves      int32                    `json:"decommissioningSlaves"`
	FinishedRequests           int32                    `json:"finishedRequests"`
	FutureTasks                int32                    `json:"futureTasks"`
	GeneratedAt                int64                    `json:"generatedAt"`
	HostStates                 SingularityHostStateList `json:"hostStates"`
	LateTasks                  int32                    `json:"lateTasks"`
	LbCleanupRequests          int32                    `json:"lbCleanupRequests"`
	LbCleanupTasks             int32                    `json:"lbCleanupTasks"`
	MaxTaskLag                 int64                    `json:"maxTaskLag"`
	NumDeploys                 int32                    `json:"numDeploys"`
	OldestDeploy               int64                    `json:"oldestDeploy"`
	OverProvisionedRequestIds  StringList               `json:"overProvisionedRequestIds"`
	OverProvisionedRequests    int32                    `json:"overProvisionedRequests"`
	PausedRequests             int32                    `json:"pausedRequests"`
	PendingRequests            int32                    `json:"pendingRequests"`
	ScheduledTasks             int32                    `json:"scheduledTasks"`
	UnderProvisionedRequestIds StringList               `json:"underProvisionedRequestIds"`
	UnderProvisionedRequests   int32                    `json:"underProvisionedRequests"`
	UnknownRacks               int32                    `json:"unknownRacks"`
	UnknownSlaves              int32                    `json:"unknownSlaves"`
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

func (list *SingularityStateList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *SingularityStateList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityStateList) FormatJSON() string {
	return FormatJSON(list)
}
