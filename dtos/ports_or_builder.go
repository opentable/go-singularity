package dtos

import "io"

type PortsOrBuilder struct {
	PortsCount int32
	//	PortsList *List[Port]
	//	PortsOrBuilderList *List[? extends org.apache.mesos.Protos$PortOrBuilder]
}

func (self *PortsOrBuilder) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *PortsOrBuilder) FormatText() string {
	return FormatText(self)
}

func (self *PortsOrBuilder) FormatJSON() string {
	return FormatJSON(self)
}

type PortsOrBuilderList []*PortsOrBuilder

func (list *PortsOrBuilderList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *PortsOrBuilderList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *PortsOrBuilderList) FormatJSON() string {
	return FormatJSON(list)
}
