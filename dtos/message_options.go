package dtos

import "io"

type MessageOptions struct {
	//	AllFields *Map[FieldDescriptor,Object] `json:"allFields"`
	DefaultInstanceForType       *MessageOptions `json:"defaultInstanceForType"`
	DescriptorForType            *Descriptor     `json:"descriptorForType"`
	InitializationErrorString    string          `json:"initializationErrorString"`
	Initialized                  bool            `json:"initialized"`
	MessageSetWireFormat         bool            `json:"messageSetWireFormat"`
	NoStandardDescriptorAccessor bool            `json:"noStandardDescriptorAccessor"`
	//	ParserForType *com.google.protobuf.Parser&lt;com.google.protobuf.DescriptorProtos$MessageOptions&gt; `json:"parserForType"`
	SerializedSize           int32 `json:"serializedSize"`
	UninterpretedOptionCount int32 `json:"uninterpretedOptionCount"`
	//	UninterpretedOptionList *List[UninterpretedOption] `json:"uninterpretedOptionList"`
	//	UninterpretedOptionOrBuilderList *List[? extends com.google.protobuf.DescriptorProtos$UninterpretedOptionOrBuilder] `json:"uninterpretedOptionOrBuilderList"`
	UnknownFields *UnknownFieldSet `json:"unknownFields"`
}

func (self *MessageOptions) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *MessageOptions) FormatText() string {
	return FormatText(self)
}

func (self *MessageOptions) FormatJSON() string {
	return FormatJSON(self)
}

type MessageOptionsList []*MessageOptions

func (list *MessageOptionsList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *MessageOptionsList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *MessageOptionsList) FormatJSON() string {
	return FormatJSON(list)
}
