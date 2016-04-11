package dtos

import "io"

type Labels struct {
	//	AllFields *Map[FieldDescriptor,Object]
	DefaultInstanceForType    *Labels
	DescriptorForType         *Descriptor
	InitializationErrorString string
	Initialized               bool
	LabelsCount               int32
	//	LabelsList *List[Label]
	//	LabelsOrBuilderList *List[? extends org.apache.mesos.Protos$LabelOrBuilder]
	//	ParserForType *com.google.protobuf.Parser&lt;org.apache.mesos.Protos$Labels&gt;
	SerializedSize int32
	UnknownFields  *UnknownFieldSet
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
