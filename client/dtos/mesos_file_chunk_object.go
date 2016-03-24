package dtos

import "io"

type MesosFileChunkObject struct {
	Data string
	NextOffset int64
	Offset int64
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
