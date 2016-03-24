package dtos

import "io"

type Deploy struct {
	Arguments []string
	AutoAdvanceDeploySteps bool
	Command string
	ConsiderHealthyAfterRunningForSeconds int64
	ContainerInfo ContainerInfo
	CustomExecutorCmd string
	CustomExecutorId string
	CustomExecutorResources Resources
	CustomExecutorSource string
	CustomExecutorUser string
	DeployHealthTimeoutSeconds int64
	DeployInstanceCountPerStep int32
	DeployStepWaitTimeMs int32
//	Env Map[string,string]
	ExecutorData ExecutorData
	HealthcheckIntervalSeconds int64
	HealthcheckMaxRetries int32
	HealthcheckMaxTotalTimeoutSeconds int64
	HealthcheckPortIndex int32
//	HealthcheckProtocol HealthcheckProtocol
	HealthcheckTimeoutSeconds int64
	HealthcheckUri string
	Id string
//	Labels Map[string,string]
	LoadBalancerGroups []string
//	LoadBalancerOptions Map[string,Object]
	LoadBalancerPortIndex int32
	MaxTaskRetries int32
//	Metadata Map[string,string]
	RequestId string
//	Resources com.hubspot.mesos.Resources
	ServiceBasePath string
	SkipHealthchecksOnDeploy bool
	Timestamp int64
	Uris []string
	Version string
}

func (self *Deploy) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *Deploy) FormatText() string {
	return FormatText(self)
}

func (self *Deploy) FormatJSON() string {
	return FormatJSON(self)
}
