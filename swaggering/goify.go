package swaggering

import (
	"fmt"
	"regexp"
	"strings"
)

var badTypes = []string{
	".",
	"FieldDescriptor",
	"DockerNetworkType",
	"PortMappingType",
	"ContainerType",
	"DockerVolumeMode",
	"Network",
	"Parameter",
	"PortMapping",
	"Type",
}

func capitalize(word string) string {
	firstRE := regexp.MustCompile(`^.`)
	return firstRE.ReplaceAllStringFunc(word, func(match string) string {
		return strings.ToTitle(match)
	})
}

func goName(name string) string {
	if name == "Object" {
		return "interface{}"
	}
	return name
}

func isAggregate(kind string) bool {
	listRE := regexp.MustCompile(`^List\[([^,]*)]`)
	mapRE := regexp.MustCompile(`^Map\[([^,]*),([^,]*)]`)
	return mapRE.FindStringSubmatch(kind) != nil || listRE.FindStringSubmatch(kind) != nil
}

func (self *DataType) findModel(context *Context) (err error) {
	return context.modelFor(self.Type, self)
}

func (self *Parameter) findGoType(context *Context) (err error) {
	if self.ParamType == "body" {
		err = self.findModel(context)
	} else {
		err = self.DataType.findGoType(context)
	}

	return
}

func (self *Operation) findGoType(context *Context) (err error) {
	switch self.Type {
	case "void":
		self.GoType = ""
	case "array":
		self.Collection.findGoType(context)
	case "":
		/* Singularity's swagger has some bugs... */
		self.Type = "array"
		self.Collection.findGoType(context)
	case "string", "bool", "integer", "number":
		typeName, err := self.goPrimitiveType()
		self.setGoType(typeName, err)
	default:
		self.findModel(context)
	}

	return
}

func (self *Collection) findGoType(context *Context) (err error) {
	if self.Type == "array" {
		return findGoType(context, &self.Items, &self.DataType)
	} else {
		return findGoType(context, &self.DataType, &self.DataType)
	}
}

func (self *DataType) findGoType(context *Context) (err error) {
	return findGoType(context, self, self)
}

func findGoType(context *Context, from, to *DataType) (err error) {
	var typeName string

	if from.Type == "" {
		err = context.modelFor(from.Ref, to)
	} else {
		typeName, err = from.goPrimitiveType()
		to.setGoType(typeName, err)
	}
	return
}

func (self *DataType) goPrimitiveType() (t string, err error) {
	switch self.Type {
	case "boolean":
		t = "bool"
	case "integer":
		t = self.Format
	case "number":
		switch self.Format {
		case "float":
			t = "float32"
		case "double":
			t = "float64"
		default:
			err = fmt.Errorf("Invalid number format: %s", self.Format)
		}
	case "string":
		switch self.Format {
		case "", "byte":
			t = "string"
		case "date", "data-time":
			t = "time.Time"
		default:
			err = fmt.Errorf("Invalid string format: %s", self.Format)
		}
	}
	return
}

func (self *DataType) setGoType(typeName string, err error) {
	if err != nil {
		self.GoTypeInvalid = true
	}
	self.GoType = typeName
}
