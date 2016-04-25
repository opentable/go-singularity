package dtos

import "io"

type FileDescriptor struct {
	Dependencies FileDescriptorList `json:"dependencies"`
	//	EnumTypes *List[EnumDescriptor] `json:"enumTypes"`
	//	Extensions *List[FieldDescriptor] `json:"extensions"`
	MessageTypes       DescriptorList     `json:"messageTypes"`
	Name               string             `json:"name"`
	Options            *FileOptions       `json:"options"`
	Package            string             `json:"package"`
	PublicDependencies FileDescriptorList `json:"publicDependencies"`
	//	Services *List[ServiceDescriptor] `json:"services"`

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
