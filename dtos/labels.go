package dtos

import "io"

type Labels struct {
	//	AllFields *Map[FieldDescriptor,Object] `json:"allFields"`
	DefaultInstanceForType    *Labels     `json:"defaultInstanceForType"`
	DescriptorForType         *Descriptor `json:"descriptorForType"`
	InitializationErrorString string      `json:"initializationErrorString"`
	Initialized               bool        `json:"initialized"`
	LabelsCount               int32       `json:"labelsCount"`
	//	LabelsList *List[Label] `json:"labelsList"`
	//	LabelsOrBuilderList *List[? extends org.apache.mesos.Protos$LabelOrBuilder] `json:"labelsOrBuilderList"`
	//	ParserForType *com.google.protobuf.Parser&lt;org.apache.mesos.Protos$Labels&gt; `json:"parserForType"`
	SerializedSize int32            `json:"serializedSize"`
	UnknownFields  *UnknownFieldSet `json:"unknownFields"`
}

func (self *Labels) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *Labels) FormatText() string {
	return FormatText(self)
}

func (self *Labels) FormatJSON() string {
	return FormatJSON(self)
}

type LabelsList []*Labels

func (list *LabelsList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *LabelsList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *LabelsList) FormatJSON() string {
	return FormatJSON(list)
}
