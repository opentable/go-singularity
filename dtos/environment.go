package dtos

import "io"

type Environment struct {
	//	AllFields *Map[FieldDescriptor,Object] `json:"allFields"`
	DefaultInstanceForType    *Environment `json:"defaultInstanceForType"`
	DescriptorForType         *Descriptor  `json:"descriptorForType"`
	InitializationErrorString string       `json:"initializationErrorString"`
	Initialized               bool         `json:"initialized"`
	//	ParserForType *com.google.protobuf.Parser&lt;org.apache.mesos.Protos$Environment&gt; `json:"parserForType"`
	SerializedSize int32            `json:"serializedSize"`
	UnknownFields  *UnknownFieldSet `json:"unknownFields"`
	VariablesCount int32            `json:"variablesCount"`
	//	VariablesList *List[Variable] `json:"variablesList"`
	//	VariablesOrBuilderList *List[? extends org.apache.mesos.Protos$Environment$VariableOrBuilder] `json:"variablesOrBuilderList"`

}

func (self *Environment) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *Environment) FormatText() string {
	return FormatText(self)
}

func (self *Environment) FormatJSON() string {
	return FormatJSON(self)
}

type EnvironmentList []*Environment

func (list *EnvironmentList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *EnvironmentList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *EnvironmentList) FormatJSON() string {
	return FormatJSON(list)
}
