package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityS3Log struct {
	GetUrl       *string `json:"getUrl,omitempty"`
	Key          *string `json:"key,omitempty"`
	LastModified *int64  `json:"lastModified,omitempty"`
	Size         *int64  `json:"size,omitempty"`
	DownloadUrl  *string `json:"downloadUrl,omitempty"`
}

func (self *SingularityS3Log) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *SingularityS3Log) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityS3Log); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityS3Log cannot copy the values from %#v", other)
}

func (self *SingularityS3Log) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityS3Log) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityS3Log) SetField(name string, value interface{}) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityS3Log", name)

	case "getUrl", "GetUrl":
		v, ok := value.(string)
		if ok {
			self.GetUrl = &v
			return nil
		}
		return fmt.Errorf("Field getUrl/GetUrl: value %v(%T) couldn't be cast to type string", value, value)

	case "key", "Key":
		v, ok := value.(string)
		if ok {
			self.Key = &v
			return nil
		}
		return fmt.Errorf("Field key/Key: value %v(%T) couldn't be cast to type string", value, value)

	case "lastModified", "LastModified":
		v, ok := value.(int64)
		if ok {
			self.LastModified = &v
			return nil
		}
		return fmt.Errorf("Field lastModified/LastModified: value %v(%T) couldn't be cast to type int64", value, value)

	case "size", "Size":
		v, ok := value.(int64)
		if ok {
			self.Size = &v
			return nil
		}
		return fmt.Errorf("Field size/Size: value %v(%T) couldn't be cast to type int64", value, value)

	case "downloadUrl", "DownloadUrl":
		v, ok := value.(string)
		if ok {
			self.DownloadUrl = &v
			return nil
		}
		return fmt.Errorf("Field downloadUrl/DownloadUrl: value %v(%T) couldn't be cast to type string", value, value)

	}
}

func (self *SingularityS3Log) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityS3Log", name)

	case "getUrl", "GetUrl":
		return *self.GetUrl, nil
		return nil, fmt.Errorf("Field GetUrl no set on GetUrl %+v", self)

	case "key", "Key":
		return *self.Key, nil
		return nil, fmt.Errorf("Field Key no set on Key %+v", self)

	case "lastModified", "LastModified":
		return *self.LastModified, nil
		return nil, fmt.Errorf("Field LastModified no set on LastModified %+v", self)

	case "size", "Size":
		return *self.Size, nil
		return nil, fmt.Errorf("Field Size no set on Size %+v", self)

	case "downloadUrl", "DownloadUrl":
		return *self.DownloadUrl, nil
		return nil, fmt.Errorf("Field DownloadUrl no set on DownloadUrl %+v", self)

	}
}

func (self *SingularityS3Log) ClearField(name string) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityS3Log", name)

	case "getUrl", "GetUrl":
		self.GetUrl = nil

	case "key", "Key":
		self.Key = nil

	case "lastModified", "LastModified":
		self.LastModified = nil

	case "size", "Size":
		self.Size = nil

	case "downloadUrl", "DownloadUrl":
		self.DownloadUrl = nil

	}

	return nil
}

func (self *SingularityS3Log) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type SingularityS3LogList []*SingularityS3Log

func (self *SingularityS3LogList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityS3LogList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityS3LogList cannot copy the values from %#v", other)
}

func (list *SingularityS3LogList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *SingularityS3LogList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityS3LogList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
