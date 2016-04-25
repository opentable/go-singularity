package dtos

import "io"

type DiscoveryInfoVisibility string

const (
	DiscoveryInfoVisibilityFRAMEWORK DiscoveryInfoVisibility = "FRAMEWORK"
	DiscoveryInfoVisibilityCLUSTER   DiscoveryInfoVisibility = "CLUSTER"
	DiscoveryInfoVisibilityEXTERNAL  DiscoveryInfoVisibility = "EXTERNAL"
)

type DiscoveryInfo struct {
	//	AllFields *Map[FieldDescriptor,Object] `json:"allFields"`
	DefaultInstanceForType    *DiscoveryInfo   `json:"defaultInstanceForType"`
	DescriptorForType         *Descriptor      `json:"descriptorForType"`
	Environment               string           `json:"environment"`
	EnvironmentBytes          *ByteString      `json:"environmentBytes"`
	InitializationErrorString string           `json:"initializationErrorString"`
	Initialized               bool             `json:"initialized"`
	Labels                    *Labels          `json:"labels"`
	LabelsOrBuilder           *LabelsOrBuilder `json:"labelsOrBuilder"`
	Location                  string           `json:"location"`
	LocationBytes             *ByteString      `json:"locationBytes"`
	Name                      string           `json:"name"`
	NameBytes                 *ByteString      `json:"nameBytes"`
	//	ParserForType *com.google.protobuf.Parser&lt;org.apache.mesos.Protos$DiscoveryInfo&gt; `json:"parserForType"`
	Ports          *Ports                  `json:"ports"`
	PortsOrBuilder *PortsOrBuilder         `json:"portsOrBuilder"`
	SerializedSize int32                   `json:"serializedSize"`
	UnknownFields  *UnknownFieldSet        `json:"unknownFields"`
	Version        string                  `json:"version"`
	VersionBytes   *ByteString             `json:"versionBytes"`
	Visibility     DiscoveryInfoVisibility `json:"visibility"`
}

func (self *DiscoveryInfo) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *DiscoveryInfo) FormatText() string {
	return FormatText(self)
}

func (self *DiscoveryInfo) FormatJSON() string {
	return FormatJSON(self)
}

type DiscoveryInfoList []*DiscoveryInfo

func (list *DiscoveryInfoList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *DiscoveryInfoList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *DiscoveryInfoList) FormatJSON() string {
	return FormatJSON(list)
}
