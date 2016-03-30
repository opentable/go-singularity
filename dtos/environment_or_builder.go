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

type EnvironmentOrBuilderList []*EnvironmentOrBuilder

func (list EnvironmentOrBuilderList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list EnvironmentOrBuilderList) FormatText() string {
	text := []byte{}
	for _, dto := range list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list EnvironmentOrBuilderList) FormatJSON() string {
	return FormatJSON(list)
}
