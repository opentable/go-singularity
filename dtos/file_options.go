package dtos

import "io"

type FileOptions struct {
	//	AllFields *Map[FieldDescriptor,Object]
	CcGenericServices         bool
	DefaultInstanceForType    *FileOptions
	DescriptorForType         *Descriptor
	GoPackage                 string
	GoPackageBytes            *ByteString
	InitializationErrorString string
	Initialized               bool
	JavaGenerateEqualsAndHash bool
	JavaGenericServices       bool
	JavaMultipleFiles         bool
	JavaOuterClassname        string
	JavaOuterClassnameBytes   *ByteString
	JavaPackage               string
	JavaPackageBytes          *ByteString
	//	OptimizeFor *OptimizeMode
	//	ParserForType *com.google.protobuf.Parser&lt;com.google.protobuf.DescriptorProtos$FileOptions&gt;
	PyGenericServices        bool
	SerializedSize           int32
	UninterpretedOptionCount int32
	//	UninterpretedOptionList *List[UninterpretedOption]
	//	UninterpretedOptionOrBuilderList *List[? extends com.google.protobuf.DescriptorProtos$UninterpretedOptionOrBuilder]
	UnknownFields *UnknownFieldSet
}

func (self *FileOptions) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *FileOptions) FormatText() string {
	return FormatText(self)
}

func (self *FileOptions) FormatJSON() string {
	return FormatJSON(self)
}

type FileOptionsList []*FileOptions

func (list *FileOptionsList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *FileOptionsList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *FileOptionsList) FormatJSON() string {
	return FormatJSON(list)
}
