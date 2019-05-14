package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type ExecutorDataSingularityExecutorLogrotateFrequency string

const (
	ExecutorDataSingularityExecutorLogrotateFrequencyHOURLY  ExecutorDataSingularityExecutorLogrotateFrequency = "HOURLY"
	ExecutorDataSingularityExecutorLogrotateFrequencyDAILY   ExecutorDataSingularityExecutorLogrotateFrequency = "DAILY"
	ExecutorDataSingularityExecutorLogrotateFrequencyWEEKLY  ExecutorDataSingularityExecutorLogrotateFrequency = "WEEKLY"
	ExecutorDataSingularityExecutorLogrotateFrequencyMONTHLY ExecutorDataSingularityExecutorLogrotateFrequency = "MONTHLY"
)

type ExecutorData struct {
	ExternalArtifacts              *ExternalArtifactList                              `json:"externalArtifacts,omitempty"`
	ExtraCmdLineArgs               *swaggering.StringList                             `json:"extraCmdLineArgs,omitempty"`
	MaxTaskThreads                 *int32                                             `json:"maxTaskThreads,omitempty"`
	PreserveTaskSandboxAfterFinish *bool                                              `json:"preserveTaskSandboxAfterFinish,omitempty"`
	User                           *string                                            `json:"user,omitempty"`
	SigKillProcessesAfterMillis    *int64                                             `json:"sigKillProcessesAfterMillis,omitempty"`
	S3ArtifactSignatures           *S3ArtifactSignatureList                           `json:"s3ArtifactSignatures,omitempty"`
	EmbeddedArtifacts              *EmbeddedArtifactList                              `json:"embeddedArtifacts,omitempty"`
	S3Artifacts                    *S3ArtifactList                                    `json:"s3Artifacts,omitempty"`
	SuccessfulExitCodes            *[]int32                                           `json:"successfulExitCodes,omitempty"`
	RunningSentinel                *string                                            `json:"runningSentinel,omitempty"`
	MaxOpenFiles                   *int32                                             `json:"maxOpenFiles,omitempty"`
	Cmd                            *string                                            `json:"cmd,omitempty"`
	LoggingTag                     *string                                            `json:"loggingTag,omitempty"`
	LoggingExtraFields             *map[string]string                                 `json:"loggingExtraFields,omitempty"`
	SkipLogrotateAndCompress       *bool                                              `json:"skipLogrotateAndCompress,omitempty"`
	LogrotateFrequency             *ExecutorDataSingularityExecutorLogrotateFrequency `json:"logrotateFrequency,omitempty"`
}

func (self *ExecutorData) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *ExecutorData) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*ExecutorData); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A ExecutorData cannot copy the values from %#v", other)
}

func (self *ExecutorData) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *ExecutorData) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *ExecutorData) SetField(name string, value interface{}) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on ExecutorData", name)

	case "externalArtifacts", "ExternalArtifacts":
		v, ok := value.(ExternalArtifactList)
		if ok {
			self.ExternalArtifacts = &v
			return nil
		}
		return fmt.Errorf("Field externalArtifacts/ExternalArtifacts: value %v(%T) couldn't be cast to type ExternalArtifactList", value, value)

	case "extraCmdLineArgs", "ExtraCmdLineArgs":
		v, ok := value.(swaggering.StringList)
		if ok {
			self.ExtraCmdLineArgs = &v
			return nil
		}
		return fmt.Errorf("Field extraCmdLineArgs/ExtraCmdLineArgs: value %v(%T) couldn't be cast to type swaggering.StringList", value, value)

	case "maxTaskThreads", "MaxTaskThreads":
		v, ok := value.(int32)
		if ok {
			self.MaxTaskThreads = &v
			return nil
		}
		return fmt.Errorf("Field maxTaskThreads/MaxTaskThreads: value %v(%T) couldn't be cast to type int32", value, value)

	case "preserveTaskSandboxAfterFinish", "PreserveTaskSandboxAfterFinish":
		v, ok := value.(bool)
		if ok {
			self.PreserveTaskSandboxAfterFinish = &v
			return nil
		}
		return fmt.Errorf("Field preserveTaskSandboxAfterFinish/PreserveTaskSandboxAfterFinish: value %v(%T) couldn't be cast to type bool", value, value)

	case "user", "User":
		v, ok := value.(string)
		if ok {
			self.User = &v
			return nil
		}
		return fmt.Errorf("Field user/User: value %v(%T) couldn't be cast to type string", value, value)

	case "sigKillProcessesAfterMillis", "SigKillProcessesAfterMillis":
		v, ok := value.(int64)
		if ok {
			self.SigKillProcessesAfterMillis = &v
			return nil
		}
		return fmt.Errorf("Field sigKillProcessesAfterMillis/SigKillProcessesAfterMillis: value %v(%T) couldn't be cast to type int64", value, value)

	case "s3ArtifactSignatures", "S3ArtifactSignatures":
		v, ok := value.(S3ArtifactSignatureList)
		if ok {
			self.S3ArtifactSignatures = &v
			return nil
		}
		return fmt.Errorf("Field s3ArtifactSignatures/S3ArtifactSignatures: value %v(%T) couldn't be cast to type S3ArtifactSignatureList", value, value)

	case "embeddedArtifacts", "EmbeddedArtifacts":
		v, ok := value.(EmbeddedArtifactList)
		if ok {
			self.EmbeddedArtifacts = &v
			return nil
		}
		return fmt.Errorf("Field embeddedArtifacts/EmbeddedArtifacts: value %v(%T) couldn't be cast to type EmbeddedArtifactList", value, value)

	case "s3Artifacts", "S3Artifacts":
		v, ok := value.(S3ArtifactList)
		if ok {
			self.S3Artifacts = &v
			return nil
		}
		return fmt.Errorf("Field s3Artifacts/S3Artifacts: value %v(%T) couldn't be cast to type S3ArtifactList", value, value)

	case "successfulExitCodes", "SuccessfulExitCodes":
		v, ok := value.([]int32)
		if ok {
			self.SuccessfulExitCodes = &v
			return nil
		}
		return fmt.Errorf("Field successfulExitCodes/SuccessfulExitCodes: value %v(%T) couldn't be cast to type []int32", value, value)

	case "runningSentinel", "RunningSentinel":
		v, ok := value.(string)
		if ok {
			self.RunningSentinel = &v
			return nil
		}
		return fmt.Errorf("Field runningSentinel/RunningSentinel: value %v(%T) couldn't be cast to type string", value, value)

	case "maxOpenFiles", "MaxOpenFiles":
		v, ok := value.(int32)
		if ok {
			self.MaxOpenFiles = &v
			return nil
		}
		return fmt.Errorf("Field maxOpenFiles/MaxOpenFiles: value %v(%T) couldn't be cast to type int32", value, value)

	case "cmd", "Cmd":
		v, ok := value.(string)
		if ok {
			self.Cmd = &v
			return nil
		}
		return fmt.Errorf("Field cmd/Cmd: value %v(%T) couldn't be cast to type string", value, value)

	case "loggingTag", "LoggingTag":
		v, ok := value.(string)
		if ok {
			self.LoggingTag = &v
			return nil
		}
		return fmt.Errorf("Field loggingTag/LoggingTag: value %v(%T) couldn't be cast to type string", value, value)

	case "loggingExtraFields", "LoggingExtraFields":
		v, ok := value.(map[string]string)
		if ok {
			self.LoggingExtraFields = &v
			return nil
		}
		return fmt.Errorf("Field loggingExtraFields/LoggingExtraFields: value %v(%T) couldn't be cast to type map[string]string", value, value)

	case "skipLogrotateAndCompress", "SkipLogrotateAndCompress":
		v, ok := value.(bool)
		if ok {
			self.SkipLogrotateAndCompress = &v
			return nil
		}
		return fmt.Errorf("Field skipLogrotateAndCompress/SkipLogrotateAndCompress: value %v(%T) couldn't be cast to type bool", value, value)

	case "logrotateFrequency", "LogrotateFrequency":
		v, ok := value.(ExecutorDataSingularityExecutorLogrotateFrequency)
		if ok {
			self.LogrotateFrequency = &v
			return nil
		}
		return fmt.Errorf("Field logrotateFrequency/LogrotateFrequency: value %v(%T) couldn't be cast to type ExecutorDataSingularityExecutorLogrotateFrequency", value, value)

	}
}

func (self *ExecutorData) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on ExecutorData", name)

	case "externalArtifacts", "ExternalArtifacts":
		return *self.ExternalArtifacts, nil
		return nil, fmt.Errorf("Field ExternalArtifacts no set on ExternalArtifacts %+v", self)

	case "extraCmdLineArgs", "ExtraCmdLineArgs":
		return *self.ExtraCmdLineArgs, nil
		return nil, fmt.Errorf("Field ExtraCmdLineArgs no set on ExtraCmdLineArgs %+v", self)

	case "maxTaskThreads", "MaxTaskThreads":
		return *self.MaxTaskThreads, nil
		return nil, fmt.Errorf("Field MaxTaskThreads no set on MaxTaskThreads %+v", self)

	case "preserveTaskSandboxAfterFinish", "PreserveTaskSandboxAfterFinish":
		return *self.PreserveTaskSandboxAfterFinish, nil
		return nil, fmt.Errorf("Field PreserveTaskSandboxAfterFinish no set on PreserveTaskSandboxAfterFinish %+v", self)

	case "user", "User":
		return *self.User, nil
		return nil, fmt.Errorf("Field User no set on User %+v", self)

	case "sigKillProcessesAfterMillis", "SigKillProcessesAfterMillis":
		return *self.SigKillProcessesAfterMillis, nil
		return nil, fmt.Errorf("Field SigKillProcessesAfterMillis no set on SigKillProcessesAfterMillis %+v", self)

	case "s3ArtifactSignatures", "S3ArtifactSignatures":
		return *self.S3ArtifactSignatures, nil
		return nil, fmt.Errorf("Field S3ArtifactSignatures no set on S3ArtifactSignatures %+v", self)

	case "embeddedArtifacts", "EmbeddedArtifacts":
		return *self.EmbeddedArtifacts, nil
		return nil, fmt.Errorf("Field EmbeddedArtifacts no set on EmbeddedArtifacts %+v", self)

	case "s3Artifacts", "S3Artifacts":
		return *self.S3Artifacts, nil
		return nil, fmt.Errorf("Field S3Artifacts no set on S3Artifacts %+v", self)

	case "successfulExitCodes", "SuccessfulExitCodes":
		return *self.SuccessfulExitCodes, nil
		return nil, fmt.Errorf("Field SuccessfulExitCodes no set on SuccessfulExitCodes %+v", self)

	case "runningSentinel", "RunningSentinel":
		return *self.RunningSentinel, nil
		return nil, fmt.Errorf("Field RunningSentinel no set on RunningSentinel %+v", self)

	case "maxOpenFiles", "MaxOpenFiles":
		return *self.MaxOpenFiles, nil
		return nil, fmt.Errorf("Field MaxOpenFiles no set on MaxOpenFiles %+v", self)

	case "cmd", "Cmd":
		return *self.Cmd, nil
		return nil, fmt.Errorf("Field Cmd no set on Cmd %+v", self)

	case "loggingTag", "LoggingTag":
		return *self.LoggingTag, nil
		return nil, fmt.Errorf("Field LoggingTag no set on LoggingTag %+v", self)

	case "loggingExtraFields", "LoggingExtraFields":
		return *self.LoggingExtraFields, nil
		return nil, fmt.Errorf("Field LoggingExtraFields no set on LoggingExtraFields %+v", self)

	case "skipLogrotateAndCompress", "SkipLogrotateAndCompress":
		return *self.SkipLogrotateAndCompress, nil
		return nil, fmt.Errorf("Field SkipLogrotateAndCompress no set on SkipLogrotateAndCompress %+v", self)

	case "logrotateFrequency", "LogrotateFrequency":
		return *self.LogrotateFrequency, nil
		return nil, fmt.Errorf("Field LogrotateFrequency no set on LogrotateFrequency %+v", self)

	}
}

func (self *ExecutorData) ClearField(name string) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on ExecutorData", name)

	case "externalArtifacts", "ExternalArtifacts":
		self.ExternalArtifacts = nil

	case "extraCmdLineArgs", "ExtraCmdLineArgs":
		self.ExtraCmdLineArgs = nil

	case "maxTaskThreads", "MaxTaskThreads":
		self.MaxTaskThreads = nil

	case "preserveTaskSandboxAfterFinish", "PreserveTaskSandboxAfterFinish":
		self.PreserveTaskSandboxAfterFinish = nil

	case "user", "User":
		self.User = nil

	case "sigKillProcessesAfterMillis", "SigKillProcessesAfterMillis":
		self.SigKillProcessesAfterMillis = nil

	case "s3ArtifactSignatures", "S3ArtifactSignatures":
		self.S3ArtifactSignatures = nil

	case "embeddedArtifacts", "EmbeddedArtifacts":
		self.EmbeddedArtifacts = nil

	case "s3Artifacts", "S3Artifacts":
		self.S3Artifacts = nil

	case "successfulExitCodes", "SuccessfulExitCodes":
		self.SuccessfulExitCodes = nil

	case "runningSentinel", "RunningSentinel":
		self.RunningSentinel = nil

	case "maxOpenFiles", "MaxOpenFiles":
		self.MaxOpenFiles = nil

	case "cmd", "Cmd":
		self.Cmd = nil

	case "loggingTag", "LoggingTag":
		self.LoggingTag = nil

	case "loggingExtraFields", "LoggingExtraFields":
		self.LoggingExtraFields = nil

	case "skipLogrotateAndCompress", "SkipLogrotateAndCompress":
		self.SkipLogrotateAndCompress = nil

	case "logrotateFrequency", "LogrotateFrequency":
		self.LogrotateFrequency = nil

	}

	return nil
}

func (self *ExecutorData) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type ExecutorDataList []*ExecutorData

func (self *ExecutorDataList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*ExecutorDataList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A ExecutorDataList cannot copy the values from %#v", other)
}

func (list *ExecutorDataList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
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
	return swaggering.FormatJSON(list)
}
