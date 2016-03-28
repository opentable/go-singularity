package dtos

import "io"

type SingularityDeploy struct {
	Arguments                             string
	AutoAdvanceDeploySteps                bool
	Command                               string
	ConsiderHealthyAfterRunningForSeconds int64
	ContainerInfo                         *SingularityContainerInfo
	CustomExecutorCmd                     string
	CustomExecutorId                      string
	CustomExecutorResources               *Resources
	CustomExecutorSource                  string
	CustomExecutorUser                    string
	DeployHealthTimeoutSeconds            int64
	DeployInstanceCountPerStep            int32
	DeployStepWaitTimeMs                  int32
	//	Env *Map[string,string]
	ExecutorData                      *ExecutorData
	HealthcheckIntervalSeconds        int64
	HealthcheckMaxRetries             int32
	HealthcheckMaxTotalTimeoutSeconds int64
	HealthcheckPortIndex              int32
	//	HealthcheckProtocol *HealthcheckProtocol
	HealthcheckTimeoutSeconds int64
	HealthcheckUri            string
	Id                        string
	//	Labels *Map[string,string]
	LoadBalancerGroups string
	//	LoadBalancerOptions *Map[string,Object]
	LoadBalancerPortIndex int32
	MaxTaskRetries        int32
	//	Metadata *Map[string,string]
	RequestId string
	//	Resources *com.hubspot.mesos.Resources
	ServiceBasePath          string
	SkipHealthchecksOnDeploy bool
	Timestamp                int64
	Uris                     string
	Version                  string
}

func (self *SingularityDeploy) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SingularityDeploy) FormatText() string {
	return FormatText(self)
}

func (self *SingularityDeploy) FormatJSON() string {
	return FormatJSON(self)
}
