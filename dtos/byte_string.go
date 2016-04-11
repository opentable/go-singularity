package dtos

import "io"

type ByteString struct {
	Empty     bool
	ValidUtf8 bool
}

func (self *ByteString) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *ByteString) FormatText() string {
	return FormatText(self)
}

func (self *ByteString) FormatJSON() string {
	return FormatJSON(self)
}

type ByteStringList []*ByteString

func (list *ByteStringList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *ByteStringList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *ByteStringList) FormatJSON() string {
	return FormatJSON(list)
}
