package dtos

import "io"

type DiscoveryInfoOrBuilderVisibility string

const (
	DiscoveryInfoOrBuilderVisibilityFRAMEWORK DiscoveryInfoOrBuilderVisibility = "FRAMEWORK"
	DiscoveryInfoOrBuilderVisibilityCLUSTER   DiscoveryInfoOrBuilderVisibility = "CLUSTER"
	DiscoveryInfoOrBuilderVisibilityEXTERNAL  DiscoveryInfoOrBuilderVisibility = "EXTERNAL"
)

type DiscoveryInfoOrBuilder struct {
	Environment      string                           `json:"environment"`
	EnvironmentBytes *ByteString                      `json:"environmentBytes"`
	Labels           *Labels                          `json:"labels"`
	LabelsOrBuilder  *LabelsOrBuilder                 `json:"labelsOrBuilder"`
	Location         string                           `json:"location"`
	LocationBytes    *ByteString                      `json:"locationBytes"`
	Name             string                           `json:"name"`
	NameBytes        *ByteString                      `json:"nameBytes"`
	Ports            *Ports                           `json:"ports"`
	PortsOrBuilder   *PortsOrBuilder                  `json:"portsOrBuilder"`
	Version          string                           `json:"version"`
	VersionBytes     *ByteString                      `json:"versionBytes"`
	Visibility       DiscoveryInfoOrBuilderVisibility `json:"visibility"`
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
