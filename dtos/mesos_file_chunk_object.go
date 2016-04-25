package dtos

import "io"

type MesosFileChunkObject struct {
	Data       string `json:"data"`
	NextOffset int64  `json:"nextOffset"`
	Offset     int64  `json:"offset"`
}

func (self *MesosFileChunkObject) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *MesosFileChunkObject) FormatText() string {
	return FormatText(self)
}

func (self *MesosFileChunkObject) FormatJSON() string {
	return FormatJSON(self)
}

type MesosFileChunkObjectList []*MesosFileChunkObject

func (list *MesosFileChunkObjectList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *MesosFileChunkObjectList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *MesosFileChunkObjectList) FormatJSON() string {
	return FormatJSON(list)
}
