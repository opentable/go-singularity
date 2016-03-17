package dtos

import "io"

type Deploy struct {
	//Resources [com.hubspot.mesos.Resources]
	//ContainerInfo [SingularityContainerInfo]
	//HealthcheckProtocol [HealthcheckProtocol]
	//LoadBalancerGroups [Set]
	//ExecutorData [ExecutorData]
	//CustomExecutorResources [Resources]

	RequestId string
	Metadata  map[string]string
	Labels    map[string]string
	Version   string
	Id        string

	Uris            []string
	ServiceBasePath string

	Command   string
	Arguments []string
	Env       map[string]string

	Timestamp uint64

	CustomExecutorId     string
	CustomExecutorUser   string
	CustomExecutorSource string
	CustomExecutorCmd    string

	HealthcheckTimeoutSeconds             uint64
	HealthcheckMaxRetries                 int
	HealthcheckMaxTotalTimeoutSeconds     uint64
	HealthcheckUri                        string
	SkipHealthchecksOnDeploy              bool
	ConsiderHealthyAfterRunningForSeconds uint64
	DeployHealthTimeoutSeconds            uint64
	HealthcheckIntervalSeconds            uint64

	LoadBalancerOptions map[string]interface{}
}

type Deploys []Deploy

func (deps *Deploys) Get(client APIClient) (err error) {
	return client.APIGet("deploys", deps)
}

func (deps *Deploys) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, deps)
	return
}

func (deps *Deploys) FormatText() string {
	return FormatText(deps)
}

func (deps *Deploys) FormatJSON() string {
	return FormatJSON(deps)
}
