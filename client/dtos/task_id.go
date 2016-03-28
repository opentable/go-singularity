package dtos

import "io"

type TaskID struct {
	//	AllFields *Map[FieldDescriptor,Object]
	DefaultInstanceForType    *TaskID
	DescriptorForType         *Descriptor
	InitializationErrorString string
	Initialized               bool
	//	ParserForType *com.google.protobuf.Parser&lt;org.apache.mesos.Protos$TaskID&gt;
	SerializedSize int32
	UnknownFields  *UnknownFieldSet
	Value          string
	ValueBytes     *ByteString
}

func (self *TaskID) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *TaskID) FormatText() string {
	return FormatText(self)
}

func (self *TaskID) FormatJSON() string {
	return FormatJSON(self)
}
