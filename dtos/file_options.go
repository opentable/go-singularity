package dtos

import "io"

type FileOptionsOptimizeMode string

const (
	FileOptionsOptimizeModeSPEED        FileOptionsOptimizeMode = "SPEED"
	FileOptionsOptimizeModeCODE_SIZE    FileOptionsOptimizeMode = "CODE_SIZE"
	FileOptionsOptimizeModeLITE_RUNTIME FileOptionsOptimizeMode = "LITE_RUNTIME"
)

type FileOptions struct {
	//	AllFields *Map[FieldDescriptor,Object] `json:"allFields"`
	CcGenericServices         bool                    `json:"ccGenericServices"`
	DefaultInstanceForType    *FileOptions            `json:"defaultInstanceForType"`
	DescriptorForType         *Descriptor             `json:"descriptorForType"`
	GoPackage                 string                  `json:"goPackage"`
	GoPackageBytes            *ByteString             `json:"goPackageBytes"`
	InitializationErrorString string                  `json:"initializationErrorString"`
	Initialized               bool                    `json:"initialized"`
	JavaGenerateEqualsAndHash bool                    `json:"javaGenerateEqualsAndHash"`
	JavaGenericServices       bool                    `json:"javaGenericServices"`
	JavaMultipleFiles         bool                    `json:"javaMultipleFiles"`
	JavaOuterClassname        string                  `json:"javaOuterClassname"`
	JavaOuterClassnameBytes   *ByteString             `json:"javaOuterClassnameBytes"`
	JavaPackage               string                  `json:"javaPackage"`
	JavaPackageBytes          *ByteString             `json:"javaPackageBytes"`
	OptimizeFor               FileOptionsOptimizeMode `json:"optimizeFor"`
	//	ParserForType *com.google.protobuf.Parser&lt;com.google.protobuf.DescriptorProtos$FileOptions&gt; `json:"parserForType"`
	PyGenericServices        bool  `json:"pyGenericServices"`
	SerializedSize           int32 `json:"serializedSize"`
	UninterpretedOptionCount int32 `json:"uninterpretedOptionCount"`
	//	UninterpretedOptionList *List[UninterpretedOption] `json:"uninterpretedOptionList"`
	//	UninterpretedOptionOrBuilderList *List[? extends com.google.protobuf.DescriptorProtos$UninterpretedOptionOrBuilder] `json:"uninterpretedOptionOrBuilderList"`
	UnknownFields *UnknownFieldSet `json:"unknownFields"`
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
