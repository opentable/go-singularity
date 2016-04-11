package dtos

import "io"

type DiscoveryInfoOrBuilder struct {
	Environment      string
	EnvironmentBytes *ByteString
	Labels           *Labels
	LabelsOrBuilder  *LabelsOrBuilder
	Location         string
	LocationBytes    *ByteString
	Name             string
	NameBytes        *ByteString
	Ports            *Ports
	PortsOrBuilder   *PortsOrBuilder
	Version          string
	VersionBytes     *ByteString
	//	Visibility *Visibility
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

type DiscoveryInfoOrBuilderList []*DiscoveryInfoOrBuilder

func (list *DiscoveryInfoOrBuilderList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *DiscoveryInfoOrBuilderList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *DiscoveryInfoOrBuilderList) FormatJSON() string {
	return FormatJSON(list)
}
