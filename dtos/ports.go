package dtos

import "io"

type Ports struct {
	//	AllFields *Map[FieldDescriptor,Object] `json:"allFields"`
	DefaultInstanceForType    *Ports      `json:"defaultInstanceForType"`
	DescriptorForType         *Descriptor `json:"descriptorForType"`
	InitializationErrorString string      `json:"initializationErrorString"`
	Initialized               bool        `json:"initialized"`
	//	ParserForType *com.google.protobuf.Parser&lt;org.apache.mesos.Protos$Ports&gt; `json:"parserForType"`
	PortsCount int32 `json:"portsCount"`
	//	PortsList *List[Port] `json:"portsList"`
	//	PortsOrBuilderList *List[? extends org.apache.mesos.Protos$PortOrBuilder] `json:"portsOrBuilderList"`
	SerializedSize int32            `json:"serializedSize"`
	UnknownFields  *UnknownFieldSet `json:"unknownFields"`
}

func (self *Ports) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *Ports) FormatText() string {
	return FormatText(self)
}

func (self *Ports) FormatJSON() string {
	return FormatJSON(self)
}

type PortsList []*Ports

func (list *PortsList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *PortsList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *PortsList) FormatJSON() string {
	return FormatJSON(list)
}
