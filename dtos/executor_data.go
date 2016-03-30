package dtos

import "io"

type ExecutorData struct {
	Cmd               string
	EmbeddedArtifacts EmbeddedArtifactList
	ExternalArtifacts ExternalArtifactList
	ExtraCmdLineArgs  StringList
	//	LoggingExtraFields *Map[string,string]
	LoggingS3Bucket                string
	LoggingTag                     string
	MaxOpenFiles                   int32
	MaxTaskThreads                 int32
	PreserveTaskSandboxAfterFinish bool
	RunningSentinel                string
	S3ArtifactSignatures           S3ArtifactSignatureList
	S3Artifacts                    S3ArtifactList
	SigKillProcessesAfterMillis    int64
	SkipLogrotateAndCompress       bool
	SuccessfulExitCodes            []int32
	User                           string
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

func (list ExecutorDataList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list ExecutorDataList) FormatText() string {
	text := []byte{}
	for _, dto := range list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list ExecutorDataList) FormatJSON() string {
	return FormatJSON(list)
}
