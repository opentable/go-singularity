package dtos

import "io"

type SingularityDeploy struct {
	Arguments                             StringList                `json:"arguments"`
	AutoAdvanceDeploySteps                bool                      `json:"autoAdvanceDeploySteps"`
	Command                               string                    `json:"command"`
	ConsiderHealthyAfterRunningForSeconds int64                     `json:"considerHealthyAfterRunningForSeconds"`
	ContainerInfo                         *SingularityContainerInfo `json:"containerInfo"`
	CustomExecutorCmd                     string                    `json:"customExecutorCmd"`
	CustomExecutorId                      string                    `json:"customExecutorId"`
	CustomExecutorResources               *Resources                `json:"customExecutorResources"`
	CustomExecutorSource                  string                    `json:"customExecutorSource"`
	CustomExecutorUser                    string                    `json:"customExecutorUser"`
	DeployHealthTimeoutSeconds            int64                     `json:"deployHealthTimeoutSeconds"`
	DeployInstanceCountPerStep            int32                     `json:"deployInstanceCountPerStep"`
	DeployStepWaitTimeMs                  int32                     `json:"deployStepWaitTimeMs"`
	Env                                   map[string]string         `json:"env"`
	ExecutorData                          *ExecutorData             `json:"executorData"`
	HealthcheckIntervalSeconds            int64                     `json:"healthcheckIntervalSeconds"`
	HealthcheckMaxRetries                 int32                     `json:"healthcheckMaxRetries"`
	HealthcheckMaxTotalTimeoutSeconds     int64                     `json:"healthcheckMaxTotalTimeoutSeconds"`
	HealthcheckPortIndex                  int32                     `json:"healthcheckPortIndex"`
	//	HealthcheckProtocol *HealthcheckProtocol `json:"healthcheckProtocol"`
	HealthcheckTimeoutSeconds int64             `json:"healthcheckTimeoutSeconds"`
	HealthcheckUri            string            `json:"healthcheckUri"`
	Id                        string            `json:"id"`
	Labels                    map[string]string `json:"labels"`
	LoadBalancerGroups        StringList        `json:"loadBalancerGroups"`
	//	LoadBalancerOptions *Map[string,Object] `json:"loadBalancerOptions"`
	LoadBalancerPortIndex int32             `json:"loadBalancerPortIndex"`
	MaxTaskRetries        int32             `json:"maxTaskRetries"`
	Metadata              map[string]string `json:"metadata"`
	RequestId             string            `json:"requestId"`
	//	Resources *com.hubspot.mesos.Resources `json:"resources"`
	ServiceBasePath          string     `json:"serviceBasePath"`
	SkipHealthchecksOnDeploy bool       `json:"skipHealthchecksOnDeploy"`
	Timestamp                int64      `json:"timestamp"`
	Uris                     StringList `json:"uris"`
	Version                  string     `json:"version"`
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

func (list *SingularityDeployList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
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
	return FormatJSON(list)
}
