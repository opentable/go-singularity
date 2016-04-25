package dtos

import "io"

type Offer struct {
	//	AllFields *Map[FieldDescriptor,Object] `json:"allFields"`
	AttributesCount int32 `json:"attributesCount"`
	//	AttributesList *List[Attribute] `json:"attributesList"`
	//	AttributesOrBuilderList *List[? extends org.apache.mesos.Protos$AttributeOrBuilder] `json:"attributesOrBuilderList"`
	DefaultInstanceForType *Offer         `json:"defaultInstanceForType"`
	DescriptorForType      *Descriptor    `json:"descriptorForType"`
	ExecutorIdsCount       int32          `json:"executorIdsCount"`
	ExecutorIdsList        ExecutorIDList `json:"executorIdsList"`
	//	ExecutorIdsOrBuilderList *List[? extends org.apache.mesos.Protos$ExecutorIDOrBuilder] `json:"executorIdsOrBuilderList"`
	FrameworkId               *FrameworkID          `json:"frameworkId"`
	FrameworkIdOrBuilder      *FrameworkIDOrBuilder `json:"frameworkIdOrBuilder"`
	Hostname                  string                `json:"hostname"`
	HostnameBytes             *ByteString           `json:"hostnameBytes"`
	Id                        *OfferID              `json:"id"`
	IdOrBuilder               *OfferIDOrBuilder     `json:"idOrBuilder"`
	InitializationErrorString string                `json:"initializationErrorString"`
	Initialized               bool                  `json:"initialized"`
	//	ParserForType *com.google.protobuf.Parser&lt;org.apache.mesos.Protos$Offer&gt; `json:"parserForType"`
	ResourcesCount int32 `json:"resourcesCount"`
	//	ResourcesList *List[Resource] `json:"resourcesList"`
	//	ResourcesOrBuilderList *List[? extends org.apache.mesos.Protos$ResourceOrBuilder] `json:"resourcesOrBuilderList"`
	SerializedSize   int32             `json:"serializedSize"`
	SlaveId          *SlaveID          `json:"slaveId"`
	SlaveIdOrBuilder *SlaveIDOrBuilder `json:"slaveIdOrBuilder"`
	UnknownFields    *UnknownFieldSet  `json:"unknownFields"`
}

func (self *Offer) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *Offer) FormatText() string {
	return FormatText(self)
}

func (self *Offer) FormatJSON() string {
	return FormatJSON(self)
}

type OfferList []*Offer

func (list *OfferList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *OfferList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *OfferList) FormatJSON() string {
	return FormatJSON(list)
}
