package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityState struct {
	LateTasks                  *int32                       `json:"lateTasks,omitempty"`
	PendingRequests            *int32                       `json:"pendingRequests,omitempty"`
	FinishedRequests           *int32                       `json:"finishedRequests,omitempty"`
	UnknownSlaves              *int32                       `json:"unknownSlaves,omitempty"`
	PausedRequests             *int32                       `json:"pausedRequests,omitempty"`
	DecomissioningRacks        *int32                       `json:"decomissioningRacks,omitempty"`
	ScheduledTasks             *int32                       `json:"scheduledTasks,omitempty"`
	MinimumPriorityLevel       *float64                     `json:"minimumPriorityLevel,omitempty"`
	OldestDeployStep           *int64                       `json:"oldestDeployStep,omitempty"`
	LbCleanupRequests          *int32                       `json:"lbCleanupRequests,omitempty"`
	DecommissioningSlaves      *int32                       `json:"decommissioningSlaves,omitempty"`
	OldestDeploy               *int64                       `json:"oldestDeploy,omitempty"`
	ActiveDeploys              *SingularityDeployMarkerList `json:"activeDeploys,omitempty"`
	HostStates                 *SingularityHostStateList    `json:"hostStates,omitempty"`
	UnderProvisionedRequestIds *swaggering.StringList       `json:"underProvisionedRequestIds,omitempty"`
	UnderProvisionedRequests   *int32                       `json:"underProvisionedRequests,omitempty"`
	LaunchingTasks             *int32                       `json:"launchingTasks,omitempty"`
	DecomissioningSlaves       *int32                       `json:"decomissioningSlaves,omitempty"`
	LbCleanupTasks             *int32                       `json:"lbCleanupTasks,omitempty"`
	MaxTaskLag                 *int64                       `json:"maxTaskLag,omitempty"`
	UnknownRacks               *int32                       `json:"unknownRacks,omitempty"`
	NumDeploys                 *int32                       `json:"numDeploys,omitempty"`
	AllRequests                *int32                       `json:"allRequests,omitempty"`
	ActiveRequests             *int32                       `json:"activeRequests,omitempty"`
	ActiveSlaves               *int32                       `json:"activeSlaves,omitempty"`
	OverProvisionedRequestIds  *swaggering.StringList       `json:"overProvisionedRequestIds,omitempty"`
	OverProvisionedRequests    *int32                       `json:"overProvisionedRequests,omitempty"`
	AuthDatastoreHealthy       *bool                        `json:"authDatastoreHealthy,omitempty"`
	CleaningTasks              *int32                       `json:"cleaningTasks,omitempty"`
	DeadSlaves                 *int32                       `json:"deadSlaves,omitempty"`
	ActiveRacks                *int32                       `json:"activeRacks,omitempty"`
	DeadRacks                  *int32                       `json:"deadRacks,omitempty"`
	DecommissioningRacks       *int32                       `json:"decommissioningRacks,omitempty"`
	AvgStatusUpdateDelayMs     *int64                       `json:"avgStatusUpdateDelayMs,omitempty"`
	CooldownRequests           *int32                       `json:"cooldownRequests,omitempty"`
	FutureTasks                *int32                       `json:"futureTasks,omitempty"`
	CleaningRequests           *int32                       `json:"cleaningRequests,omitempty"`
	GeneratedAt                *int64                       `json:"generatedAt,omitempty"`
	ActiveTasks                *int32                       `json:"activeTasks,omitempty"`
}

func (self *SingularityState) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *SingularityState) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityState); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityState cannot copy the values from %#v", other)
}

func (self *SingularityState) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityState) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityState) SetField(name string, value interface{}) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityState", name)

	case "lateTasks", "LateTasks":
		v, ok := value.(int32)
		if ok {
			self.LateTasks = &v
			return nil
		}
		return fmt.Errorf("Field lateTasks/LateTasks: value %v(%T) couldn't be cast to type int32", value, value)

	case "pendingRequests", "PendingRequests":
		v, ok := value.(int32)
		if ok {
			self.PendingRequests = &v
			return nil
		}
		return fmt.Errorf("Field pendingRequests/PendingRequests: value %v(%T) couldn't be cast to type int32", value, value)

	case "finishedRequests", "FinishedRequests":
		v, ok := value.(int32)
		if ok {
			self.FinishedRequests = &v
			return nil
		}
		return fmt.Errorf("Field finishedRequests/FinishedRequests: value %v(%T) couldn't be cast to type int32", value, value)

	case "unknownSlaves", "UnknownSlaves":
		v, ok := value.(int32)
		if ok {
			self.UnknownSlaves = &v
			return nil
		}
		return fmt.Errorf("Field unknownSlaves/UnknownSlaves: value %v(%T) couldn't be cast to type int32", value, value)

	case "pausedRequests", "PausedRequests":
		v, ok := value.(int32)
		if ok {
			self.PausedRequests = &v
			return nil
		}
		return fmt.Errorf("Field pausedRequests/PausedRequests: value %v(%T) couldn't be cast to type int32", value, value)

	case "decomissioningRacks", "DecomissioningRacks":
		v, ok := value.(int32)
		if ok {
			self.DecomissioningRacks = &v
			return nil
		}
		return fmt.Errorf("Field decomissioningRacks/DecomissioningRacks: value %v(%T) couldn't be cast to type int32", value, value)

	case "scheduledTasks", "ScheduledTasks":
		v, ok := value.(int32)
		if ok {
			self.ScheduledTasks = &v
			return nil
		}
		return fmt.Errorf("Field scheduledTasks/ScheduledTasks: value %v(%T) couldn't be cast to type int32", value, value)

	case "minimumPriorityLevel", "MinimumPriorityLevel":
		v, ok := value.(float64)
		if ok {
			self.MinimumPriorityLevel = &v
			return nil
		}
		return fmt.Errorf("Field minimumPriorityLevel/MinimumPriorityLevel: value %v(%T) couldn't be cast to type float64", value, value)

	case "oldestDeployStep", "OldestDeployStep":
		v, ok := value.(int64)
		if ok {
			self.OldestDeployStep = &v
			return nil
		}
		return fmt.Errorf("Field oldestDeployStep/OldestDeployStep: value %v(%T) couldn't be cast to type int64", value, value)

	case "lbCleanupRequests", "LbCleanupRequests":
		v, ok := value.(int32)
		if ok {
			self.LbCleanupRequests = &v
			return nil
		}
		return fmt.Errorf("Field lbCleanupRequests/LbCleanupRequests: value %v(%T) couldn't be cast to type int32", value, value)

	case "decommissioningSlaves", "DecommissioningSlaves":
		v, ok := value.(int32)
		if ok {
			self.DecommissioningSlaves = &v
			return nil
		}
		return fmt.Errorf("Field decommissioningSlaves/DecommissioningSlaves: value %v(%T) couldn't be cast to type int32", value, value)

	case "oldestDeploy", "OldestDeploy":
		v, ok := value.(int64)
		if ok {
			self.OldestDeploy = &v
			return nil
		}
		return fmt.Errorf("Field oldestDeploy/OldestDeploy: value %v(%T) couldn't be cast to type int64", value, value)

	case "activeDeploys", "ActiveDeploys":
		v, ok := value.(SingularityDeployMarkerList)
		if ok {
			self.ActiveDeploys = &v
			return nil
		}
		return fmt.Errorf("Field activeDeploys/ActiveDeploys: value %v(%T) couldn't be cast to type SingularityDeployMarkerList", value, value)

	case "hostStates", "HostStates":
		v, ok := value.(SingularityHostStateList)
		if ok {
			self.HostStates = &v
			return nil
		}
		return fmt.Errorf("Field hostStates/HostStates: value %v(%T) couldn't be cast to type SingularityHostStateList", value, value)

	case "underProvisionedRequestIds", "UnderProvisionedRequestIds":
		v, ok := value.(swaggering.StringList)
		if ok {
			self.UnderProvisionedRequestIds = &v
			return nil
		}
		return fmt.Errorf("Field underProvisionedRequestIds/UnderProvisionedRequestIds: value %v(%T) couldn't be cast to type swaggering.StringList", value, value)

	case "underProvisionedRequests", "UnderProvisionedRequests":
		v, ok := value.(int32)
		if ok {
			self.UnderProvisionedRequests = &v
			return nil
		}
		return fmt.Errorf("Field underProvisionedRequests/UnderProvisionedRequests: value %v(%T) couldn't be cast to type int32", value, value)

	case "launchingTasks", "LaunchingTasks":
		v, ok := value.(int32)
		if ok {
			self.LaunchingTasks = &v
			return nil
		}
		return fmt.Errorf("Field launchingTasks/LaunchingTasks: value %v(%T) couldn't be cast to type int32", value, value)

	case "decomissioningSlaves", "DecomissioningSlaves":
		v, ok := value.(int32)
		if ok {
			self.DecomissioningSlaves = &v
			return nil
		}
		return fmt.Errorf("Field decomissioningSlaves/DecomissioningSlaves: value %v(%T) couldn't be cast to type int32", value, value)

	case "lbCleanupTasks", "LbCleanupTasks":
		v, ok := value.(int32)
		if ok {
			self.LbCleanupTasks = &v
			return nil
		}
		return fmt.Errorf("Field lbCleanupTasks/LbCleanupTasks: value %v(%T) couldn't be cast to type int32", value, value)

	case "maxTaskLag", "MaxTaskLag":
		v, ok := value.(int64)
		if ok {
			self.MaxTaskLag = &v
			return nil
		}
		return fmt.Errorf("Field maxTaskLag/MaxTaskLag: value %v(%T) couldn't be cast to type int64", value, value)

	case "unknownRacks", "UnknownRacks":
		v, ok := value.(int32)
		if ok {
			self.UnknownRacks = &v
			return nil
		}
		return fmt.Errorf("Field unknownRacks/UnknownRacks: value %v(%T) couldn't be cast to type int32", value, value)

	case "numDeploys", "NumDeploys":
		v, ok := value.(int32)
		if ok {
			self.NumDeploys = &v
			return nil
		}
		return fmt.Errorf("Field numDeploys/NumDeploys: value %v(%T) couldn't be cast to type int32", value, value)

	case "allRequests", "AllRequests":
		v, ok := value.(int32)
		if ok {
			self.AllRequests = &v
			return nil
		}
		return fmt.Errorf("Field allRequests/AllRequests: value %v(%T) couldn't be cast to type int32", value, value)

	case "activeRequests", "ActiveRequests":
		v, ok := value.(int32)
		if ok {
			self.ActiveRequests = &v
			return nil
		}
		return fmt.Errorf("Field activeRequests/ActiveRequests: value %v(%T) couldn't be cast to type int32", value, value)

	case "activeSlaves", "ActiveSlaves":
		v, ok := value.(int32)
		if ok {
			self.ActiveSlaves = &v
			return nil
		}
		return fmt.Errorf("Field activeSlaves/ActiveSlaves: value %v(%T) couldn't be cast to type int32", value, value)

	case "overProvisionedRequestIds", "OverProvisionedRequestIds":
		v, ok := value.(swaggering.StringList)
		if ok {
			self.OverProvisionedRequestIds = &v
			return nil
		}
		return fmt.Errorf("Field overProvisionedRequestIds/OverProvisionedRequestIds: value %v(%T) couldn't be cast to type swaggering.StringList", value, value)

	case "overProvisionedRequests", "OverProvisionedRequests":
		v, ok := value.(int32)
		if ok {
			self.OverProvisionedRequests = &v
			return nil
		}
		return fmt.Errorf("Field overProvisionedRequests/OverProvisionedRequests: value %v(%T) couldn't be cast to type int32", value, value)

	case "authDatastoreHealthy", "AuthDatastoreHealthy":
		v, ok := value.(bool)
		if ok {
			self.AuthDatastoreHealthy = &v
			return nil
		}
		return fmt.Errorf("Field authDatastoreHealthy/AuthDatastoreHealthy: value %v(%T) couldn't be cast to type bool", value, value)

	case "cleaningTasks", "CleaningTasks":
		v, ok := value.(int32)
		if ok {
			self.CleaningTasks = &v
			return nil
		}
		return fmt.Errorf("Field cleaningTasks/CleaningTasks: value %v(%T) couldn't be cast to type int32", value, value)

	case "deadSlaves", "DeadSlaves":
		v, ok := value.(int32)
		if ok {
			self.DeadSlaves = &v
			return nil
		}
		return fmt.Errorf("Field deadSlaves/DeadSlaves: value %v(%T) couldn't be cast to type int32", value, value)

	case "activeRacks", "ActiveRacks":
		v, ok := value.(int32)
		if ok {
			self.ActiveRacks = &v
			return nil
		}
		return fmt.Errorf("Field activeRacks/ActiveRacks: value %v(%T) couldn't be cast to type int32", value, value)

	case "deadRacks", "DeadRacks":
		v, ok := value.(int32)
		if ok {
			self.DeadRacks = &v
			return nil
		}
		return fmt.Errorf("Field deadRacks/DeadRacks: value %v(%T) couldn't be cast to type int32", value, value)

	case "decommissioningRacks", "DecommissioningRacks":
		v, ok := value.(int32)
		if ok {
			self.DecommissioningRacks = &v
			return nil
		}
		return fmt.Errorf("Field decommissioningRacks/DecommissioningRacks: value %v(%T) couldn't be cast to type int32", value, value)

	case "avgStatusUpdateDelayMs", "AvgStatusUpdateDelayMs":
		v, ok := value.(int64)
		if ok {
			self.AvgStatusUpdateDelayMs = &v
			return nil
		}
		return fmt.Errorf("Field avgStatusUpdateDelayMs/AvgStatusUpdateDelayMs: value %v(%T) couldn't be cast to type int64", value, value)

	case "cooldownRequests", "CooldownRequests":
		v, ok := value.(int32)
		if ok {
			self.CooldownRequests = &v
			return nil
		}
		return fmt.Errorf("Field cooldownRequests/CooldownRequests: value %v(%T) couldn't be cast to type int32", value, value)

	case "futureTasks", "FutureTasks":
		v, ok := value.(int32)
		if ok {
			self.FutureTasks = &v
			return nil
		}
		return fmt.Errorf("Field futureTasks/FutureTasks: value %v(%T) couldn't be cast to type int32", value, value)

	case "cleaningRequests", "CleaningRequests":
		v, ok := value.(int32)
		if ok {
			self.CleaningRequests = &v
			return nil
		}
		return fmt.Errorf("Field cleaningRequests/CleaningRequests: value %v(%T) couldn't be cast to type int32", value, value)

	case "generatedAt", "GeneratedAt":
		v, ok := value.(int64)
		if ok {
			self.GeneratedAt = &v
			return nil
		}
		return fmt.Errorf("Field generatedAt/GeneratedAt: value %v(%T) couldn't be cast to type int64", value, value)

	case "activeTasks", "ActiveTasks":
		v, ok := value.(int32)
		if ok {
			self.ActiveTasks = &v
			return nil
		}
		return fmt.Errorf("Field activeTasks/ActiveTasks: value %v(%T) couldn't be cast to type int32", value, value)

	}
}

func (self *SingularityState) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityState", name)

	case "lateTasks", "LateTasks":
		return *self.LateTasks, nil
		return nil, fmt.Errorf("Field LateTasks no set on LateTasks %+v", self)

	case "pendingRequests", "PendingRequests":
		return *self.PendingRequests, nil
		return nil, fmt.Errorf("Field PendingRequests no set on PendingRequests %+v", self)

	case "finishedRequests", "FinishedRequests":
		return *self.FinishedRequests, nil
		return nil, fmt.Errorf("Field FinishedRequests no set on FinishedRequests %+v", self)

	case "unknownSlaves", "UnknownSlaves":
		return *self.UnknownSlaves, nil
		return nil, fmt.Errorf("Field UnknownSlaves no set on UnknownSlaves %+v", self)

	case "pausedRequests", "PausedRequests":
		return *self.PausedRequests, nil
		return nil, fmt.Errorf("Field PausedRequests no set on PausedRequests %+v", self)

	case "decomissioningRacks", "DecomissioningRacks":
		return *self.DecomissioningRacks, nil
		return nil, fmt.Errorf("Field DecomissioningRacks no set on DecomissioningRacks %+v", self)

	case "scheduledTasks", "ScheduledTasks":
		return *self.ScheduledTasks, nil
		return nil, fmt.Errorf("Field ScheduledTasks no set on ScheduledTasks %+v", self)

	case "minimumPriorityLevel", "MinimumPriorityLevel":
		return *self.MinimumPriorityLevel, nil
		return nil, fmt.Errorf("Field MinimumPriorityLevel no set on MinimumPriorityLevel %+v", self)

	case "oldestDeployStep", "OldestDeployStep":
		return *self.OldestDeployStep, nil
		return nil, fmt.Errorf("Field OldestDeployStep no set on OldestDeployStep %+v", self)

	case "lbCleanupRequests", "LbCleanupRequests":
		return *self.LbCleanupRequests, nil
		return nil, fmt.Errorf("Field LbCleanupRequests no set on LbCleanupRequests %+v", self)

	case "decommissioningSlaves", "DecommissioningSlaves":
		return *self.DecommissioningSlaves, nil
		return nil, fmt.Errorf("Field DecommissioningSlaves no set on DecommissioningSlaves %+v", self)

	case "oldestDeploy", "OldestDeploy":
		return *self.OldestDeploy, nil
		return nil, fmt.Errorf("Field OldestDeploy no set on OldestDeploy %+v", self)

	case "activeDeploys", "ActiveDeploys":
		return *self.ActiveDeploys, nil
		return nil, fmt.Errorf("Field ActiveDeploys no set on ActiveDeploys %+v", self)

	case "hostStates", "HostStates":
		return *self.HostStates, nil
		return nil, fmt.Errorf("Field HostStates no set on HostStates %+v", self)

	case "underProvisionedRequestIds", "UnderProvisionedRequestIds":
		return *self.UnderProvisionedRequestIds, nil
		return nil, fmt.Errorf("Field UnderProvisionedRequestIds no set on UnderProvisionedRequestIds %+v", self)

	case "underProvisionedRequests", "UnderProvisionedRequests":
		return *self.UnderProvisionedRequests, nil
		return nil, fmt.Errorf("Field UnderProvisionedRequests no set on UnderProvisionedRequests %+v", self)

	case "launchingTasks", "LaunchingTasks":
		return *self.LaunchingTasks, nil
		return nil, fmt.Errorf("Field LaunchingTasks no set on LaunchingTasks %+v", self)

	case "decomissioningSlaves", "DecomissioningSlaves":
		return *self.DecomissioningSlaves, nil
		return nil, fmt.Errorf("Field DecomissioningSlaves no set on DecomissioningSlaves %+v", self)

	case "lbCleanupTasks", "LbCleanupTasks":
		return *self.LbCleanupTasks, nil
		return nil, fmt.Errorf("Field LbCleanupTasks no set on LbCleanupTasks %+v", self)

	case "maxTaskLag", "MaxTaskLag":
		return *self.MaxTaskLag, nil
		return nil, fmt.Errorf("Field MaxTaskLag no set on MaxTaskLag %+v", self)

	case "unknownRacks", "UnknownRacks":
		return *self.UnknownRacks, nil
		return nil, fmt.Errorf("Field UnknownRacks no set on UnknownRacks %+v", self)

	case "numDeploys", "NumDeploys":
		return *self.NumDeploys, nil
		return nil, fmt.Errorf("Field NumDeploys no set on NumDeploys %+v", self)

	case "allRequests", "AllRequests":
		return *self.AllRequests, nil
		return nil, fmt.Errorf("Field AllRequests no set on AllRequests %+v", self)

	case "activeRequests", "ActiveRequests":
		return *self.ActiveRequests, nil
		return nil, fmt.Errorf("Field ActiveRequests no set on ActiveRequests %+v", self)

	case "activeSlaves", "ActiveSlaves":
		return *self.ActiveSlaves, nil
		return nil, fmt.Errorf("Field ActiveSlaves no set on ActiveSlaves %+v", self)

	case "overProvisionedRequestIds", "OverProvisionedRequestIds":
		return *self.OverProvisionedRequestIds, nil
		return nil, fmt.Errorf("Field OverProvisionedRequestIds no set on OverProvisionedRequestIds %+v", self)

	case "overProvisionedRequests", "OverProvisionedRequests":
		return *self.OverProvisionedRequests, nil
		return nil, fmt.Errorf("Field OverProvisionedRequests no set on OverProvisionedRequests %+v", self)

	case "authDatastoreHealthy", "AuthDatastoreHealthy":
		return *self.AuthDatastoreHealthy, nil
		return nil, fmt.Errorf("Field AuthDatastoreHealthy no set on AuthDatastoreHealthy %+v", self)

	case "cleaningTasks", "CleaningTasks":
		return *self.CleaningTasks, nil
		return nil, fmt.Errorf("Field CleaningTasks no set on CleaningTasks %+v", self)

	case "deadSlaves", "DeadSlaves":
		return *self.DeadSlaves, nil
		return nil, fmt.Errorf("Field DeadSlaves no set on DeadSlaves %+v", self)

	case "activeRacks", "ActiveRacks":
		return *self.ActiveRacks, nil
		return nil, fmt.Errorf("Field ActiveRacks no set on ActiveRacks %+v", self)

	case "deadRacks", "DeadRacks":
		return *self.DeadRacks, nil
		return nil, fmt.Errorf("Field DeadRacks no set on DeadRacks %+v", self)

	case "decommissioningRacks", "DecommissioningRacks":
		return *self.DecommissioningRacks, nil
		return nil, fmt.Errorf("Field DecommissioningRacks no set on DecommissioningRacks %+v", self)

	case "avgStatusUpdateDelayMs", "AvgStatusUpdateDelayMs":
		return *self.AvgStatusUpdateDelayMs, nil
		return nil, fmt.Errorf("Field AvgStatusUpdateDelayMs no set on AvgStatusUpdateDelayMs %+v", self)

	case "cooldownRequests", "CooldownRequests":
		return *self.CooldownRequests, nil
		return nil, fmt.Errorf("Field CooldownRequests no set on CooldownRequests %+v", self)

	case "futureTasks", "FutureTasks":
		return *self.FutureTasks, nil
		return nil, fmt.Errorf("Field FutureTasks no set on FutureTasks %+v", self)

	case "cleaningRequests", "CleaningRequests":
		return *self.CleaningRequests, nil
		return nil, fmt.Errorf("Field CleaningRequests no set on CleaningRequests %+v", self)

	case "generatedAt", "GeneratedAt":
		return *self.GeneratedAt, nil
		return nil, fmt.Errorf("Field GeneratedAt no set on GeneratedAt %+v", self)

	case "activeTasks", "ActiveTasks":
		return *self.ActiveTasks, nil
		return nil, fmt.Errorf("Field ActiveTasks no set on ActiveTasks %+v", self)

	}
}

func (self *SingularityState) ClearField(name string) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityState", name)

	case "lateTasks", "LateTasks":
		self.LateTasks = nil

	case "pendingRequests", "PendingRequests":
		self.PendingRequests = nil

	case "finishedRequests", "FinishedRequests":
		self.FinishedRequests = nil

	case "unknownSlaves", "UnknownSlaves":
		self.UnknownSlaves = nil

	case "pausedRequests", "PausedRequests":
		self.PausedRequests = nil

	case "decomissioningRacks", "DecomissioningRacks":
		self.DecomissioningRacks = nil

	case "scheduledTasks", "ScheduledTasks":
		self.ScheduledTasks = nil

	case "minimumPriorityLevel", "MinimumPriorityLevel":
		self.MinimumPriorityLevel = nil

	case "oldestDeployStep", "OldestDeployStep":
		self.OldestDeployStep = nil

	case "lbCleanupRequests", "LbCleanupRequests":
		self.LbCleanupRequests = nil

	case "decommissioningSlaves", "DecommissioningSlaves":
		self.DecommissioningSlaves = nil

	case "oldestDeploy", "OldestDeploy":
		self.OldestDeploy = nil

	case "activeDeploys", "ActiveDeploys":
		self.ActiveDeploys = nil

	case "hostStates", "HostStates":
		self.HostStates = nil

	case "underProvisionedRequestIds", "UnderProvisionedRequestIds":
		self.UnderProvisionedRequestIds = nil

	case "underProvisionedRequests", "UnderProvisionedRequests":
		self.UnderProvisionedRequests = nil

	case "launchingTasks", "LaunchingTasks":
		self.LaunchingTasks = nil

	case "decomissioningSlaves", "DecomissioningSlaves":
		self.DecomissioningSlaves = nil

	case "lbCleanupTasks", "LbCleanupTasks":
		self.LbCleanupTasks = nil

	case "maxTaskLag", "MaxTaskLag":
		self.MaxTaskLag = nil

	case "unknownRacks", "UnknownRacks":
		self.UnknownRacks = nil

	case "numDeploys", "NumDeploys":
		self.NumDeploys = nil

	case "allRequests", "AllRequests":
		self.AllRequests = nil

	case "activeRequests", "ActiveRequests":
		self.ActiveRequests = nil

	case "activeSlaves", "ActiveSlaves":
		self.ActiveSlaves = nil

	case "overProvisionedRequestIds", "OverProvisionedRequestIds":
		self.OverProvisionedRequestIds = nil

	case "overProvisionedRequests", "OverProvisionedRequests":
		self.OverProvisionedRequests = nil

	case "authDatastoreHealthy", "AuthDatastoreHealthy":
		self.AuthDatastoreHealthy = nil

	case "cleaningTasks", "CleaningTasks":
		self.CleaningTasks = nil

	case "deadSlaves", "DeadSlaves":
		self.DeadSlaves = nil

	case "activeRacks", "ActiveRacks":
		self.ActiveRacks = nil

	case "deadRacks", "DeadRacks":
		self.DeadRacks = nil

	case "decommissioningRacks", "DecommissioningRacks":
		self.DecommissioningRacks = nil

	case "avgStatusUpdateDelayMs", "AvgStatusUpdateDelayMs":
		self.AvgStatusUpdateDelayMs = nil

	case "cooldownRequests", "CooldownRequests":
		self.CooldownRequests = nil

	case "futureTasks", "FutureTasks":
		self.FutureTasks = nil

	case "cleaningRequests", "CleaningRequests":
		self.CleaningRequests = nil

	case "generatedAt", "GeneratedAt":
		self.GeneratedAt = nil

	case "activeTasks", "ActiveTasks":
		self.ActiveTasks = nil

	}

	return nil
}

func (self *SingularityState) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type SingularityStateList []*SingularityState

func (self *SingularityStateList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityStateList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityStateList cannot copy the values from %#v", other)
}

func (list *SingularityStateList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
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
	return swaggering.FormatJSON(list)
}
