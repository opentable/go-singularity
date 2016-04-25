package dtos

import "io"

type Descriptor struct {
	ContainingType *Descriptor `json:"containingType"`
	//	EnumTypes *List[EnumDescriptor] `json:"enumTypes"`
	//	Extensions *List[FieldDescriptor] `json:"extensions"`
	//	Fields *List[FieldDescriptor] `json:"fields"`
	File        *FileDescriptor `json:"file"`
	FullName    string          `json:"fullName"`
	Index       int32           `json:"index"`
	Name        string          `json:"name"`
	NestedTypes DescriptorList  `json:"nestedTypes"`
	Options     *MessageOptions `json:"options"`
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

func (list *DescriptorList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *DescriptorList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *DescriptorList) FormatJSON() string {
	return FormatJSON(list)
}
