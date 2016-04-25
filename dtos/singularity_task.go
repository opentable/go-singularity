package dtos

import "io"

type SingularityTask struct {
	MesosTask   *TaskInfo               `json:"mesosTask"`
	Offer       *Offer                  `json:"offer"`
	RackId      string                  `json:"rackId"`
	TaskId      *SingularityTaskId      `json:"taskId"`
	TaskRequest *SingularityTaskRequest `json:"taskRequest"`
}

func (self *SingularityTask) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SingularityTask) FormatText() string {
	return FormatText(self)
}

func (self *SingularityTask) FormatJSON() string {
	return FormatJSON(self)
}

type SingularityTaskList []*SingularityTask

func (list *SingularityTaskList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *SingularityTaskList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityTaskList) FormatJSON() string {
	return FormatJSON(list)
}
