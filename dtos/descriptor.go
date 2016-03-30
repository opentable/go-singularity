package dtos

import "io"

type Descriptor struct {
	ContainingType *Descriptor
	//	EnumTypes *List[EnumDescriptor]
	//	Extensions *List[FieldDescriptor]
	//	Fields *List[FieldDescriptor]
	File     *FileDescriptor
	FullName string
	Index    int32
	Name     string
	//	NestedTypes *List[Descriptor]
	Options *MessageOptions
}

func (self *Descriptor) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *Descriptor) FormatText() string {
	return FormatText(self)
}

func (self *Descriptor) FormatJSON() string {
	return FormatJSON(self)
}

type DescriptorList []*Descriptor

func (list DescriptorList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list DescriptorList) FormatText() string {
	text := []byte{}
	for _, dto := range list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list DescriptorList) FormatJSON() string {
	return FormatJSON(list)
}
