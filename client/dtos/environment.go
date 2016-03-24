package dtos

import "io"

type Environment struct {
//	AllFields Map[FieldDescriptor,Object]
	DefaultInstanceForType Environment
	DescriptorForType Descriptor
	InitializationErrorString string
	Initialized bool
//	ParserForType com.google.protobuf.Parser&lt;org.apache.mesos.Protos$Environment&gt;
	SerializedSize int32
	UnknownFields UnknownFieldSet
	VariablesCount int32
//	VariablesList List[Variable]
//	VariablesOrBuilderList List[? extends org.apache.mesos.Protos$Environment$VariableOrBuilder]
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
