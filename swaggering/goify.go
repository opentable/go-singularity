package swaggering

import (
	"fmt"
	"log"
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

func goify(swagger *Swagger) {

	for name, model := range swagger.Models {
		model.GoName = goName(model.Id)
		for name, property := range model.Properties {
			property.GoName = capitalize(name)
			property.goType()
			/*
				for _, bad := range badTypes {
					if strings.Index(property.goType, bad) != -1 {
						property.goTypeInvalid = true
						break
					}
				}
			*/

			model.Properties[name] = property
		}
		swagger.Models[name] = model
	}
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
	singRE := regexp.MustCompile("^Singularity")
	return singRE.ReplaceAllString(name, "")
}

func isAggregate(kind string) bool {
	listRE := regexp.MustCompile(`^List\[([^,]*)]`)
	mapRE := regexp.MustCompile(`^Map\[([^,]*),([^,]*)]`)
	return mapRE.FindStringSubmatch(kind) != nil || listRE.FindStringSubmatch(kind) != nil
}

func goAggregates(kind string) string {
	listRE := regexp.MustCompile(`^List\[([^,]*)]`)
	mapRE := regexp.MustCompile(`^Map\[([^,]*),([^,]*)]`)

	maps := mapRE.FindStringSubmatch(kind)
	if maps == nil {
		lists := listRE.FindStringSubmatch(kind)
		if lists == nil {
			return goName(kind)
		}
		return fmt.Sprintf("[]%s", goName(lists[1]))
	}
	return fmt.Sprintf("map[%s]%s", goName(maps[1]), goAggregates(maps[2]))
}

func goType(kind string) string {
	log.Print(kind)
	switch kind {
	case "boolean":
		return "bool"
	default:
		return goAggregates(kind)
	}
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

func (property *Property) goType() {
	switch property.Type {
	case "integer":
		property.DataType.GoType = "u" + property.Format
	case "array":
		if property.Items.Type != "" {
			property.DataType.GoType = fmt.Sprintf("[]%s", property.Items.Type)
		} else {
			property.DataType.GoType = fmt.Sprintf("[]%s", goName(property.Items.Ref))
		}
	case "":
		if isAggregate(property.Ref) {
			property.GoTypeInvalid = true
			property.DataType.GoType = property.Ref
			return
		}
		property.DataType.GoType = goName(property.Ref)
	default:
		property.DataType.GoType = goType(property.Type)
	}
}
