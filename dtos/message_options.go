package dtos

import "io"

type MessageOptions struct {
	//	AllFields *Map[FieldDescriptor,Object]
	DefaultInstanceForType       *MessageOptions
	DescriptorForType            *Descriptor
	InitializationErrorString    string
	Initialized                  bool
	MessageSetWireFormat         bool
	NoStandardDescriptorAccessor bool
	//	ParserForType *com.google.protobuf.Parser&lt;com.google.protobuf.DescriptorProtos$MessageOptions&gt;
	SerializedSize           int32
	UninterpretedOptionCount int32
	//	UninterpretedOptionList *List[UninterpretedOption]
	//	UninterpretedOptionOrBuilderList *List[? extends com.google.protobuf.DescriptorProtos$UninterpretedOptionOrBuilder]
	UnknownFields *UnknownFieldSet
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
