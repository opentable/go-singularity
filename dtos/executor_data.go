package dtos

import "io"

type ExecutorData struct {
	Cmd                            string                  `json:"cmd"`
	EmbeddedArtifacts              EmbeddedArtifactList    `json:"embeddedArtifacts"`
	ExternalArtifacts              ExternalArtifactList    `json:"externalArtifacts"`
	ExtraCmdLineArgs               StringList              `json:"extraCmdLineArgs"`
	LoggingExtraFields             map[string]string       `json:"loggingExtraFields"`
	LoggingS3Bucket                string                  `json:"loggingS3Bucket"`
	LoggingTag                     string                  `json:"loggingTag"`
	MaxOpenFiles                   int32                   `json:"maxOpenFiles"`
	MaxTaskThreads                 int32                   `json:"maxTaskThreads"`
	PreserveTaskSandboxAfterFinish bool                    `json:"preserveTaskSandboxAfterFinish"`
	RunningSentinel                string                  `json:"runningSentinel"`
	S3ArtifactSignatures           S3ArtifactSignatureList `json:"s3ArtifactSignatures"`
	S3Artifacts                    S3ArtifactList          `json:"s3Artifacts"`
	SigKillProcessesAfterMillis    int64                   `json:"sigKillProcessesAfterMillis"`
	SkipLogrotateAndCompress       bool                    `json:"skipLogrotateAndCompress"`
	SuccessfulExitCodes            []int32                 `json:"successfulExitCodes"`
	User                           string                  `json:"user"`
}

func (self *ExecutorData) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *ExecutorData) FormatText() string {
	return FormatText(self)
}

func (self *ExecutorData) FormatJSON() string {
	return FormatJSON(self)
}

type ExecutorDataList []*ExecutorData

func (list *ExecutorDataList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *ExecutorDataList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *ExecutorDataList) FormatJSON() string {
	return FormatJSON(list)
}
