package dtos

import "io"

type FileDescriptor struct {
	//	Dependencies *List[FileDescriptor]
	//	EnumTypes *List[EnumDescriptor]
	//	Extensions *List[FieldDescriptor]
	//	MessageTypes *List[Descriptor]
	Name    string
	Options *FileOptions
	Package string
	//	PublicDependencies *List[FileDescriptor]
	//	Services *List[ServiceDescriptor]
}

func (self *FileDescriptor) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *FileDescriptor) FormatText() string {
	return FormatText(self)
}

func (self *FileDescriptor) FormatJSON() string {
	return FormatJSON(self)
}

type FileDescriptorList []*FileDescriptor

func (list *FileDescriptorList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *FileDescriptorList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *FileDescriptorList) FormatJSON() string {
	return FormatJSON(list)
}
