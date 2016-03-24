package dtos

import "io"

type OfferIDOrBuilder struct {
	Value string
	ValueBytes ByteString
}

func (self *OfferIDOrBuilder) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *OfferIDOrBuilder) FormatText() string {
	return FormatText(self)
}

func (self *OfferIDOrBuilder) FormatJSON() string {
	return FormatJSON(self)
}
