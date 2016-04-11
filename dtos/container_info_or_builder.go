package dtos

import "io"

type ContainerInfoOrBuilder struct {
	Docker          *DockerInfo
	DockerOrBuilder *DockerInfoOrBuilder
	Hostname        string
	HostnameBytes   *ByteString
	//	Type *Type
	VolumesCount int32
	//	VolumesList *List[Volume]
	//	VolumesOrBuilderList *List[? extends org.apache.mesos.Protos$VolumeOrBuilder]
}

func (self *ContainerInfoOrBuilder) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *ContainerInfoOrBuilder) FormatText() string {
	return FormatText(self)
}

func (self *ContainerInfoOrBuilder) FormatJSON() string {
	return FormatJSON(self)
}

type ContainerInfoOrBuilderList []*ContainerInfoOrBuilder

func (list *ContainerInfoOrBuilderList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *ContainerInfoOrBuilderList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *ContainerInfoOrBuilderList) FormatJSON() string {
	return FormatJSON(list)
}
