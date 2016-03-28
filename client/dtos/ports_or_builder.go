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
