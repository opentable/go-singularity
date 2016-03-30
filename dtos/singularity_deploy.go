package dtos

import "io"

type SingularityDeploy struct {
	Arguments                             StringList
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
	LoadBalancerGroups StringList
	//	LoadBalancerOptions *Map[string,Object]
	LoadBalancerPortIndex int32
	MaxTaskRetries        int32
	//	Metadata *Map[string,string]
	RequestId string
	//	Resources *com.hubspot.mesos.Resources
	ServiceBasePath          string
	SkipHealthchecksOnDeploy bool
	Timestamp                int64
	Uris                     StringList
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

type SingularityDeployList []*SingularityDeploy

func (list SingularityDeployList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list SingularityDeployList) FormatText() string {
	text := []byte{}
	for _, dto := range list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list SingularityDeployList) FormatJSON() string {
	return FormatJSON(list)
}
