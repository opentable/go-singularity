package dtos

import "io"

type Offer struct {
	//	AllFields *Map[FieldDescriptor,Object]
	AttributesCount int32
	//	AttributesList *List[Attribute]
	//	AttributesOrBuilderList *List[? extends org.apache.mesos.Protos$AttributeOrBuilder]
	DefaultInstanceForType *Offer
	DescriptorForType      *Descriptor
	ExecutorIdsCount       int32
	//	ExecutorIdsList *List[ExecutorID]
	//	ExecutorIdsOrBuilderList *List[? extends org.apache.mesos.Protos$ExecutorIDOrBuilder]
	FrameworkId               *FrameworkID
	FrameworkIdOrBuilder      *FrameworkIDOrBuilder
	Hostname                  string
	HostnameBytes             *ByteString
	Id                        *OfferID
	IdOrBuilder               *OfferIDOrBuilder
	InitializationErrorString string
	Initialized               bool
	//	ParserForType *com.google.protobuf.Parser&lt;org.apache.mesos.Protos$Offer&gt;
	ResourcesCount int32
	//	ResourcesList *List[Resource]
	//	ResourcesOrBuilderList *List[? extends org.apache.mesos.Protos$ResourceOrBuilder]
	SerializedSize   int32
	SlaveId          *SlaveID
	SlaveIdOrBuilder *SlaveIDOrBuilder
	UnknownFields    *UnknownFieldSet
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
