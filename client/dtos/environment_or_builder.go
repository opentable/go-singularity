package dtos

import "io"

type EnvironmentOrBuilder struct {
	VariablesCount int32
	//	VariablesList *List[Variable]
	//	VariablesOrBuilderList *List[? extends org.apache.mesos.Protos$Environment$VariableOrBuilder]
}

func (self *EnvironmentOrBuilder) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *EnvironmentOrBuilder) FormatText() string {
	return FormatText(self)
}

func (self *EnvironmentOrBuilder) FormatJSON() string {
	return FormatJSON(self)
}
