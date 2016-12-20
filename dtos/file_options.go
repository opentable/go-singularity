package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type FileOptionsOptimizeMode string

const (
	FileOptionsOptimizeModeSPEED        FileOptionsOptimizeMode = "SPEED"
	FileOptionsOptimizeModeCODE_SIZE    FileOptionsOptimizeMode = "CODE_SIZE"
	FileOptionsOptimizeModeLITE_RUNTIME FileOptionsOptimizeMode = "LITE_RUNTIME"
)

type FileOptions struct {
	present map[string]bool

	// AllFields *Map[FieldDescriptor,Object] `json:"allFields"`

	CcGenericServices bool `json:"ccGenericServices"`

	DefaultInstanceForType *FileOptions `json:"defaultInstanceForType"`

	DescriptorForType *Descriptor `json:"descriptorForType"`

	GoPackage string `json:"goPackage,omitempty"`

	GoPackageBytes *ByteString `json:"goPackageBytes"`

	InitializationErrorString string `json:"initializationErrorString,omitempty"`

	Initialized bool `json:"initialized"`

	JavaGenerateEqualsAndHash bool `json:"javaGenerateEqualsAndHash"`

	JavaGenericServices bool `json:"javaGenericServices"`

	JavaMultipleFiles bool `json:"javaMultipleFiles"`

	JavaOuterClassname string `json:"javaOuterClassname,omitempty"`

	JavaOuterClassnameBytes *ByteString `json:"javaOuterClassnameBytes"`

	JavaPackage string `json:"javaPackage,omitempty"`

	JavaPackageBytes *ByteString `json:"javaPackageBytes"`

	OptimizeFor FileOptionsOptimizeMode `json:"optimizeFor"`

	// ParserForType *com.google.protobuf.Parser<com.google.protobuf.DescriptorProtos$FileOptions> `json:"parserForType"`

	PyGenericServices bool `json:"pyGenericServices"`

	SerializedSize int32 `json:"serializedSize"`

	UninterpretedOptionCount int32 `json:"uninterpretedOptionCount"`

	// UninterpretedOptionList *List[UninterpretedOption] `json:"uninterpretedOptionList"`

	// UninterpretedOptionOrBuilderList *List[? extends com.google.protobuf.DescriptorProtos$UninterpretedOptionOrBuilder] `json:"uninterpretedOptionOrBuilderList"`

	UnknownFields *UnknownFieldSet `json:"unknownFields"`
}

func (self *FileOptions) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *FileOptions) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*FileOptions); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A FileOptions cannot copy the values from %#v", other)
}

func (self *FileOptions) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *FileOptions) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *FileOptions) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *FileOptions) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *FileOptions) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on FileOptions", name)

	case "ccGenericServices", "CcGenericServices":
		v, ok := value.(bool)
		if ok {
			self.CcGenericServices = v
			self.present["ccGenericServices"] = true
			return nil
		} else {
			return fmt.Errorf("Field ccGenericServices/CcGenericServices: value %v(%T) couldn't be cast to type bool", value, value)
		}

	case "defaultInstanceForType", "DefaultInstanceForType":
		v, ok := value.(*FileOptions)
		if ok {
			self.DefaultInstanceForType = v
			self.present["defaultInstanceForType"] = true
			return nil
		} else {
			return fmt.Errorf("Field defaultInstanceForType/DefaultInstanceForType: value %v(%T) couldn't be cast to type *FileOptions", value, value)
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

	case "goPackage", "GoPackage":
		v, ok := value.(string)
		if ok {
			self.GoPackage = v
			self.present["goPackage"] = true
			return nil
		} else {
			return fmt.Errorf("Field goPackage/GoPackage: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "goPackageBytes", "GoPackageBytes":
		v, ok := value.(*ByteString)
		if ok {
			self.GoPackageBytes = v
			self.present["goPackageBytes"] = true
			return nil
		} else {
			return fmt.Errorf("Field goPackageBytes/GoPackageBytes: value %v(%T) couldn't be cast to type *ByteString", value, value)
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

	case "javaGenerateEqualsAndHash", "JavaGenerateEqualsAndHash":
		v, ok := value.(bool)
		if ok {
			self.JavaGenerateEqualsAndHash = v
			self.present["javaGenerateEqualsAndHash"] = true
			return nil
		} else {
			return fmt.Errorf("Field javaGenerateEqualsAndHash/JavaGenerateEqualsAndHash: value %v(%T) couldn't be cast to type bool", value, value)
		}

	case "javaGenericServices", "JavaGenericServices":
		v, ok := value.(bool)
		if ok {
			self.JavaGenericServices = v
			self.present["javaGenericServices"] = true
			return nil
		} else {
			return fmt.Errorf("Field javaGenericServices/JavaGenericServices: value %v(%T) couldn't be cast to type bool", value, value)
		}

	case "javaMultipleFiles", "JavaMultipleFiles":
		v, ok := value.(bool)
		if ok {
			self.JavaMultipleFiles = v
			self.present["javaMultipleFiles"] = true
			return nil
		} else {
			return fmt.Errorf("Field javaMultipleFiles/JavaMultipleFiles: value %v(%T) couldn't be cast to type bool", value, value)
		}

	case "javaOuterClassname", "JavaOuterClassname":
		v, ok := value.(string)
		if ok {
			self.JavaOuterClassname = v
			self.present["javaOuterClassname"] = true
			return nil
		} else {
			return fmt.Errorf("Field javaOuterClassname/JavaOuterClassname: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "javaOuterClassnameBytes", "JavaOuterClassnameBytes":
		v, ok := value.(*ByteString)
		if ok {
			self.JavaOuterClassnameBytes = v
			self.present["javaOuterClassnameBytes"] = true
			return nil
		} else {
			return fmt.Errorf("Field javaOuterClassnameBytes/JavaOuterClassnameBytes: value %v(%T) couldn't be cast to type *ByteString", value, value)
		}

	case "javaPackage", "JavaPackage":
		v, ok := value.(string)
		if ok {
			self.JavaPackage = v
			self.present["javaPackage"] = true
			return nil
		} else {
			return fmt.Errorf("Field javaPackage/JavaPackage: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "javaPackageBytes", "JavaPackageBytes":
		v, ok := value.(*ByteString)
		if ok {
			self.JavaPackageBytes = v
			self.present["javaPackageBytes"] = true
			return nil
		} else {
			return fmt.Errorf("Field javaPackageBytes/JavaPackageBytes: value %v(%T) couldn't be cast to type *ByteString", value, value)
		}

	case "optimizeFor", "OptimizeFor":
		v, ok := value.(FileOptionsOptimizeMode)
		if ok {
			self.OptimizeFor = v
			self.present["optimizeFor"] = true
			return nil
		} else {
			return fmt.Errorf("Field optimizeFor/OptimizeFor: value %v(%T) couldn't be cast to type FileOptionsOptimizeMode", value, value)
		}

	case "pyGenericServices", "PyGenericServices":
		v, ok := value.(bool)
		if ok {
			self.PyGenericServices = v
			self.present["pyGenericServices"] = true
			return nil
		} else {
			return fmt.Errorf("Field pyGenericServices/PyGenericServices: value %v(%T) couldn't be cast to type bool", value, value)
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

	case "uninterpretedOptionCount", "UninterpretedOptionCount":
		v, ok := value.(int32)
		if ok {
			self.UninterpretedOptionCount = v
			self.present["uninterpretedOptionCount"] = true
			return nil
		} else {
			return fmt.Errorf("Field uninterpretedOptionCount/UninterpretedOptionCount: value %v(%T) couldn't be cast to type int32", value, value)
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

func (self *FileOptions) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on FileOptions", name)

	case "ccGenericServices", "CcGenericServices":
		if self.present != nil {
			if _, ok := self.present["ccGenericServices"]; ok {
				return self.CcGenericServices, nil
			}
		}
		return nil, fmt.Errorf("Field CcGenericServices no set on CcGenericServices %+v", self)

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

	case "goPackage", "GoPackage":
		if self.present != nil {
			if _, ok := self.present["goPackage"]; ok {
				return self.GoPackage, nil
			}
		}
		return nil, fmt.Errorf("Field GoPackage no set on GoPackage %+v", self)

	case "goPackageBytes", "GoPackageBytes":
		if self.present != nil {
			if _, ok := self.present["goPackageBytes"]; ok {
				return self.GoPackageBytes, nil
			}
		}
		return nil, fmt.Errorf("Field GoPackageBytes no set on GoPackageBytes %+v", self)

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

	case "javaGenerateEqualsAndHash", "JavaGenerateEqualsAndHash":
		if self.present != nil {
			if _, ok := self.present["javaGenerateEqualsAndHash"]; ok {
				return self.JavaGenerateEqualsAndHash, nil
			}
		}
		return nil, fmt.Errorf("Field JavaGenerateEqualsAndHash no set on JavaGenerateEqualsAndHash %+v", self)

	case "javaGenericServices", "JavaGenericServices":
		if self.present != nil {
			if _, ok := self.present["javaGenericServices"]; ok {
				return self.JavaGenericServices, nil
			}
		}
		return nil, fmt.Errorf("Field JavaGenericServices no set on JavaGenericServices %+v", self)

	case "javaMultipleFiles", "JavaMultipleFiles":
		if self.present != nil {
			if _, ok := self.present["javaMultipleFiles"]; ok {
				return self.JavaMultipleFiles, nil
			}
		}
		return nil, fmt.Errorf("Field JavaMultipleFiles no set on JavaMultipleFiles %+v", self)

	case "javaOuterClassname", "JavaOuterClassname":
		if self.present != nil {
			if _, ok := self.present["javaOuterClassname"]; ok {
				return self.JavaOuterClassname, nil
			}
		}
		return nil, fmt.Errorf("Field JavaOuterClassname no set on JavaOuterClassname %+v", self)

	case "javaOuterClassnameBytes", "JavaOuterClassnameBytes":
		if self.present != nil {
			if _, ok := self.present["javaOuterClassnameBytes"]; ok {
				return self.JavaOuterClassnameBytes, nil
			}
		}
		return nil, fmt.Errorf("Field JavaOuterClassnameBytes no set on JavaOuterClassnameBytes %+v", self)

	case "javaPackage", "JavaPackage":
		if self.present != nil {
			if _, ok := self.present["javaPackage"]; ok {
				return self.JavaPackage, nil
			}
		}
		return nil, fmt.Errorf("Field JavaPackage no set on JavaPackage %+v", self)

	case "javaPackageBytes", "JavaPackageBytes":
		if self.present != nil {
			if _, ok := self.present["javaPackageBytes"]; ok {
				return self.JavaPackageBytes, nil
			}
		}
		return nil, fmt.Errorf("Field JavaPackageBytes no set on JavaPackageBytes %+v", self)

	case "optimizeFor", "OptimizeFor":
		if self.present != nil {
			if _, ok := self.present["optimizeFor"]; ok {
				return self.OptimizeFor, nil
			}
		}
		return nil, fmt.Errorf("Field OptimizeFor no set on OptimizeFor %+v", self)

	case "pyGenericServices", "PyGenericServices":
		if self.present != nil {
			if _, ok := self.present["pyGenericServices"]; ok {
				return self.PyGenericServices, nil
			}
		}
		return nil, fmt.Errorf("Field PyGenericServices no set on PyGenericServices %+v", self)

	case "serializedSize", "SerializedSize":
		if self.present != nil {
			if _, ok := self.present["serializedSize"]; ok {
				return self.SerializedSize, nil
			}
		}
		return nil, fmt.Errorf("Field SerializedSize no set on SerializedSize %+v", self)

	case "uninterpretedOptionCount", "UninterpretedOptionCount":
		if self.present != nil {
			if _, ok := self.present["uninterpretedOptionCount"]; ok {
				return self.UninterpretedOptionCount, nil
			}
		}
		return nil, fmt.Errorf("Field UninterpretedOptionCount no set on UninterpretedOptionCount %+v", self)

	case "unknownFields", "UnknownFields":
		if self.present != nil {
			if _, ok := self.present["unknownFields"]; ok {
				return self.UnknownFields, nil
			}
		}
		return nil, fmt.Errorf("Field UnknownFields no set on UnknownFields %+v", self)

	}
}

func (self *FileOptions) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on FileOptions", name)

	case "ccGenericServices", "CcGenericServices":
		self.present["ccGenericServices"] = false

	case "defaultInstanceForType", "DefaultInstanceForType":
		self.present["defaultInstanceForType"] = false

	case "descriptorForType", "DescriptorForType":
		self.present["descriptorForType"] = false

	case "goPackage", "GoPackage":
		self.present["goPackage"] = false

	case "goPackageBytes", "GoPackageBytes":
		self.present["goPackageBytes"] = false

	case "initializationErrorString", "InitializationErrorString":
		self.present["initializationErrorString"] = false

	case "initialized", "Initialized":
		self.present["initialized"] = false

	case "javaGenerateEqualsAndHash", "JavaGenerateEqualsAndHash":
		self.present["javaGenerateEqualsAndHash"] = false

	case "javaGenericServices", "JavaGenericServices":
		self.present["javaGenericServices"] = false

	case "javaMultipleFiles", "JavaMultipleFiles":
		self.present["javaMultipleFiles"] = false

	case "javaOuterClassname", "JavaOuterClassname":
		self.present["javaOuterClassname"] = false

	case "javaOuterClassnameBytes", "JavaOuterClassnameBytes":
		self.present["javaOuterClassnameBytes"] = false

	case "javaPackage", "JavaPackage":
		self.present["javaPackage"] = false

	case "javaPackageBytes", "JavaPackageBytes":
		self.present["javaPackageBytes"] = false

	case "optimizeFor", "OptimizeFor":
		self.present["optimizeFor"] = false

	case "pyGenericServices", "PyGenericServices":
		self.present["pyGenericServices"] = false

	case "serializedSize", "SerializedSize":
		self.present["serializedSize"] = false

	case "uninterpretedOptionCount", "UninterpretedOptionCount":
		self.present["uninterpretedOptionCount"] = false

	case "unknownFields", "UnknownFields":
		self.present["unknownFields"] = false

	}

	return nil
}

func (self *FileOptions) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type FileOptionsList []*FileOptions

func (self *FileOptionsList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*FileOptionsList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A FileOptionsList cannot copy the values from %#v", other)
}

func (list *FileOptionsList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *FileOptionsList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *FileOptionsList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
