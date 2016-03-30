package dtos

import "io"

type OfferIDOrBuilder struct {
	Value      string
	ValueBytes *ByteString
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

type OfferIDOrBuilderList []*OfferIDOrBuilder

func (list OfferIDOrBuilderList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list OfferIDOrBuilderList) FormatText() string {
	text := []byte{}
	for _, dto := range list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list OfferIDOrBuilderList) FormatJSON() string {
	return FormatJSON(list)
}
