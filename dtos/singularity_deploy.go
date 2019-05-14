package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityDeployHealthcheckProtocol string

const (
	SingularityDeployHealthcheckProtocolhttp  SingularityDeployHealthcheckProtocol = "http"
	SingularityDeployHealthcheckProtocolhttps SingularityDeployHealthcheckProtocol = "https"
)

type SingularityDeploy struct {
	LoadBalancerGroups                    *swaggering.StringList                   `json:"loadBalancerGroups,omitempty"`
	LoadBalancerDomains                   *swaggering.StringList                   `json:"loadBalancerDomains,omitempty"`
	LoadBalancerUpstreamGroup             *string                                  `json:"loadBalancerUpstreamGroup,omitempty"`
	Resources                             *Resources                               `json:"resources,omitempty"`
	Command                               *string                                  `json:"command,omitempty"`
	Env                                   *map[string]string                       `json:"env,omitempty"`
	MesosTaskLabels                       *map[int64]SingularityMesosTaskLabelList `json:"mesosTaskLabels,omitempty"`
	HealthcheckMaxRetries                 *int32                                   `json:"healthcheckMaxRetries,omitempty"`
	MaxTaskRetries                        *int32                                   `json:"maxTaskRetries,omitempty"`
	LoadBalancerOptions                   *map[string]interface{}                  `json:"loadBalancerOptions,omitempty"`
	LoadBalancerServiceIdOverride         *string                                  `json:"loadBalancerServiceIdOverride,omitempty"`
	RequestId                             *string                                  `json:"requestId,omitempty"`
	Version                               *string                                  `json:"version,omitempty"`
	Uris                                  *SingularityMesosArtifactList            `json:"uris,omitempty"`
	HealthcheckIntervalSeconds            *int64                                   `json:"healthcheckIntervalSeconds,omitempty"`
	ServiceBasePath                       *string                                  `json:"serviceBasePath,omitempty"`
	Arguments                             *swaggering.StringList                   `json:"arguments,omitempty"`
	ExecutorData                          *ExecutorData                            `json:"executorData,omitempty"`
	Labels                                *map[string]string                       `json:"labels,omitempty"`
	AutoAdvanceDeploySteps                *bool                                    `json:"autoAdvanceDeploySteps,omitempty"`
	Metadata                              *map[string]string                       `json:"metadata,omitempty"`
	CustomExecutorResources               *Resources                               `json:"customExecutorResources,omitempty"`
	TaskEnv                               *map[int64]map[string]string             `json:"taskEnv,omitempty"`
	HealthcheckMaxTotalTimeoutSeconds     *int64                                   `json:"healthcheckMaxTotalTimeoutSeconds,omitempty"`
	LoadBalancerTemplate                  *string                                  `json:"loadBalancerTemplate,omitempty"`
	DeployInstanceCountPerStep            *int32                                   `json:"deployInstanceCountPerStep,omitempty"`
	User                                  *string                                  `json:"user,omitempty"`
	CustomExecutorId                      *string                                  `json:"customExecutorId,omitempty"`
	CustomExecutorSource                  *string                                  `json:"customExecutorSource,omitempty"`
	HealthcheckProtocol                   *SingularityDeployHealthcheckProtocol    `json:"healthcheckProtocol,omitempty"`
	SkipHealthchecksOnDeploy              *bool                                    `json:"skipHealthchecksOnDeploy,omitempty"`
	ConsiderHealthyAfterRunningForSeconds *int64                                   `json:"considerHealthyAfterRunningForSeconds,omitempty"`
	Timestamp                             *int64                                   `json:"timestamp,omitempty"`
	CustomExecutorCmd                     *string                                  `json:"customExecutorCmd,omitempty"`
	HealthcheckTimeoutSeconds             *int64                                   `json:"healthcheckTimeoutSeconds,omitempty"`
	DeployHealthTimeoutSeconds            *int64                                   `json:"deployHealthTimeoutSeconds,omitempty"`
	Shell                                 *bool                                    `json:"shell,omitempty"`
	ContainerInfo                         *SingularityContainerInfo                `json:"containerInfo,omitempty"`
	TaskLabels                            *map[int64]map[string]string             `json:"taskLabels,omitempty"`
	HealthcheckUri                        *string                                  `json:"healthcheckUri,omitempty"`
	Healthcheck                           *HealthcheckOptions                      `json:"healthcheck,omitempty"`
	DeployStepWaitTimeMs                  *int32                                   `json:"deployStepWaitTimeMs,omitempty"`
	Id                                    *string                                  `json:"id,omitempty"`
	MesosLabels                           *SingularityMesosTaskLabelList           `json:"mesosLabels,omitempty"`
	HealthcheckPortIndex                  *int32                                   `json:"healthcheckPortIndex,omitempty"`
	LoadBalancerPortIndex                 *int32                                   `json:"loadBalancerPortIndex,omitempty"`
	LoadBalancerAdditionalRoutes          *swaggering.StringList                   `json:"loadBalancerAdditionalRoutes,omitempty"`
}

func (self *SingularityDeploy) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *SingularityDeploy) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityDeploy); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityDeploy cannot copy the values from %#v", other)
}

func (self *SingularityDeploy) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityDeploy) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityDeploy) SetField(name string, value interface{}) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityDeploy", name)

	case "loadBalancerGroups", "LoadBalancerGroups":
		v, ok := value.(swaggering.StringList)
		if ok {
			self.LoadBalancerGroups = &v
			return nil
		}
		return fmt.Errorf("Field loadBalancerGroups/LoadBalancerGroups: value %v(%T) couldn't be cast to type swaggering.StringList", value, value)

	case "loadBalancerDomains", "LoadBalancerDomains":
		v, ok := value.(swaggering.StringList)
		if ok {
			self.LoadBalancerDomains = &v
			return nil
		}
		return fmt.Errorf("Field loadBalancerDomains/LoadBalancerDomains: value %v(%T) couldn't be cast to type swaggering.StringList", value, value)

	case "loadBalancerUpstreamGroup", "LoadBalancerUpstreamGroup":
		v, ok := value.(string)
		if ok {
			self.LoadBalancerUpstreamGroup = &v
			return nil
		}
		return fmt.Errorf("Field loadBalancerUpstreamGroup/LoadBalancerUpstreamGroup: value %v(%T) couldn't be cast to type string", value, value)

	case "resources", "Resources":
		v, ok := value.(*Resources)
		if ok {
			self.Resources = v
			return nil
		}
		return fmt.Errorf("Field resources/Resources: value %v(%T) couldn't be cast to type *Resources", value, value)

	case "command", "Command":
		v, ok := value.(string)
		if ok {
			self.Command = &v
			return nil
		}
		return fmt.Errorf("Field command/Command: value %v(%T) couldn't be cast to type string", value, value)

	case "env", "Env":
		v, ok := value.(map[string]string)
		if ok {
			self.Env = &v
			return nil
		}
		return fmt.Errorf("Field env/Env: value %v(%T) couldn't be cast to type map[string]string", value, value)

	case "mesosTaskLabels", "MesosTaskLabels":
		v, ok := value.(map[int64]SingularityMesosTaskLabelList)
		if ok {
			self.MesosTaskLabels = &v
			return nil
		}
		return fmt.Errorf("Field mesosTaskLabels/MesosTaskLabels: value %v(%T) couldn't be cast to type map[int64]SingularityMesosTaskLabelList", value, value)

	case "healthcheckMaxRetries", "HealthcheckMaxRetries":
		v, ok := value.(int32)
		if ok {
			self.HealthcheckMaxRetries = &v
			return nil
		}
		return fmt.Errorf("Field healthcheckMaxRetries/HealthcheckMaxRetries: value %v(%T) couldn't be cast to type int32", value, value)

	case "maxTaskRetries", "MaxTaskRetries":
		v, ok := value.(int32)
		if ok {
			self.MaxTaskRetries = &v
			return nil
		}
		return fmt.Errorf("Field maxTaskRetries/MaxTaskRetries: value %v(%T) couldn't be cast to type int32", value, value)

	case "loadBalancerOptions", "LoadBalancerOptions":
		v, ok := value.(map[string]interface{})
		if ok {
			self.LoadBalancerOptions = &v
			return nil
		}
		return fmt.Errorf("Field loadBalancerOptions/LoadBalancerOptions: value %v(%T) couldn't be cast to type map[string]interface{}", value, value)

	case "loadBalancerServiceIdOverride", "LoadBalancerServiceIdOverride":
		v, ok := value.(string)
		if ok {
			self.LoadBalancerServiceIdOverride = &v
			return nil
		}
		return fmt.Errorf("Field loadBalancerServiceIdOverride/LoadBalancerServiceIdOverride: value %v(%T) couldn't be cast to type string", value, value)

	case "requestId", "RequestId":
		v, ok := value.(string)
		if ok {
			self.RequestId = &v
			return nil
		}
		return fmt.Errorf("Field requestId/RequestId: value %v(%T) couldn't be cast to type string", value, value)

	case "version", "Version":
		v, ok := value.(string)
		if ok {
			self.Version = &v
			return nil
		}
		return fmt.Errorf("Field version/Version: value %v(%T) couldn't be cast to type string", value, value)

	case "uris", "Uris":
		v, ok := value.(SingularityMesosArtifactList)
		if ok {
			self.Uris = &v
			return nil
		}
		return fmt.Errorf("Field uris/Uris: value %v(%T) couldn't be cast to type SingularityMesosArtifactList", value, value)

	case "healthcheckIntervalSeconds", "HealthcheckIntervalSeconds":
		v, ok := value.(int64)
		if ok {
			self.HealthcheckIntervalSeconds = &v
			return nil
		}
		return fmt.Errorf("Field healthcheckIntervalSeconds/HealthcheckIntervalSeconds: value %v(%T) couldn't be cast to type int64", value, value)

	case "serviceBasePath", "ServiceBasePath":
		v, ok := value.(string)
		if ok {
			self.ServiceBasePath = &v
			return nil
		}
		return fmt.Errorf("Field serviceBasePath/ServiceBasePath: value %v(%T) couldn't be cast to type string", value, value)

	case "arguments", "Arguments":
		v, ok := value.(swaggering.StringList)
		if ok {
			self.Arguments = &v
			return nil
		}
		return fmt.Errorf("Field arguments/Arguments: value %v(%T) couldn't be cast to type swaggering.StringList", value, value)

	case "executorData", "ExecutorData":
		v, ok := value.(*ExecutorData)
		if ok {
			self.ExecutorData = v
			return nil
		}
		return fmt.Errorf("Field executorData/ExecutorData: value %v(%T) couldn't be cast to type *ExecutorData", value, value)

	case "labels", "Labels":
		v, ok := value.(map[string]string)
		if ok {
			self.Labels = &v
			return nil
		}
		return fmt.Errorf("Field labels/Labels: value %v(%T) couldn't be cast to type map[string]string", value, value)

	case "autoAdvanceDeploySteps", "AutoAdvanceDeploySteps":
		v, ok := value.(bool)
		if ok {
			self.AutoAdvanceDeploySteps = &v
			return nil
		}
		return fmt.Errorf("Field autoAdvanceDeploySteps/AutoAdvanceDeploySteps: value %v(%T) couldn't be cast to type bool", value, value)

	case "metadata", "Metadata":
		v, ok := value.(map[string]string)
		if ok {
			self.Metadata = &v
			return nil
		}
		return fmt.Errorf("Field metadata/Metadata: value %v(%T) couldn't be cast to type map[string]string", value, value)

	case "customExecutorResources", "CustomExecutorResources":
		v, ok := value.(*Resources)
		if ok {
			self.CustomExecutorResources = v
			return nil
		}
		return fmt.Errorf("Field customExecutorResources/CustomExecutorResources: value %v(%T) couldn't be cast to type *Resources", value, value)

	case "taskEnv", "TaskEnv":
		v, ok := value.(map[int64]map[string]string)
		if ok {
			self.TaskEnv = &v
			return nil
		}
		return fmt.Errorf("Field taskEnv/TaskEnv: value %v(%T) couldn't be cast to type map[int64]map[string]string", value, value)

	case "healthcheckMaxTotalTimeoutSeconds", "HealthcheckMaxTotalTimeoutSeconds":
		v, ok := value.(int64)
		if ok {
			self.HealthcheckMaxTotalTimeoutSeconds = &v
			return nil
		}
		return fmt.Errorf("Field healthcheckMaxTotalTimeoutSeconds/HealthcheckMaxTotalTimeoutSeconds: value %v(%T) couldn't be cast to type int64", value, value)

	case "loadBalancerTemplate", "LoadBalancerTemplate":
		v, ok := value.(string)
		if ok {
			self.LoadBalancerTemplate = &v
			return nil
		}
		return fmt.Errorf("Field loadBalancerTemplate/LoadBalancerTemplate: value %v(%T) couldn't be cast to type string", value, value)

	case "deployInstanceCountPerStep", "DeployInstanceCountPerStep":
		v, ok := value.(int32)
		if ok {
			self.DeployInstanceCountPerStep = &v
			return nil
		}
		return fmt.Errorf("Field deployInstanceCountPerStep/DeployInstanceCountPerStep: value %v(%T) couldn't be cast to type int32", value, value)

	case "user", "User":
		v, ok := value.(string)
		if ok {
			self.User = &v
			return nil
		}
		return fmt.Errorf("Field user/User: value %v(%T) couldn't be cast to type string", value, value)

	case "customExecutorId", "CustomExecutorId":
		v, ok := value.(string)
		if ok {
			self.CustomExecutorId = &v
			return nil
		}
		return fmt.Errorf("Field customExecutorId/CustomExecutorId: value %v(%T) couldn't be cast to type string", value, value)

	case "customExecutorSource", "CustomExecutorSource":
		v, ok := value.(string)
		if ok {
			self.CustomExecutorSource = &v
			return nil
		}
		return fmt.Errorf("Field customExecutorSource/CustomExecutorSource: value %v(%T) couldn't be cast to type string", value, value)

	case "healthcheckProtocol", "HealthcheckProtocol":
		v, ok := value.(SingularityDeployHealthcheckProtocol)
		if ok {
			self.HealthcheckProtocol = &v
			return nil
		}
		return fmt.Errorf("Field healthcheckProtocol/HealthcheckProtocol: value %v(%T) couldn't be cast to type SingularityDeployHealthcheckProtocol", value, value)

	case "skipHealthchecksOnDeploy", "SkipHealthchecksOnDeploy":
		v, ok := value.(bool)
		if ok {
			self.SkipHealthchecksOnDeploy = &v
			return nil
		}
		return fmt.Errorf("Field skipHealthchecksOnDeploy/SkipHealthchecksOnDeploy: value %v(%T) couldn't be cast to type bool", value, value)

	case "considerHealthyAfterRunningForSeconds", "ConsiderHealthyAfterRunningForSeconds":
		v, ok := value.(int64)
		if ok {
			self.ConsiderHealthyAfterRunningForSeconds = &v
			return nil
		}
		return fmt.Errorf("Field considerHealthyAfterRunningForSeconds/ConsiderHealthyAfterRunningForSeconds: value %v(%T) couldn't be cast to type int64", value, value)

	case "timestamp", "Timestamp":
		v, ok := value.(int64)
		if ok {
			self.Timestamp = &v
			return nil
		}
		return fmt.Errorf("Field timestamp/Timestamp: value %v(%T) couldn't be cast to type int64", value, value)

	case "customExecutorCmd", "CustomExecutorCmd":
		v, ok := value.(string)
		if ok {
			self.CustomExecutorCmd = &v
			return nil
		}
		return fmt.Errorf("Field customExecutorCmd/CustomExecutorCmd: value %v(%T) couldn't be cast to type string", value, value)

	case "healthcheckTimeoutSeconds", "HealthcheckTimeoutSeconds":
		v, ok := value.(int64)
		if ok {
			self.HealthcheckTimeoutSeconds = &v
			return nil
		}
		return fmt.Errorf("Field healthcheckTimeoutSeconds/HealthcheckTimeoutSeconds: value %v(%T) couldn't be cast to type int64", value, value)

	case "deployHealthTimeoutSeconds", "DeployHealthTimeoutSeconds":
		v, ok := value.(int64)
		if ok {
			self.DeployHealthTimeoutSeconds = &v
			return nil
		}
		return fmt.Errorf("Field deployHealthTimeoutSeconds/DeployHealthTimeoutSeconds: value %v(%T) couldn't be cast to type int64", value, value)

	case "shell", "Shell":
		v, ok := value.(bool)
		if ok {
			self.Shell = &v
			return nil
		}
		return fmt.Errorf("Field shell/Shell: value %v(%T) couldn't be cast to type bool", value, value)

	case "containerInfo", "ContainerInfo":
		v, ok := value.(*SingularityContainerInfo)
		if ok {
			self.ContainerInfo = v
			return nil
		}
		return fmt.Errorf("Field containerInfo/ContainerInfo: value %v(%T) couldn't be cast to type *SingularityContainerInfo", value, value)

	case "taskLabels", "TaskLabels":
		v, ok := value.(map[int64]map[string]string)
		if ok {
			self.TaskLabels = &v
			return nil
		}
		return fmt.Errorf("Field taskLabels/TaskLabels: value %v(%T) couldn't be cast to type map[int64]map[string]string", value, value)

	case "healthcheckUri", "HealthcheckUri":
		v, ok := value.(string)
		if ok {
			self.HealthcheckUri = &v
			return nil
		}
		return fmt.Errorf("Field healthcheckUri/HealthcheckUri: value %v(%T) couldn't be cast to type string", value, value)

	case "healthcheck", "Healthcheck":
		v, ok := value.(*HealthcheckOptions)
		if ok {
			self.Healthcheck = v
			return nil
		}
		return fmt.Errorf("Field healthcheck/Healthcheck: value %v(%T) couldn't be cast to type *HealthcheckOptions", value, value)

	case "deployStepWaitTimeMs", "DeployStepWaitTimeMs":
		v, ok := value.(int32)
		if ok {
			self.DeployStepWaitTimeMs = &v
			return nil
		}
		return fmt.Errorf("Field deployStepWaitTimeMs/DeployStepWaitTimeMs: value %v(%T) couldn't be cast to type int32", value, value)

	case "id", "Id":
		v, ok := value.(string)
		if ok {
			self.Id = &v
			return nil
		}
		return fmt.Errorf("Field id/Id: value %v(%T) couldn't be cast to type string", value, value)

	case "mesosLabels", "MesosLabels":
		v, ok := value.(SingularityMesosTaskLabelList)
		if ok {
			self.MesosLabels = &v
			return nil
		}
		return fmt.Errorf("Field mesosLabels/MesosLabels: value %v(%T) couldn't be cast to type SingularityMesosTaskLabelList", value, value)

	case "healthcheckPortIndex", "HealthcheckPortIndex":
		v, ok := value.(int32)
		if ok {
			self.HealthcheckPortIndex = &v
			return nil
		}
		return fmt.Errorf("Field healthcheckPortIndex/HealthcheckPortIndex: value %v(%T) couldn't be cast to type int32", value, value)

	case "loadBalancerPortIndex", "LoadBalancerPortIndex":
		v, ok := value.(int32)
		if ok {
			self.LoadBalancerPortIndex = &v
			return nil
		}
		return fmt.Errorf("Field loadBalancerPortIndex/LoadBalancerPortIndex: value %v(%T) couldn't be cast to type int32", value, value)

	case "loadBalancerAdditionalRoutes", "LoadBalancerAdditionalRoutes":
		v, ok := value.(swaggering.StringList)
		if ok {
			self.LoadBalancerAdditionalRoutes = &v
			return nil
		}
		return fmt.Errorf("Field loadBalancerAdditionalRoutes/LoadBalancerAdditionalRoutes: value %v(%T) couldn't be cast to type swaggering.StringList", value, value)

	}
}

func (self *SingularityDeploy) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityDeploy", name)

	case "loadBalancerGroups", "LoadBalancerGroups":
		return *self.LoadBalancerGroups, nil
		return nil, fmt.Errorf("Field LoadBalancerGroups no set on LoadBalancerGroups %+v", self)

	case "loadBalancerDomains", "LoadBalancerDomains":
		return *self.LoadBalancerDomains, nil
		return nil, fmt.Errorf("Field LoadBalancerDomains no set on LoadBalancerDomains %+v", self)

	case "loadBalancerUpstreamGroup", "LoadBalancerUpstreamGroup":
		return *self.LoadBalancerUpstreamGroup, nil
		return nil, fmt.Errorf("Field LoadBalancerUpstreamGroup no set on LoadBalancerUpstreamGroup %+v", self)

	case "resources", "Resources":
		return self.Resources, nil
		return nil, fmt.Errorf("Field Resources no set on Resources %+v", self)

	case "command", "Command":
		return *self.Command, nil
		return nil, fmt.Errorf("Field Command no set on Command %+v", self)

	case "env", "Env":
		return *self.Env, nil
		return nil, fmt.Errorf("Field Env no set on Env %+v", self)

	case "mesosTaskLabels", "MesosTaskLabels":
		return *self.MesosTaskLabels, nil
		return nil, fmt.Errorf("Field MesosTaskLabels no set on MesosTaskLabels %+v", self)

	case "healthcheckMaxRetries", "HealthcheckMaxRetries":
		return *self.HealthcheckMaxRetries, nil
		return nil, fmt.Errorf("Field HealthcheckMaxRetries no set on HealthcheckMaxRetries %+v", self)

	case "maxTaskRetries", "MaxTaskRetries":
		return *self.MaxTaskRetries, nil
		return nil, fmt.Errorf("Field MaxTaskRetries no set on MaxTaskRetries %+v", self)

	case "loadBalancerOptions", "LoadBalancerOptions":
		return *self.LoadBalancerOptions, nil
		return nil, fmt.Errorf("Field LoadBalancerOptions no set on LoadBalancerOptions %+v", self)

	case "loadBalancerServiceIdOverride", "LoadBalancerServiceIdOverride":
		return *self.LoadBalancerServiceIdOverride, nil
		return nil, fmt.Errorf("Field LoadBalancerServiceIdOverride no set on LoadBalancerServiceIdOverride %+v", self)

	case "requestId", "RequestId":
		return *self.RequestId, nil
		return nil, fmt.Errorf("Field RequestId no set on RequestId %+v", self)

	case "version", "Version":
		return *self.Version, nil
		return nil, fmt.Errorf("Field Version no set on Version %+v", self)

	case "uris", "Uris":
		return *self.Uris, nil
		return nil, fmt.Errorf("Field Uris no set on Uris %+v", self)

	case "healthcheckIntervalSeconds", "HealthcheckIntervalSeconds":
		return *self.HealthcheckIntervalSeconds, nil
		return nil, fmt.Errorf("Field HealthcheckIntervalSeconds no set on HealthcheckIntervalSeconds %+v", self)

	case "serviceBasePath", "ServiceBasePath":
		return *self.ServiceBasePath, nil
		return nil, fmt.Errorf("Field ServiceBasePath no set on ServiceBasePath %+v", self)

	case "arguments", "Arguments":
		return *self.Arguments, nil
		return nil, fmt.Errorf("Field Arguments no set on Arguments %+v", self)

	case "executorData", "ExecutorData":
		return self.ExecutorData, nil
		return nil, fmt.Errorf("Field ExecutorData no set on ExecutorData %+v", self)

	case "labels", "Labels":
		return *self.Labels, nil
		return nil, fmt.Errorf("Field Labels no set on Labels %+v", self)

	case "autoAdvanceDeploySteps", "AutoAdvanceDeploySteps":
		return *self.AutoAdvanceDeploySteps, nil
		return nil, fmt.Errorf("Field AutoAdvanceDeploySteps no set on AutoAdvanceDeploySteps %+v", self)

	case "metadata", "Metadata":
		return *self.Metadata, nil
		return nil, fmt.Errorf("Field Metadata no set on Metadata %+v", self)

	case "customExecutorResources", "CustomExecutorResources":
		return self.CustomExecutorResources, nil
		return nil, fmt.Errorf("Field CustomExecutorResources no set on CustomExecutorResources %+v", self)

	case "taskEnv", "TaskEnv":
		return *self.TaskEnv, nil
		return nil, fmt.Errorf("Field TaskEnv no set on TaskEnv %+v", self)

	case "healthcheckMaxTotalTimeoutSeconds", "HealthcheckMaxTotalTimeoutSeconds":
		return *self.HealthcheckMaxTotalTimeoutSeconds, nil
		return nil, fmt.Errorf("Field HealthcheckMaxTotalTimeoutSeconds no set on HealthcheckMaxTotalTimeoutSeconds %+v", self)

	case "loadBalancerTemplate", "LoadBalancerTemplate":
		return *self.LoadBalancerTemplate, nil
		return nil, fmt.Errorf("Field LoadBalancerTemplate no set on LoadBalancerTemplate %+v", self)

	case "deployInstanceCountPerStep", "DeployInstanceCountPerStep":
		return *self.DeployInstanceCountPerStep, nil
		return nil, fmt.Errorf("Field DeployInstanceCountPerStep no set on DeployInstanceCountPerStep %+v", self)

	case "user", "User":
		return *self.User, nil
		return nil, fmt.Errorf("Field User no set on User %+v", self)

	case "customExecutorId", "CustomExecutorId":
		return *self.CustomExecutorId, nil
		return nil, fmt.Errorf("Field CustomExecutorId no set on CustomExecutorId %+v", self)

	case "customExecutorSource", "CustomExecutorSource":
		return *self.CustomExecutorSource, nil
		return nil, fmt.Errorf("Field CustomExecutorSource no set on CustomExecutorSource %+v", self)

	case "healthcheckProtocol", "HealthcheckProtocol":
		return *self.HealthcheckProtocol, nil
		return nil, fmt.Errorf("Field HealthcheckProtocol no set on HealthcheckProtocol %+v", self)

	case "skipHealthchecksOnDeploy", "SkipHealthchecksOnDeploy":
		return *self.SkipHealthchecksOnDeploy, nil
		return nil, fmt.Errorf("Field SkipHealthchecksOnDeploy no set on SkipHealthchecksOnDeploy %+v", self)

	case "considerHealthyAfterRunningForSeconds", "ConsiderHealthyAfterRunningForSeconds":
		return *self.ConsiderHealthyAfterRunningForSeconds, nil
		return nil, fmt.Errorf("Field ConsiderHealthyAfterRunningForSeconds no set on ConsiderHealthyAfterRunningForSeconds %+v", self)

	case "timestamp", "Timestamp":
		return *self.Timestamp, nil
		return nil, fmt.Errorf("Field Timestamp no set on Timestamp %+v", self)

	case "customExecutorCmd", "CustomExecutorCmd":
		return *self.CustomExecutorCmd, nil
		return nil, fmt.Errorf("Field CustomExecutorCmd no set on CustomExecutorCmd %+v", self)

	case "healthcheckTimeoutSeconds", "HealthcheckTimeoutSeconds":
		return *self.HealthcheckTimeoutSeconds, nil
		return nil, fmt.Errorf("Field HealthcheckTimeoutSeconds no set on HealthcheckTimeoutSeconds %+v", self)

	case "deployHealthTimeoutSeconds", "DeployHealthTimeoutSeconds":
		return *self.DeployHealthTimeoutSeconds, nil
		return nil, fmt.Errorf("Field DeployHealthTimeoutSeconds no set on DeployHealthTimeoutSeconds %+v", self)

	case "shell", "Shell":
		return *self.Shell, nil
		return nil, fmt.Errorf("Field Shell no set on Shell %+v", self)

	case "containerInfo", "ContainerInfo":
		return self.ContainerInfo, nil
		return nil, fmt.Errorf("Field ContainerInfo no set on ContainerInfo %+v", self)

	case "taskLabels", "TaskLabels":
		return *self.TaskLabels, nil
		return nil, fmt.Errorf("Field TaskLabels no set on TaskLabels %+v", self)

	case "healthcheckUri", "HealthcheckUri":
		return *self.HealthcheckUri, nil
		return nil, fmt.Errorf("Field HealthcheckUri no set on HealthcheckUri %+v", self)

	case "healthcheck", "Healthcheck":
		return self.Healthcheck, nil
		return nil, fmt.Errorf("Field Healthcheck no set on Healthcheck %+v", self)

	case "deployStepWaitTimeMs", "DeployStepWaitTimeMs":
		return *self.DeployStepWaitTimeMs, nil
		return nil, fmt.Errorf("Field DeployStepWaitTimeMs no set on DeployStepWaitTimeMs %+v", self)

	case "id", "Id":
		return *self.Id, nil
		return nil, fmt.Errorf("Field Id no set on Id %+v", self)

	case "mesosLabels", "MesosLabels":
		return *self.MesosLabels, nil
		return nil, fmt.Errorf("Field MesosLabels no set on MesosLabels %+v", self)

	case "healthcheckPortIndex", "HealthcheckPortIndex":
		return *self.HealthcheckPortIndex, nil
		return nil, fmt.Errorf("Field HealthcheckPortIndex no set on HealthcheckPortIndex %+v", self)

	case "loadBalancerPortIndex", "LoadBalancerPortIndex":
		return *self.LoadBalancerPortIndex, nil
		return nil, fmt.Errorf("Field LoadBalancerPortIndex no set on LoadBalancerPortIndex %+v", self)

	case "loadBalancerAdditionalRoutes", "LoadBalancerAdditionalRoutes":
		return *self.LoadBalancerAdditionalRoutes, nil
		return nil, fmt.Errorf("Field LoadBalancerAdditionalRoutes no set on LoadBalancerAdditionalRoutes %+v", self)

	}
}

func (self *SingularityDeploy) ClearField(name string) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityDeploy", name)

	case "loadBalancerGroups", "LoadBalancerGroups":
		self.LoadBalancerGroups = nil

	case "loadBalancerDomains", "LoadBalancerDomains":
		self.LoadBalancerDomains = nil

	case "loadBalancerUpstreamGroup", "LoadBalancerUpstreamGroup":
		self.LoadBalancerUpstreamGroup = nil

	case "resources", "Resources":
		self.Resources = nil

	case "command", "Command":
		self.Command = nil

	case "env", "Env":
		self.Env = nil

	case "mesosTaskLabels", "MesosTaskLabels":
		self.MesosTaskLabels = nil

	case "healthcheckMaxRetries", "HealthcheckMaxRetries":
		self.HealthcheckMaxRetries = nil

	case "maxTaskRetries", "MaxTaskRetries":
		self.MaxTaskRetries = nil

	case "loadBalancerOptions", "LoadBalancerOptions":
		self.LoadBalancerOptions = nil

	case "loadBalancerServiceIdOverride", "LoadBalancerServiceIdOverride":
		self.LoadBalancerServiceIdOverride = nil

	case "requestId", "RequestId":
		self.RequestId = nil

	case "version", "Version":
		self.Version = nil

	case "uris", "Uris":
		self.Uris = nil

	case "healthcheckIntervalSeconds", "HealthcheckIntervalSeconds":
		self.HealthcheckIntervalSeconds = nil

	case "serviceBasePath", "ServiceBasePath":
		self.ServiceBasePath = nil

	case "arguments", "Arguments":
		self.Arguments = nil

	case "executorData", "ExecutorData":
		self.ExecutorData = nil

	case "labels", "Labels":
		self.Labels = nil

	case "autoAdvanceDeploySteps", "AutoAdvanceDeploySteps":
		self.AutoAdvanceDeploySteps = nil

	case "metadata", "Metadata":
		self.Metadata = nil

	case "customExecutorResources", "CustomExecutorResources":
		self.CustomExecutorResources = nil

	case "taskEnv", "TaskEnv":
		self.TaskEnv = nil

	case "healthcheckMaxTotalTimeoutSeconds", "HealthcheckMaxTotalTimeoutSeconds":
		self.HealthcheckMaxTotalTimeoutSeconds = nil

	case "loadBalancerTemplate", "LoadBalancerTemplate":
		self.LoadBalancerTemplate = nil

	case "deployInstanceCountPerStep", "DeployInstanceCountPerStep":
		self.DeployInstanceCountPerStep = nil

	case "user", "User":
		self.User = nil

	case "customExecutorId", "CustomExecutorId":
		self.CustomExecutorId = nil

	case "customExecutorSource", "CustomExecutorSource":
		self.CustomExecutorSource = nil

	case "healthcheckProtocol", "HealthcheckProtocol":
		self.HealthcheckProtocol = nil

	case "skipHealthchecksOnDeploy", "SkipHealthchecksOnDeploy":
		self.SkipHealthchecksOnDeploy = nil

	case "considerHealthyAfterRunningForSeconds", "ConsiderHealthyAfterRunningForSeconds":
		self.ConsiderHealthyAfterRunningForSeconds = nil

	case "timestamp", "Timestamp":
		self.Timestamp = nil

	case "customExecutorCmd", "CustomExecutorCmd":
		self.CustomExecutorCmd = nil

	case "healthcheckTimeoutSeconds", "HealthcheckTimeoutSeconds":
		self.HealthcheckTimeoutSeconds = nil

	case "deployHealthTimeoutSeconds", "DeployHealthTimeoutSeconds":
		self.DeployHealthTimeoutSeconds = nil

	case "shell", "Shell":
		self.Shell = nil

	case "containerInfo", "ContainerInfo":
		self.ContainerInfo = nil

	case "taskLabels", "TaskLabels":
		self.TaskLabels = nil

	case "healthcheckUri", "HealthcheckUri":
		self.HealthcheckUri = nil

	case "healthcheck", "Healthcheck":
		self.Healthcheck = nil

	case "deployStepWaitTimeMs", "DeployStepWaitTimeMs":
		self.DeployStepWaitTimeMs = nil

	case "id", "Id":
		self.Id = nil

	case "mesosLabels", "MesosLabels":
		self.MesosLabels = nil

	case "healthcheckPortIndex", "HealthcheckPortIndex":
		self.HealthcheckPortIndex = nil

	case "loadBalancerPortIndex", "LoadBalancerPortIndex":
		self.LoadBalancerPortIndex = nil

	case "loadBalancerAdditionalRoutes", "LoadBalancerAdditionalRoutes":
		self.LoadBalancerAdditionalRoutes = nil

	}

	return nil
}

func (self *SingularityDeploy) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type SingularityDeployList []*SingularityDeploy

func (self *SingularityDeployList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityDeployList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityDeployList cannot copy the values from %#v", other)
}

func (list *SingularityDeployList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *SingularityDeployList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityDeployList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
