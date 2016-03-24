package dtos

import "io"

type FileDescriptor struct {
//	Dependencies List[FileDescriptor]
//	EnumTypes List[EnumDescriptor]
//	Extensions List[FieldDescriptor]
//	MessageTypes List[Descriptor]
	Name string
	Options FileOptions
	Package string
//	PublicDependencies List[FileDescriptor]
//	Services List[ServiceDescriptor]
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
