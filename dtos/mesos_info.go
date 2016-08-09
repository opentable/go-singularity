package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type MesosInfo struct {
	present map[string]bool

	// AllFields *Map[FieldDescriptor,Object] `json:"allFields"`

	DefaultInstanceForType *MesosInfo `json:"defaultInstanceForType"`

	DescriptorForType *Descriptor `json:"descriptorForType"`

	Image *Image `json:"image"`

	ImageOrBuilder *ImageOrBuilder `json:"imageOrBuilder"`

	InitializationErrorString string `json:"initializationErrorString,omitempty"`

	Initialized bool `json:"initialized"`

	// ParserForType *com.google.protobuf.Parser<org.apache.mesos.Protos$ContainerInfo$MesosInfo> `json:"parserForType"`

	SerializedSize int32 `json:"serializedSize"`

	UnknownFields *UnknownFieldSet `json:"unknownFields"`
}

func (self *MesosInfo) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *MesosInfo) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*MesosInfo); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A MesosInfo cannot absorb the values from %v", other)
}

func (self *MesosInfo) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *MesosInfo) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *MesosInfo) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *MesosInfo) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *MesosInfo) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on MesosInfo", name)

	case "defaultInstanceForType", "DefaultInstanceForType":
		v, ok := value.(*MesosInfo)
		if ok {
			self.DefaultInstanceForType = v
			self.present["defaultInstanceForType"] = true
			return nil
		} else {
			return fmt.Errorf("Field defaultInstanceForType/DefaultInstanceForType: value %v(%T) couldn't be cast to type *MesosInfo", value, value)
		}

	case "descriptorForType", "DescriptorForType":
		v, ok := value.(*Descriptor)
		if ok {
			self.DescriptorForType = v
			self.present["descriptorForType"] = true
			return nil
		} else {
			return fmt.Errorf("Field descriptorForType/DescriptorForType: value %v(%T) couldn't be cast to type *Descriptor", value, value)
		}

	case "image", "Image":
		v, ok := value.(*Image)
		if ok {
			self.Image = v
			self.present["image"] = true
			return nil
		} else {
			return fmt.Errorf("Field image/Image: value %v(%T) couldn't be cast to type *Image", value, value)
		}

	case "imageOrBuilder", "ImageOrBuilder":
		v, ok := value.(*ImageOrBuilder)
		if ok {
			self.ImageOrBuilder = v
			self.present["imageOrBuilder"] = true
			return nil
		} else {
			return fmt.Errorf("Field imageOrBuilder/ImageOrBuilder: value %v(%T) couldn't be cast to type *ImageOrBuilder", value, value)
		}

	case "initializationErrorString", "InitializationErrorString":
		v, ok := value.(string)
		if ok {
			self.InitializationErrorString = v
			self.present["initializationErrorString"] = true
			return nil
		} else {
			return fmt.Errorf("Field initializationErrorString/InitializationErrorString: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "initialized", "Initialized":
		v, ok := value.(bool)
		if ok {
			self.Initialized = v
			self.present["initialized"] = true
			return nil
		} else {
			return fmt.Errorf("Field initialized/Initialized: value %v(%T) couldn't be cast to type bool", value, value)
		}

	case "serializedSize", "SerializedSize":
		v, ok := value.(int32)
		if ok {
			self.SerializedSize = v
			self.present["serializedSize"] = true
			return nil
		} else {
			return fmt.Errorf("Field serializedSize/SerializedSize: value %v(%T) couldn't be cast to type int32", value, value)
		}

	case "unknownFields", "UnknownFields":
		v, ok := value.(*UnknownFieldSet)
		if ok {
			self.UnknownFields = v
			self.present["unknownFields"] = true
			return nil
		} else {
			return fmt.Errorf("Field unknownFields/UnknownFields: value %v(%T) couldn't be cast to type *UnknownFieldSet", value, value)
		}

	}
}

func (self *MesosInfo) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on MesosInfo", name)

	case "defaultInstanceForType", "DefaultInstanceForType":
		if self.present != nil {
			if _, ok := self.present["defaultInstanceForType"]; ok {
				return self.DefaultInstanceForType, nil
			}
		}
		return nil, fmt.Errorf("Field DefaultInstanceForType no set on DefaultInstanceForType %+v", self)

	case "descriptorForType", "DescriptorForType":
		if self.present != nil {
			if _, ok := self.present["descriptorForType"]; ok {
				return self.DescriptorForType, nil
			}
		}
		return nil, fmt.Errorf("Field DescriptorForType no set on DescriptorForType %+v", self)

	case "image", "Image":
		if self.present != nil {
			if _, ok := self.present["image"]; ok {
				return self.Image, nil
			}
		}
		return nil, fmt.Errorf("Field Image no set on Image %+v", self)

	case "imageOrBuilder", "ImageOrBuilder":
		if self.present != nil {
			if _, ok := self.present["imageOrBuilder"]; ok {
				return self.ImageOrBuilder, nil
			}
		}
		return nil, fmt.Errorf("Field ImageOrBuilder no set on ImageOrBuilder %+v", self)

	case "initializationErrorString", "InitializationErrorString":
		if self.present != nil {
			if _, ok := self.present["initializationErrorString"]; ok {
				return self.InitializationErrorString, nil
			}
		}
		return nil, fmt.Errorf("Field InitializationErrorString no set on InitializationErrorString %+v", self)

	case "initialized", "Initialized":
		if self.present != nil {
			if _, ok := self.present["initialized"]; ok {
				return self.Initialized, nil
			}
		}
		return nil, fmt.Errorf("Field Initialized no set on Initialized %+v", self)

	case "serializedSize", "SerializedSize":
		if self.present != nil {
			if _, ok := self.present["serializedSize"]; ok {
				return self.SerializedSize, nil
			}
		}
		return nil, fmt.Errorf("Field SerializedSize no set on SerializedSize %+v", self)

	case "unknownFields", "UnknownFields":
		if self.present != nil {
			if _, ok := self.present["unknownFields"]; ok {
				return self.UnknownFields, nil
			}
		}
		return nil, fmt.Errorf("Field UnknownFields no set on UnknownFields %+v", self)

	}
}

func (self *MesosInfo) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on MesosInfo", name)

	case "defaultInstanceForType", "DefaultInstanceForType":
		self.present["defaultInstanceForType"] = false

	case "descriptorForType", "DescriptorForType":
		self.present["descriptorForType"] = false

	case "image", "Image":
		self.present["image"] = false

	case "imageOrBuilder", "ImageOrBuilder":
		self.present["imageOrBuilder"] = false

	case "initializationErrorString", "InitializationErrorString":
		self.present["initializationErrorString"] = false

	case "initialized", "Initialized":
		self.present["initialized"] = false

	case "serializedSize", "SerializedSize":
		self.present["serializedSize"] = false

	case "unknownFields", "UnknownFields":
		self.present["unknownFields"] = false

	}

	return nil
}

func (self *MesosInfo) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type MesosInfoList []*MesosInfo

func (self *MesosInfoList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*MesosInfoList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A MesosInfo cannot absorb the values from %v", other)
}

func (list *MesosInfoList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *MesosInfoList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *MesosInfoList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
