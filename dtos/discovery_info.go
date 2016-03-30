package dtos

import "io"

type DiscoveryInfo struct {
	//	AllFields *Map[FieldDescriptor,Object]
	DefaultInstanceForType    *DiscoveryInfo
	DescriptorForType         *Descriptor
	Environment               string
	EnvironmentBytes          *ByteString
	InitializationErrorString string
	Initialized               bool
	Labels                    *Labels
	LabelsOrBuilder           *LabelsOrBuilder
	Location                  string
	LocationBytes             *ByteString
	Name                      string
	NameBytes                 *ByteString
	//	ParserForType *com.google.protobuf.Parser&lt;org.apache.mesos.Protos$DiscoveryInfo&gt;
	Ports          *Ports
	PortsOrBuilder *PortsOrBuilder
	SerializedSize int32
	UnknownFields  *UnknownFieldSet
	Version        string
	VersionBytes   *ByteString
	//	Visibility *Visibility
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

func (list DiscoveryInfoList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list DiscoveryInfoList) FormatText() string {
	text := []byte{}
	for _, dto := range list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list DiscoveryInfoList) FormatJSON() string {
	return FormatJSON(list)
}
