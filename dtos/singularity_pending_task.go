package dtos

import "io"

type SingularityPendingTask struct {
	CmdLineArgsList  StringList                `json:"cmdLineArgsList"`
	Message          string                    `json:"message"`
	PendingTaskId    *SingularityPendingTaskId `json:"pendingTaskId"`
	RunId            string                    `json:"runId"`
	SkipHealthchecks bool                      `json:"skipHealthchecks"`
	User             string                    `json:"user"`
}

func (self *SingularityPendingTask) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SingularityPendingTask) FormatText() string {
	return FormatText(self)
}

func (self *SingularityPendingTask) FormatJSON() string {
	return FormatJSON(self)
}

type SingularityPendingTaskList []*SingularityPendingTask

func (list *SingularityPendingTaskList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *SingularityPendingTaskList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityPendingTaskList) FormatJSON() string {
	return FormatJSON(list)
}
