package dtos

import "io"

type Resources struct {
	Cpus     float64
	MemoryMb float64
	NumPorts int32
}

func (self *Resources) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *Resources) FormatText() string {
	return FormatText(self)
}

func (self *Resources) FormatJSON() string {
	return FormatJSON(self)
}

type ResourcesList []*Resources

func (list *ResourcesList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *ResourcesList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *ResourcesList) FormatJSON() string {
	return FormatJSON(list)
}
