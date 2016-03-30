package dtos

import "io"

type LabelsOrBuilder struct {
	LabelsCount int32
	//	LabelsList *List[Label]
	//	LabelsOrBuilderList *List[? extends org.apache.mesos.Protos$LabelOrBuilder]
}

func (self *LabelsOrBuilder) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *LabelsOrBuilder) FormatText() string {
	return FormatText(self)
}

func (self *LabelsOrBuilder) FormatJSON() string {
	return FormatJSON(self)
}

type LabelsOrBuilderList []*LabelsOrBuilder

func (list LabelsOrBuilderList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list LabelsOrBuilderList) FormatText() string {
	text := []byte{}
	for _, dto := range list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list LabelsOrBuilderList) FormatJSON() string {
	return FormatJSON(list)
}
