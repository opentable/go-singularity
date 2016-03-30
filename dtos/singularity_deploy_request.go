package dtos

import "io"

type SingularityDeployRequest struct {
	Deploy                    *SingularityDeploy
	Message                   string
	UnpauseOnSuccessfulDeploy bool
}

func (self *SingularityDeployRequest) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SingularityDeployRequest) FormatText() string {
	return FormatText(self)
}

func (self *SingularityDeployRequest) FormatJSON() string {
	return FormatJSON(self)
}

type SingularityDeployRequestList []*SingularityDeployRequest

func (list SingularityDeployRequestList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list SingularityDeployRequestList) FormatText() string {
	text := []byte{}
	for _, dto := range list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list SingularityDeployRequestList) FormatJSON() string {
	return FormatJSON(list)
}
