package dtos

import "io"

type UnknownFieldSet struct {
	DefaultInstanceForType *UnknownFieldSet `json:"defaultInstanceForType"`
	Initialized            bool             `json:"initialized"`
	//	ParserForType *Parser `json:"parserForType"`
	SerializedSize             int32 `json:"serializedSize"`
	SerializedSizeAsMessageSet int32 `json:"serializedSizeAsMessageSet"`
}

func (self *UnknownFieldSet) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *UnknownFieldSet) FormatText() string {
	return FormatText(self)
}

func (self *UnknownFieldSet) FormatJSON() string {
	return FormatJSON(self)
}

type UnknownFieldSetList []*UnknownFieldSet

func (list *UnknownFieldSetList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *UnknownFieldSetList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *UnknownFieldSetList) FormatJSON() string {
	return FormatJSON(list)
}
