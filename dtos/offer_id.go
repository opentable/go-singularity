package dtos

import "io"

type OfferID struct {
	//	AllFields *Map[FieldDescriptor,Object]
	DefaultInstanceForType    *OfferID
	DescriptorForType         *Descriptor
	InitializationErrorString string
	Initialized               bool
	//	ParserForType *com.google.protobuf.Parser&lt;org.apache.mesos.Protos$OfferID&gt;
	SerializedSize int32
	UnknownFields  *UnknownFieldSet
	Value          string
	ValueBytes     *ByteString
}

func (self *OfferID) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *OfferID) FormatText() string {
	return FormatText(self)
}

func (self *OfferID) FormatJSON() string {
	return FormatJSON(self)
}

type OfferIDList []*OfferID

func (list OfferIDList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list OfferIDList) FormatText() string {
	text := []byte{}
	for _, dto := range list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list OfferIDList) FormatJSON() string {
	return FormatJSON(list)
}
