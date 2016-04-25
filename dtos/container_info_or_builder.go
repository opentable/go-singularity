package dtos

import "io"

type ContainerInfoOrBuilderType string

const (
	ContainerInfoOrBuilderTypeDOCKER ContainerInfoOrBuilderType = "DOCKER"
	ContainerInfoOrBuilderTypeMESOS  ContainerInfoOrBuilderType = "MESOS"
)

type ContainerInfoOrBuilder struct {
	Docker          *DockerInfo                `json:"docker"`
	DockerOrBuilder *DockerInfoOrBuilder       `json:"dockerOrBuilder"`
	Hostname        string                     `json:"hostname"`
	HostnameBytes   *ByteString                `json:"hostnameBytes"`
	Type            ContainerInfoOrBuilderType `json:"type"`
	VolumesCount    int32                      `json:"volumesCount"`
	//	VolumesList *List[Volume] `json:"volumesList"`
	//	VolumesOrBuilderList *List[? extends org.apache.mesos.Protos$VolumeOrBuilder] `json:"volumesOrBuilderList"`

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
