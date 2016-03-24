package dtos

import "io"

type DiscoveryInfoOrBuilder struct {
	Environment string
	EnvironmentBytes ByteString
	Labels Labels
	LabelsOrBuilder LabelsOrBuilder
	Location string
	LocationBytes ByteString
	Name string
	NameBytes ByteString
	Ports Ports
	PortsOrBuilder PortsOrBuilder
	Version string
	VersionBytes ByteString
//	Visibility Visibility
}

func (self *DiscoveryInfoOrBuilder) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *DiscoveryInfoOrBuilder) FormatText() string {
	return FormatText(self)
}

func (self *DiscoveryInfoOrBuilder) FormatJSON() string {
	return FormatJSON(self)
}
