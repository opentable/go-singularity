package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityState struct {
	present map[string]bool

	ActiveRacks int32 `json:"activeRacks"`

	ActiveRequests int32 `json:"activeRequests"`

	ActiveSlaves int32 `json:"activeSlaves"`

	ActiveTasks int32 `json:"activeTasks"`

	AllRequests int32 `json:"allRequests"`

	AuthDatastoreHealthy bool `json:"authDatastoreHealthy"`

	CleaningRequests int32 `json:"cleaningRequests"`

	CleaningTasks int32 `json:"cleaningTasks"`

	CooldownRequests int32 `json:"cooldownRequests"`

	DeadRacks int32 `json:"deadRacks"`

	DeadSlaves int32 `json:"deadSlaves"`

	DecomissioningRacks int32 `json:"decomissioningRacks"`

	DecomissioningSlaves int32 `json:"decomissioningSlaves"`

	DecommissioningRacks int32 `json:"decommissioningRacks"`

	DecommissioningSlaves int32 `json:"decommissioningSlaves"`

	FinishedRequests int32 `json:"finishedRequests"`

	FutureTasks int32 `json:"futureTasks"`

	GeneratedAt int64 `json:"generatedAt"`

	HostStates SingularityHostStateList `json:"hostStates"`

	LateTasks int32 `json:"lateTasks"`

	LbCleanupRequests int32 `json:"lbCleanupRequests"`

	LbCleanupTasks int32 `json:"lbCleanupTasks"`

	MaxTaskLag int64 `json:"maxTaskLag"`

	NumDeploys int32 `json:"numDeploys"`

	OldestDeploy int64 `json:"oldestDeploy"`

	OverProvisionedRequestIds swaggering.StringList `json:"overProvisionedRequestIds"`

	OverProvisionedRequests int32 `json:"overProvisionedRequests"`

	PausedRequests int32 `json:"pausedRequests"`

	PendingRequests int32 `json:"pendingRequests"`

	ScheduledTasks int32 `json:"scheduledTasks"`

	UnderProvisionedRequestIds swaggering.StringList `json:"underProvisionedRequestIds"`

	UnderProvisionedRequests int32 `json:"underProvisionedRequests"`

	UnknownRacks int32 `json:"unknownRacks"`

	UnknownSlaves int32 `json:"unknownSlaves"`
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

func (self *SingularityState) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *SingularityState) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityState) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityState) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *SingularityState) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityState", name)

	case "activeRacks", "ActiveRacks":
		v, ok := value.(int32)
		if ok {
			self.ActiveRacks = v
			self.present["activeRacks"] = true
			return nil
		} else {
			return fmt.Errorf("Field activeRacks/ActiveRacks: value %v(%T) couldn't be cast to type int32", value, value)
		}

	case "activeRequests", "ActiveRequests":
		v, ok := value.(int32)
		if ok {
			self.ActiveRequests = v
			self.present["activeRequests"] = true
			return nil
		} else {
			return fmt.Errorf("Field activeRequests/ActiveRequests: value %v(%T) couldn't be cast to type int32", value, value)
		}

	case "activeSlaves", "ActiveSlaves":
		v, ok := value.(int32)
		if ok {
			self.ActiveSlaves = v
			self.present["activeSlaves"] = true
			return nil
		} else {
			return fmt.Errorf("Field activeSlaves/ActiveSlaves: value %v(%T) couldn't be cast to type int32", value, value)
		}

	case "activeTasks", "ActiveTasks":
		v, ok := value.(int32)
		if ok {
			self.ActiveTasks = v
			self.present["activeTasks"] = true
			return nil
		} else {
			return fmt.Errorf("Field activeTasks/ActiveTasks: value %v(%T) couldn't be cast to type int32", value, value)
		}

	case "allRequests", "AllRequests":
		v, ok := value.(int32)
		if ok {
			self.AllRequests = v
			self.present["allRequests"] = true
			return nil
		} else {
			return fmt.Errorf("Field allRequests/AllRequests: value %v(%T) couldn't be cast to type int32", value, value)
		}

	case "authDatastoreHealthy", "AuthDatastoreHealthy":
		v, ok := value.(bool)
		if ok {
			self.AuthDatastoreHealthy = v
			self.present["authDatastoreHealthy"] = true
			return nil
		} else {
			return fmt.Errorf("Field authDatastoreHealthy/AuthDatastoreHealthy: value %v(%T) couldn't be cast to type bool", value, value)
		}

	case "cleaningRequests", "CleaningRequests":
		v, ok := value.(int32)
		if ok {
			self.CleaningRequests = v
			self.present["cleaningRequests"] = true
			return nil
		} else {
			return fmt.Errorf("Field cleaningRequests/CleaningRequests: value %v(%T) couldn't be cast to type int32", value, value)
		}

	case "cleaningTasks", "CleaningTasks":
		v, ok := value.(int32)
		if ok {
			self.CleaningTasks = v
			self.present["cleaningTasks"] = true
			return nil
		} else {
			return fmt.Errorf("Field cleaningTasks/CleaningTasks: value %v(%T) couldn't be cast to type int32", value, value)
		}

	case "cooldownRequests", "CooldownRequests":
		v, ok := value.(int32)
		if ok {
			self.CooldownRequests = v
			self.present["cooldownRequests"] = true
			return nil
		} else {
			return fmt.Errorf("Field cooldownRequests/CooldownRequests: value %v(%T) couldn't be cast to type int32", value, value)
		}

	case "deadRacks", "DeadRacks":
		v, ok := value.(int32)
		if ok {
			self.DeadRacks = v
			self.present["deadRacks"] = true
			return nil
		} else {
			return fmt.Errorf("Field deadRacks/DeadRacks: value %v(%T) couldn't be cast to type int32", value, value)
		}

	case "deadSlaves", "DeadSlaves":
		v, ok := value.(int32)
		if ok {
			self.DeadSlaves = v
			self.present["deadSlaves"] = true
			return nil
		} else {
			return fmt.Errorf("Field deadSlaves/DeadSlaves: value %v(%T) couldn't be cast to type int32", value, value)
		}

	case "decomissioningRacks", "DecomissioningRacks":
		v, ok := value.(int32)
		if ok {
			self.DecomissioningRacks = v
			self.present["decomissioningRacks"] = true
			return nil
		} else {
			return fmt.Errorf("Field decomissioningRacks/DecomissioningRacks: value %v(%T) couldn't be cast to type int32", value, value)
		}

	case "decomissioningSlaves", "DecomissioningSlaves":
		v, ok := value.(int32)
		if ok {
			self.DecomissioningSlaves = v
			self.present["decomissioningSlaves"] = true
			return nil
		} else {
			return fmt.Errorf("Field decomissioningSlaves/DecomissioningSlaves: value %v(%T) couldn't be cast to type int32", value, value)
		}

	case "decommissioningRacks", "DecommissioningRacks":
		v, ok := value.(int32)
		if ok {
			self.DecommissioningRacks = v
			self.present["decommissioningRacks"] = true
			return nil
		} else {
			return fmt.Errorf("Field decommissioningRacks/DecommissioningRacks: value %v(%T) couldn't be cast to type int32", value, value)
		}

	case "decommissioningSlaves", "DecommissioningSlaves":
		v, ok := value.(int32)
		if ok {
			self.DecommissioningSlaves = v
			self.present["decommissioningSlaves"] = true
			return nil
		} else {
			return fmt.Errorf("Field decommissioningSlaves/DecommissioningSlaves: value %v(%T) couldn't be cast to type int32", value, value)
		}

	case "finishedRequests", "FinishedRequests":
		v, ok := value.(int32)
		if ok {
			self.FinishedRequests = v
			self.present["finishedRequests"] = true
			return nil
		} else {
			return fmt.Errorf("Field finishedRequests/FinishedRequests: value %v(%T) couldn't be cast to type int32", value, value)
		}

	case "futureTasks", "FutureTasks":
		v, ok := value.(int32)
		if ok {
			self.FutureTasks = v
			self.present["futureTasks"] = true
			return nil
		} else {
			return fmt.Errorf("Field futureTasks/FutureTasks: value %v(%T) couldn't be cast to type int32", value, value)
		}

	case "generatedAt", "GeneratedAt":
		v, ok := value.(int64)
		if ok {
			self.GeneratedAt = v
			self.present["generatedAt"] = true
			return nil
		} else {
			return fmt.Errorf("Field generatedAt/GeneratedAt: value %v(%T) couldn't be cast to type int64", value, value)
		}

	case "hostStates", "HostStates":
		v, ok := value.(SingularityHostStateList)
		if ok {
			self.HostStates = v
			self.present["hostStates"] = true
			return nil
		} else {
			return fmt.Errorf("Field hostStates/HostStates: value %v(%T) couldn't be cast to type SingularityHostStateList", value, value)
		}

	case "lateTasks", "LateTasks":
		v, ok := value.(int32)
		if ok {
			self.LateTasks = v
			self.present["lateTasks"] = true
			return nil
		} else {
			return fmt.Errorf("Field lateTasks/LateTasks: value %v(%T) couldn't be cast to type int32", value, value)
		}

	case "lbCleanupRequests", "LbCleanupRequests":
		v, ok := value.(int32)
		if ok {
			self.LbCleanupRequests = v
			self.present["lbCleanupRequests"] = true
			return nil
		} else {
			return fmt.Errorf("Field lbCleanupRequests/LbCleanupRequests: value %v(%T) couldn't be cast to type int32", value, value)
		}

	case "lbCleanupTasks", "LbCleanupTasks":
		v, ok := value.(int32)
		if ok {
			self.LbCleanupTasks = v
			self.present["lbCleanupTasks"] = true
			return nil
		} else {
			return fmt.Errorf("Field lbCleanupTasks/LbCleanupTasks: value %v(%T) couldn't be cast to type int32", value, value)
		}

	case "maxTaskLag", "MaxTaskLag":
		v, ok := value.(int64)
		if ok {
			self.MaxTaskLag = v
			self.present["maxTaskLag"] = true
			return nil
		} else {
			return fmt.Errorf("Field maxTaskLag/MaxTaskLag: value %v(%T) couldn't be cast to type int64", value, value)
		}

	case "numDeploys", "NumDeploys":
		v, ok := value.(int32)
		if ok {
			self.NumDeploys = v
			self.present["numDeploys"] = true
			return nil
		} else {
			return fmt.Errorf("Field numDeploys/NumDeploys: value %v(%T) couldn't be cast to type int32", value, value)
		}

	case "oldestDeploy", "OldestDeploy":
		v, ok := value.(int64)
		if ok {
			self.OldestDeploy = v
			self.present["oldestDeploy"] = true
			return nil
		} else {
			return fmt.Errorf("Field oldestDeploy/OldestDeploy: value %v(%T) couldn't be cast to type int64", value, value)
		}

	case "overProvisionedRequestIds", "OverProvisionedRequestIds":
		v, ok := value.(swaggering.StringList)
		if ok {
			self.OverProvisionedRequestIds = v
			self.present["overProvisionedRequestIds"] = true
			return nil
		} else {
			return fmt.Errorf("Field overProvisionedRequestIds/OverProvisionedRequestIds: value %v(%T) couldn't be cast to type StringList", value, value)
		}

	case "overProvisionedRequests", "OverProvisionedRequests":
		v, ok := value.(int32)
		if ok {
			self.OverProvisionedRequests = v
			self.present["overProvisionedRequests"] = true
			return nil
		} else {
			return fmt.Errorf("Field overProvisionedRequests/OverProvisionedRequests: value %v(%T) couldn't be cast to type int32", value, value)
		}

	case "pausedRequests", "PausedRequests":
		v, ok := value.(int32)
		if ok {
			self.PausedRequests = v
			self.present["pausedRequests"] = true
			return nil
		} else {
			return fmt.Errorf("Field pausedRequests/PausedRequests: value %v(%T) couldn't be cast to type int32", value, value)
		}

	case "pendingRequests", "PendingRequests":
		v, ok := value.(int32)
		if ok {
			self.PendingRequests = v
			self.present["pendingRequests"] = true
			return nil
		} else {
			return fmt.Errorf("Field pendingRequests/PendingRequests: value %v(%T) couldn't be cast to type int32", value, value)
		}

	case "scheduledTasks", "ScheduledTasks":
		v, ok := value.(int32)
		if ok {
			self.ScheduledTasks = v
			self.present["scheduledTasks"] = true
			return nil
		} else {
			return fmt.Errorf("Field scheduledTasks/ScheduledTasks: value %v(%T) couldn't be cast to type int32", value, value)
		}

	case "underProvisionedRequestIds", "UnderProvisionedRequestIds":
		v, ok := value.(swaggering.StringList)
		if ok {
			self.UnderProvisionedRequestIds = v
			self.present["underProvisionedRequestIds"] = true
			return nil
		} else {
			return fmt.Errorf("Field underProvisionedRequestIds/UnderProvisionedRequestIds: value %v(%T) couldn't be cast to type StringList", value, value)
		}

	case "underProvisionedRequests", "UnderProvisionedRequests":
		v, ok := value.(int32)
		if ok {
			self.UnderProvisionedRequests = v
			self.present["underProvisionedRequests"] = true
			return nil
		} else {
			return fmt.Errorf("Field underProvisionedRequests/UnderProvisionedRequests: value %v(%T) couldn't be cast to type int32", value, value)
		}

	case "unknownRacks", "UnknownRacks":
		v, ok := value.(int32)
		if ok {
			self.UnknownRacks = v
			self.present["unknownRacks"] = true
			return nil
		} else {
			return fmt.Errorf("Field unknownRacks/UnknownRacks: value %v(%T) couldn't be cast to type int32", value, value)
		}

	case "unknownSlaves", "UnknownSlaves":
		v, ok := value.(int32)
		if ok {
			self.UnknownSlaves = v
			self.present["unknownSlaves"] = true
			return nil
		} else {
			return fmt.Errorf("Field unknownSlaves/UnknownSlaves: value %v(%T) couldn't be cast to type int32", value, value)
		}

	}
}

func (self *SingularityState) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityState", name)

	case "activeRacks", "ActiveRacks":
		if self.present != nil {
			if _, ok := self.present["activeRacks"]; ok {
				return self.ActiveRacks, nil
			}
		}
		return nil, fmt.Errorf("Field ActiveRacks no set on ActiveRacks %+v", self)

	case "activeRequests", "ActiveRequests":
		if self.present != nil {
			if _, ok := self.present["activeRequests"]; ok {
				return self.ActiveRequests, nil
			}
		}
		return nil, fmt.Errorf("Field ActiveRequests no set on ActiveRequests %+v", self)

	case "activeSlaves", "ActiveSlaves":
		if self.present != nil {
			if _, ok := self.present["activeSlaves"]; ok {
				return self.ActiveSlaves, nil
			}
		}
		return nil, fmt.Errorf("Field ActiveSlaves no set on ActiveSlaves %+v", self)

	case "activeTasks", "ActiveTasks":
		if self.present != nil {
			if _, ok := self.present["activeTasks"]; ok {
				return self.ActiveTasks, nil
			}
		}
		return nil, fmt.Errorf("Field ActiveTasks no set on ActiveTasks %+v", self)

	case "allRequests", "AllRequests":
		if self.present != nil {
			if _, ok := self.present["allRequests"]; ok {
				return self.AllRequests, nil
			}
		}
		return nil, fmt.Errorf("Field AllRequests no set on AllRequests %+v", self)

	case "authDatastoreHealthy", "AuthDatastoreHealthy":
		if self.present != nil {
			if _, ok := self.present["authDatastoreHealthy"]; ok {
				return self.AuthDatastoreHealthy, nil
			}
		}
		return nil, fmt.Errorf("Field AuthDatastoreHealthy no set on AuthDatastoreHealthy %+v", self)

	case "cleaningRequests", "CleaningRequests":
		if self.present != nil {
			if _, ok := self.present["cleaningRequests"]; ok {
				return self.CleaningRequests, nil
			}
		}
		return nil, fmt.Errorf("Field CleaningRequests no set on CleaningRequests %+v", self)

	case "cleaningTasks", "CleaningTasks":
		if self.present != nil {
			if _, ok := self.present["cleaningTasks"]; ok {
				return self.CleaningTasks, nil
			}
		}
		return nil, fmt.Errorf("Field CleaningTasks no set on CleaningTasks %+v", self)

	case "cooldownRequests", "CooldownRequests":
		if self.present != nil {
			if _, ok := self.present["cooldownRequests"]; ok {
				return self.CooldownRequests, nil
			}
		}
		return nil, fmt.Errorf("Field CooldownRequests no set on CooldownRequests %+v", self)

	case "deadRacks", "DeadRacks":
		if self.present != nil {
			if _, ok := self.present["deadRacks"]; ok {
				return self.DeadRacks, nil
			}
		}
		return nil, fmt.Errorf("Field DeadRacks no set on DeadRacks %+v", self)

	case "deadSlaves", "DeadSlaves":
		if self.present != nil {
			if _, ok := self.present["deadSlaves"]; ok {
				return self.DeadSlaves, nil
			}
		}
		return nil, fmt.Errorf("Field DeadSlaves no set on DeadSlaves %+v", self)

	case "decomissioningRacks", "DecomissioningRacks":
		if self.present != nil {
			if _, ok := self.present["decomissioningRacks"]; ok {
				return self.DecomissioningRacks, nil
			}
		}
		return nil, fmt.Errorf("Field DecomissioningRacks no set on DecomissioningRacks %+v", self)

	case "decomissioningSlaves", "DecomissioningSlaves":
		if self.present != nil {
			if _, ok := self.present["decomissioningSlaves"]; ok {
				return self.DecomissioningSlaves, nil
			}
		}
		return nil, fmt.Errorf("Field DecomissioningSlaves no set on DecomissioningSlaves %+v", self)

	case "decommissioningRacks", "DecommissioningRacks":
		if self.present != nil {
			if _, ok := self.present["decommissioningRacks"]; ok {
				return self.DecommissioningRacks, nil
			}
		}
		return nil, fmt.Errorf("Field DecommissioningRacks no set on DecommissioningRacks %+v", self)

	case "decommissioningSlaves", "DecommissioningSlaves":
		if self.present != nil {
			if _, ok := self.present["decommissioningSlaves"]; ok {
				return self.DecommissioningSlaves, nil
			}
		}
		return nil, fmt.Errorf("Field DecommissioningSlaves no set on DecommissioningSlaves %+v", self)

	case "finishedRequests", "FinishedRequests":
		if self.present != nil {
			if _, ok := self.present["finishedRequests"]; ok {
				return self.FinishedRequests, nil
			}
		}
		return nil, fmt.Errorf("Field FinishedRequests no set on FinishedRequests %+v", self)

	case "futureTasks", "FutureTasks":
		if self.present != nil {
			if _, ok := self.present["futureTasks"]; ok {
				return self.FutureTasks, nil
			}
		}
		return nil, fmt.Errorf("Field FutureTasks no set on FutureTasks %+v", self)

	case "generatedAt", "GeneratedAt":
		if self.present != nil {
			if _, ok := self.present["generatedAt"]; ok {
				return self.GeneratedAt, nil
			}
		}
		return nil, fmt.Errorf("Field GeneratedAt no set on GeneratedAt %+v", self)

	case "hostStates", "HostStates":
		if self.present != nil {
			if _, ok := self.present["hostStates"]; ok {
				return self.HostStates, nil
			}
		}
		return nil, fmt.Errorf("Field HostStates no set on HostStates %+v", self)

	case "lateTasks", "LateTasks":
		if self.present != nil {
			if _, ok := self.present["lateTasks"]; ok {
				return self.LateTasks, nil
			}
		}
		return nil, fmt.Errorf("Field LateTasks no set on LateTasks %+v", self)

	case "lbCleanupRequests", "LbCleanupRequests":
		if self.present != nil {
			if _, ok := self.present["lbCleanupRequests"]; ok {
				return self.LbCleanupRequests, nil
			}
		}
		return nil, fmt.Errorf("Field LbCleanupRequests no set on LbCleanupRequests %+v", self)

	case "lbCleanupTasks", "LbCleanupTasks":
		if self.present != nil {
			if _, ok := self.present["lbCleanupTasks"]; ok {
				return self.LbCleanupTasks, nil
			}
		}
		return nil, fmt.Errorf("Field LbCleanupTasks no set on LbCleanupTasks %+v", self)

	case "maxTaskLag", "MaxTaskLag":
		if self.present != nil {
			if _, ok := self.present["maxTaskLag"]; ok {
				return self.MaxTaskLag, nil
			}
		}
		return nil, fmt.Errorf("Field MaxTaskLag no set on MaxTaskLag %+v", self)

	case "numDeploys", "NumDeploys":
		if self.present != nil {
			if _, ok := self.present["numDeploys"]; ok {
				return self.NumDeploys, nil
			}
		}
		return nil, fmt.Errorf("Field NumDeploys no set on NumDeploys %+v", self)

	case "oldestDeploy", "OldestDeploy":
		if self.present != nil {
			if _, ok := self.present["oldestDeploy"]; ok {
				return self.OldestDeploy, nil
			}
		}
		return nil, fmt.Errorf("Field OldestDeploy no set on OldestDeploy %+v", self)

	case "overProvisionedRequestIds", "OverProvisionedRequestIds":
		if self.present != nil {
			if _, ok := self.present["overProvisionedRequestIds"]; ok {
				return self.OverProvisionedRequestIds, nil
			}
		}
		return nil, fmt.Errorf("Field OverProvisionedRequestIds no set on OverProvisionedRequestIds %+v", self)

	case "overProvisionedRequests", "OverProvisionedRequests":
		if self.present != nil {
			if _, ok := self.present["overProvisionedRequests"]; ok {
				return self.OverProvisionedRequests, nil
			}
		}
		return nil, fmt.Errorf("Field OverProvisionedRequests no set on OverProvisionedRequests %+v", self)

	case "pausedRequests", "PausedRequests":
		if self.present != nil {
			if _, ok := self.present["pausedRequests"]; ok {
				return self.PausedRequests, nil
			}
		}
		return nil, fmt.Errorf("Field PausedRequests no set on PausedRequests %+v", self)

	case "pendingRequests", "PendingRequests":
		if self.present != nil {
			if _, ok := self.present["pendingRequests"]; ok {
				return self.PendingRequests, nil
			}
		}
		return nil, fmt.Errorf("Field PendingRequests no set on PendingRequests %+v", self)

	case "scheduledTasks", "ScheduledTasks":
		if self.present != nil {
			if _, ok := self.present["scheduledTasks"]; ok {
				return self.ScheduledTasks, nil
			}
		}
		return nil, fmt.Errorf("Field ScheduledTasks no set on ScheduledTasks %+v", self)

	case "underProvisionedRequestIds", "UnderProvisionedRequestIds":
		if self.present != nil {
			if _, ok := self.present["underProvisionedRequestIds"]; ok {
				return self.UnderProvisionedRequestIds, nil
			}
		}
		return nil, fmt.Errorf("Field UnderProvisionedRequestIds no set on UnderProvisionedRequestIds %+v", self)

	case "underProvisionedRequests", "UnderProvisionedRequests":
		if self.present != nil {
			if _, ok := self.present["underProvisionedRequests"]; ok {
				return self.UnderProvisionedRequests, nil
			}
		}
		return nil, fmt.Errorf("Field UnderProvisionedRequests no set on UnderProvisionedRequests %+v", self)

	case "unknownRacks", "UnknownRacks":
		if self.present != nil {
			if _, ok := self.present["unknownRacks"]; ok {
				return self.UnknownRacks, nil
			}
		}
		return nil, fmt.Errorf("Field UnknownRacks no set on UnknownRacks %+v", self)

	case "unknownSlaves", "UnknownSlaves":
		if self.present != nil {
			if _, ok := self.present["unknownSlaves"]; ok {
				return self.UnknownSlaves, nil
			}
		}
		return nil, fmt.Errorf("Field UnknownSlaves no set on UnknownSlaves %+v", self)

	}
}

func (self *SingularityState) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityState", name)

	case "activeRacks", "ActiveRacks":
		self.present["activeRacks"] = false

	case "activeRequests", "ActiveRequests":
		self.present["activeRequests"] = false

	case "activeSlaves", "ActiveSlaves":
		self.present["activeSlaves"] = false

	case "activeTasks", "ActiveTasks":
		self.present["activeTasks"] = false

	case "allRequests", "AllRequests":
		self.present["allRequests"] = false

	case "authDatastoreHealthy", "AuthDatastoreHealthy":
		self.present["authDatastoreHealthy"] = false

	case "cleaningRequests", "CleaningRequests":
		self.present["cleaningRequests"] = false

	case "cleaningTasks", "CleaningTasks":
		self.present["cleaningTasks"] = false

	case "cooldownRequests", "CooldownRequests":
		self.present["cooldownRequests"] = false

	case "deadRacks", "DeadRacks":
		self.present["deadRacks"] = false

	case "deadSlaves", "DeadSlaves":
		self.present["deadSlaves"] = false

	case "decomissioningRacks", "DecomissioningRacks":
		self.present["decomissioningRacks"] = false

	case "decomissioningSlaves", "DecomissioningSlaves":
		self.present["decomissioningSlaves"] = false

	case "decommissioningRacks", "DecommissioningRacks":
		self.present["decommissioningRacks"] = false

	case "decommissioningSlaves", "DecommissioningSlaves":
		self.present["decommissioningSlaves"] = false

	case "finishedRequests", "FinishedRequests":
		self.present["finishedRequests"] = false

	case "futureTasks", "FutureTasks":
		self.present["futureTasks"] = false

	case "generatedAt", "GeneratedAt":
		self.present["generatedAt"] = false

	case "hostStates", "HostStates":
		self.present["hostStates"] = false

	case "lateTasks", "LateTasks":
		self.present["lateTasks"] = false

	case "lbCleanupRequests", "LbCleanupRequests":
		self.present["lbCleanupRequests"] = false

	case "lbCleanupTasks", "LbCleanupTasks":
		self.present["lbCleanupTasks"] = false

	case "maxTaskLag", "MaxTaskLag":
		self.present["maxTaskLag"] = false

	case "numDeploys", "NumDeploys":
		self.present["numDeploys"] = false

	case "oldestDeploy", "OldestDeploy":
		self.present["oldestDeploy"] = false

	case "overProvisionedRequestIds", "OverProvisionedRequestIds":
		self.present["overProvisionedRequestIds"] = false

	case "overProvisionedRequests", "OverProvisionedRequests":
		self.present["overProvisionedRequests"] = false

	case "pausedRequests", "PausedRequests":
		self.present["pausedRequests"] = false

	case "pendingRequests", "PendingRequests":
		self.present["pendingRequests"] = false

	case "scheduledTasks", "ScheduledTasks":
		self.present["scheduledTasks"] = false

	case "underProvisionedRequestIds", "UnderProvisionedRequestIds":
		self.present["underProvisionedRequestIds"] = false

	case "underProvisionedRequests", "UnderProvisionedRequests":
		self.present["underProvisionedRequests"] = false

	case "unknownRacks", "UnknownRacks":
		self.present["unknownRacks"] = false

	case "unknownSlaves", "UnknownSlaves":
		self.present["unknownSlaves"] = false

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
