package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type ExecutorData struct {
	present map[string]bool

	Cmd string `json:"cmd,omitempty"`

	EmbeddedArtifacts EmbeddedArtifactList `json:"embeddedArtifacts"`

	ExternalArtifacts ExternalArtifactList `json:"externalArtifacts"`

	ExtraCmdLineArgs swaggering.StringList `json:"extraCmdLineArgs"`

	LoggingExtraFields map[string]string `json:"loggingExtraFields"`

	LoggingS3Bucket string `json:"loggingS3Bucket,omitempty"`

	LoggingTag string `json:"loggingTag,omitempty"`

	MaxOpenFiles int32 `json:"maxOpenFiles"`

	MaxTaskThreads int32 `json:"maxTaskThreads"`

	PreserveTaskSandboxAfterFinish bool `json:"preserveTaskSandboxAfterFinish"`

	RunningSentinel string `json:"runningSentinel,omitempty"`

	S3ArtifactSignatures S3ArtifactSignatureList `json:"s3ArtifactSignatures"`

	S3Artifacts S3ArtifactList `json:"s3Artifacts"`

	SigKillProcessesAfterMillis int64 `json:"sigKillProcessesAfterMillis"`

	SkipLogrotateAndCompress bool `json:"skipLogrotateAndCompress"`

	SuccessfulExitCodes []int32 `json:"successfulExitCodes"`

	User string `json:"user,omitempty"`
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

func (self *ExecutorData) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *ExecutorData) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *ExecutorData) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *ExecutorData) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *ExecutorData) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on ExecutorData", name)

	case "cmd", "Cmd":
		v, ok := value.(string)
		if ok {
			self.Cmd = v
			self.present["cmd"] = true
			return nil
		} else {
			return fmt.Errorf("Field cmd/Cmd: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "embeddedArtifacts", "EmbeddedArtifacts":
		v, ok := value.(EmbeddedArtifactList)
		if ok {
			self.EmbeddedArtifacts = v
			self.present["embeddedArtifacts"] = true
			return nil
		} else {
			return fmt.Errorf("Field embeddedArtifacts/EmbeddedArtifacts: value %v(%T) couldn't be cast to type EmbeddedArtifactList", value, value)
		}

	case "externalArtifacts", "ExternalArtifacts":
		v, ok := value.(ExternalArtifactList)
		if ok {
			self.ExternalArtifacts = v
			self.present["externalArtifacts"] = true
			return nil
		} else {
			return fmt.Errorf("Field externalArtifacts/ExternalArtifacts: value %v(%T) couldn't be cast to type ExternalArtifactList", value, value)
		}

	case "extraCmdLineArgs", "ExtraCmdLineArgs":
		v, ok := value.(swaggering.StringList)
		if ok {
			self.ExtraCmdLineArgs = v
			self.present["extraCmdLineArgs"] = true
			return nil
		} else {
			return fmt.Errorf("Field extraCmdLineArgs/ExtraCmdLineArgs: value %v(%T) couldn't be cast to type StringList", value, value)
		}

	case "loggingExtraFields", "LoggingExtraFields":
		v, ok := value.(map[string]string)
		if ok {
			self.LoggingExtraFields = v
			self.present["loggingExtraFields"] = true
			return nil
		} else {
			return fmt.Errorf("Field loggingExtraFields/LoggingExtraFields: value %v(%T) couldn't be cast to type map[string]string", value, value)
		}

	case "loggingS3Bucket", "LoggingS3Bucket":
		v, ok := value.(string)
		if ok {
			self.LoggingS3Bucket = v
			self.present["loggingS3Bucket"] = true
			return nil
		} else {
			return fmt.Errorf("Field loggingS3Bucket/LoggingS3Bucket: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "loggingTag", "LoggingTag":
		v, ok := value.(string)
		if ok {
			self.LoggingTag = v
			self.present["loggingTag"] = true
			return nil
		} else {
			return fmt.Errorf("Field loggingTag/LoggingTag: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "maxOpenFiles", "MaxOpenFiles":
		v, ok := value.(int32)
		if ok {
			self.MaxOpenFiles = v
			self.present["maxOpenFiles"] = true
			return nil
		} else {
			return fmt.Errorf("Field maxOpenFiles/MaxOpenFiles: value %v(%T) couldn't be cast to type int32", value, value)
		}

	case "maxTaskThreads", "MaxTaskThreads":
		v, ok := value.(int32)
		if ok {
			self.MaxTaskThreads = v
			self.present["maxTaskThreads"] = true
			return nil
		} else {
			return fmt.Errorf("Field maxTaskThreads/MaxTaskThreads: value %v(%T) couldn't be cast to type int32", value, value)
		}

	case "preserveTaskSandboxAfterFinish", "PreserveTaskSandboxAfterFinish":
		v, ok := value.(bool)
		if ok {
			self.PreserveTaskSandboxAfterFinish = v
			self.present["preserveTaskSandboxAfterFinish"] = true
			return nil
		} else {
			return fmt.Errorf("Field preserveTaskSandboxAfterFinish/PreserveTaskSandboxAfterFinish: value %v(%T) couldn't be cast to type bool", value, value)
		}

	case "runningSentinel", "RunningSentinel":
		v, ok := value.(string)
		if ok {
			self.RunningSentinel = v
			self.present["runningSentinel"] = true
			return nil
		} else {
			return fmt.Errorf("Field runningSentinel/RunningSentinel: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "s3ArtifactSignatures", "S3ArtifactSignatures":
		v, ok := value.(S3ArtifactSignatureList)
		if ok {
			self.S3ArtifactSignatures = v
			self.present["s3ArtifactSignatures"] = true
			return nil
		} else {
			return fmt.Errorf("Field s3ArtifactSignatures/S3ArtifactSignatures: value %v(%T) couldn't be cast to type S3ArtifactSignatureList", value, value)
		}

	case "s3Artifacts", "S3Artifacts":
		v, ok := value.(S3ArtifactList)
		if ok {
			self.S3Artifacts = v
			self.present["s3Artifacts"] = true
			return nil
		} else {
			return fmt.Errorf("Field s3Artifacts/S3Artifacts: value %v(%T) couldn't be cast to type S3ArtifactList", value, value)
		}

	case "sigKillProcessesAfterMillis", "SigKillProcessesAfterMillis":
		v, ok := value.(int64)
		if ok {
			self.SigKillProcessesAfterMillis = v
			self.present["sigKillProcessesAfterMillis"] = true
			return nil
		} else {
			return fmt.Errorf("Field sigKillProcessesAfterMillis/SigKillProcessesAfterMillis: value %v(%T) couldn't be cast to type int64", value, value)
		}

	case "skipLogrotateAndCompress", "SkipLogrotateAndCompress":
		v, ok := value.(bool)
		if ok {
			self.SkipLogrotateAndCompress = v
			self.present["skipLogrotateAndCompress"] = true
			return nil
		} else {
			return fmt.Errorf("Field skipLogrotateAndCompress/SkipLogrotateAndCompress: value %v(%T) couldn't be cast to type bool", value, value)
		}

	case "successfulExitCodes", "SuccessfulExitCodes":
		v, ok := value.([]int32)
		if ok {
			self.SuccessfulExitCodes = v
			self.present["successfulExitCodes"] = true
			return nil
		} else {
			return fmt.Errorf("Field successfulExitCodes/SuccessfulExitCodes: value %v(%T) couldn't be cast to type []int32", value, value)
		}

	case "user", "User":
		v, ok := value.(string)
		if ok {
			self.User = v
			self.present["user"] = true
			return nil
		} else {
			return fmt.Errorf("Field user/User: value %v(%T) couldn't be cast to type string", value, value)
		}

	}
}

func (self *ExecutorData) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on ExecutorData", name)

	case "cmd", "Cmd":
		if self.present != nil {
			if _, ok := self.present["cmd"]; ok {
				return self.Cmd, nil
			}
		}
		return nil, fmt.Errorf("Field Cmd no set on Cmd %+v", self)

	case "embeddedArtifacts", "EmbeddedArtifacts":
		if self.present != nil {
			if _, ok := self.present["embeddedArtifacts"]; ok {
				return self.EmbeddedArtifacts, nil
			}
		}
		return nil, fmt.Errorf("Field EmbeddedArtifacts no set on EmbeddedArtifacts %+v", self)

	case "externalArtifacts", "ExternalArtifacts":
		if self.present != nil {
			if _, ok := self.present["externalArtifacts"]; ok {
				return self.ExternalArtifacts, nil
			}
		}
		return nil, fmt.Errorf("Field ExternalArtifacts no set on ExternalArtifacts %+v", self)

	case "extraCmdLineArgs", "ExtraCmdLineArgs":
		if self.present != nil {
			if _, ok := self.present["extraCmdLineArgs"]; ok {
				return self.ExtraCmdLineArgs, nil
			}
		}
		return nil, fmt.Errorf("Field ExtraCmdLineArgs no set on ExtraCmdLineArgs %+v", self)

	case "loggingExtraFields", "LoggingExtraFields":
		if self.present != nil {
			if _, ok := self.present["loggingExtraFields"]; ok {
				return self.LoggingExtraFields, nil
			}
		}
		return nil, fmt.Errorf("Field LoggingExtraFields no set on LoggingExtraFields %+v", self)

	case "loggingS3Bucket", "LoggingS3Bucket":
		if self.present != nil {
			if _, ok := self.present["loggingS3Bucket"]; ok {
				return self.LoggingS3Bucket, nil
			}
		}
		return nil, fmt.Errorf("Field LoggingS3Bucket no set on LoggingS3Bucket %+v", self)

	case "loggingTag", "LoggingTag":
		if self.present != nil {
			if _, ok := self.present["loggingTag"]; ok {
				return self.LoggingTag, nil
			}
		}
		return nil, fmt.Errorf("Field LoggingTag no set on LoggingTag %+v", self)

	case "maxOpenFiles", "MaxOpenFiles":
		if self.present != nil {
			if _, ok := self.present["maxOpenFiles"]; ok {
				return self.MaxOpenFiles, nil
			}
		}
		return nil, fmt.Errorf("Field MaxOpenFiles no set on MaxOpenFiles %+v", self)

	case "maxTaskThreads", "MaxTaskThreads":
		if self.present != nil {
			if _, ok := self.present["maxTaskThreads"]; ok {
				return self.MaxTaskThreads, nil
			}
		}
		return nil, fmt.Errorf("Field MaxTaskThreads no set on MaxTaskThreads %+v", self)

	case "preserveTaskSandboxAfterFinish", "PreserveTaskSandboxAfterFinish":
		if self.present != nil {
			if _, ok := self.present["preserveTaskSandboxAfterFinish"]; ok {
				return self.PreserveTaskSandboxAfterFinish, nil
			}
		}
		return nil, fmt.Errorf("Field PreserveTaskSandboxAfterFinish no set on PreserveTaskSandboxAfterFinish %+v", self)

	case "runningSentinel", "RunningSentinel":
		if self.present != nil {
			if _, ok := self.present["runningSentinel"]; ok {
				return self.RunningSentinel, nil
			}
		}
		return nil, fmt.Errorf("Field RunningSentinel no set on RunningSentinel %+v", self)

	case "s3ArtifactSignatures", "S3ArtifactSignatures":
		if self.present != nil {
			if _, ok := self.present["s3ArtifactSignatures"]; ok {
				return self.S3ArtifactSignatures, nil
			}
		}
		return nil, fmt.Errorf("Field S3ArtifactSignatures no set on S3ArtifactSignatures %+v", self)

	case "s3Artifacts", "S3Artifacts":
		if self.present != nil {
			if _, ok := self.present["s3Artifacts"]; ok {
				return self.S3Artifacts, nil
			}
		}
		return nil, fmt.Errorf("Field S3Artifacts no set on S3Artifacts %+v", self)

	case "sigKillProcessesAfterMillis", "SigKillProcessesAfterMillis":
		if self.present != nil {
			if _, ok := self.present["sigKillProcessesAfterMillis"]; ok {
				return self.SigKillProcessesAfterMillis, nil
			}
		}
		return nil, fmt.Errorf("Field SigKillProcessesAfterMillis no set on SigKillProcessesAfterMillis %+v", self)

	case "skipLogrotateAndCompress", "SkipLogrotateAndCompress":
		if self.present != nil {
			if _, ok := self.present["skipLogrotateAndCompress"]; ok {
				return self.SkipLogrotateAndCompress, nil
			}
		}
		return nil, fmt.Errorf("Field SkipLogrotateAndCompress no set on SkipLogrotateAndCompress %+v", self)

	case "successfulExitCodes", "SuccessfulExitCodes":
		if self.present != nil {
			if _, ok := self.present["successfulExitCodes"]; ok {
				return self.SuccessfulExitCodes, nil
			}
		}
		return nil, fmt.Errorf("Field SuccessfulExitCodes no set on SuccessfulExitCodes %+v", self)

	case "user", "User":
		if self.present != nil {
			if _, ok := self.present["user"]; ok {
				return self.User, nil
			}
		}
		return nil, fmt.Errorf("Field User no set on User %+v", self)

	}
}

func (self *ExecutorData) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on ExecutorData", name)

	case "cmd", "Cmd":
		self.present["cmd"] = false

	case "embeddedArtifacts", "EmbeddedArtifacts":
		self.present["embeddedArtifacts"] = false

	case "externalArtifacts", "ExternalArtifacts":
		self.present["externalArtifacts"] = false

	case "extraCmdLineArgs", "ExtraCmdLineArgs":
		self.present["extraCmdLineArgs"] = false

	case "loggingExtraFields", "LoggingExtraFields":
		self.present["loggingExtraFields"] = false

	case "loggingS3Bucket", "LoggingS3Bucket":
		self.present["loggingS3Bucket"] = false

	case "loggingTag", "LoggingTag":
		self.present["loggingTag"] = false

	case "maxOpenFiles", "MaxOpenFiles":
		self.present["maxOpenFiles"] = false

	case "maxTaskThreads", "MaxTaskThreads":
		self.present["maxTaskThreads"] = false

	case "preserveTaskSandboxAfterFinish", "PreserveTaskSandboxAfterFinish":
		self.present["preserveTaskSandboxAfterFinish"] = false

	case "runningSentinel", "RunningSentinel":
		self.present["runningSentinel"] = false

	case "s3ArtifactSignatures", "S3ArtifactSignatures":
		self.present["s3ArtifactSignatures"] = false

	case "s3Artifacts", "S3Artifacts":
		self.present["s3Artifacts"] = false

	case "sigKillProcessesAfterMillis", "SigKillProcessesAfterMillis":
		self.present["sigKillProcessesAfterMillis"] = false

	case "skipLogrotateAndCompress", "SkipLogrotateAndCompress":
		self.present["skipLogrotateAndCompress"] = false

	case "successfulExitCodes", "SuccessfulExitCodes":
		self.present["successfulExitCodes"] = false

	case "user", "User":
		self.present["user"] = false

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
