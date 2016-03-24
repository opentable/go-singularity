package dtos

import "io"

type ExecutorData struct {
	Cmd string
	EmbeddedArtifacts []
	ExternalArtifacts []
	ExtraCmdLineArgs []string
//	LoggingExtraFields Map[string,string]
	LoggingS3Bucket string
	LoggingTag string
	MaxOpenFiles int32
	MaxTaskThreads int32
	PreserveTaskSandboxAfterFinish bool
	RunningSentinel string
	S3ArtifactSignatures []
	S3Artifacts []
	SigKillProcessesAfterMillis int64
	SkipLogrotateAndCompress bool
	SuccessfulExitCodes []int32
	User string
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
