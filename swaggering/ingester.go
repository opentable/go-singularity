package swaggering

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

type ServiceJSON struct {
	Apis []*ServiceApiJSON
}

type ServiceApiJSON struct {
	Path, Desc string
}

type Ingester struct {
	sourceDir string
	swaggers  []*Swagger
	apis      []*Api
	models    []*Model
}

func NewIngester(src string) (ingester *Ingester) {
	ingester = &Ingester{sourceDir: src}
	ingester.swaggers = make([]*Swagger, 0)

	return
}

func (ingester *Ingester) ProcessService() {
	dir := ingester.sourceDir
	defer log.SetFlags(log.Flags())
	log.SetFlags(log.Lshortfile)
	fullpath := filepath.Join(dir, "service.json")

	apis := &ServiceJSON{}

	loadJSON(fullpath, apis)

	fileRE := regexp.MustCompile(`^/(\w+).{format}$`)

	for _, api := range apis.Apis {
		smi := fileRE.FindStringSubmatchIndex(api.Path)
		file := []byte("")
		file = fileRE.ExpandString(file, "$1.json", api.Path, smi)

		ingester.ingestApi(filepath.Join(dir, string(file)))
	}

	ingester.resolve()
}

func (self *Ingester) resolve() {
	var err error

	log.Print("Resolving types")

	log.Print(self.swaggers)
	for _, swagger := range self.swaggers {
		log.Printf("%#v", swagger)
		log.Print(len(swagger.Apis))
		for adx := range swagger.Apis {
			self.apis = append(self.apis, &swagger.Apis[adx])
		}
		for _, model := range swagger.Models {
			self.models = append(self.models, model)
		}
	}

	log.Printf("  Found %d apis", len(self.apis))
	log.Printf("  Found %d models", len(self.models))

	openModels := make([]*Model, 0, len(self.models))

	for _, api := range self.apis {
		for _, op := range api.Operations {
			modelName := op.Type
			openModels, err = self.modelUsed(modelName, openModels)
			if err != nil {
				log.Printf("Operation %s invalid: its type %v", op.Nickname, err)
				op.GoTypeInvalid = true
			}

			for _, parm := range op.Parameters {
				if parm.ParamType == "body" {
					openModels, err = self.modelUsed(parm.Type, openModels)
					if err != nil {
						log.Printf("Operation %s invalid: parameter %s type: %v", op.Nickname, parm.Name, err)
						op.GoTypeInvalid = true
					}
				}
			}
		}
	}

	self.resolveModels(openModels)
}

func (self *Ingester) resolveModels(openModels []*Model) {
	var cur *Model
	var err error
	var typeName string

	for len(openModels) > 0 {
		cur, openModels = openModels[0], openModels[1:]
		if cur.GoUses {
			continue
		}
		cur.GoUses = true

		for name, prop := range cur.Properties {
			if prop.Type == "" {
				openModels, err = self.modelUsed(prop.Ref, openModels)
				prop.setGoType(goName(prop.Ref), err, name, cur.Id, false)
			} else {
				if prop.Type == "array" {
					if prop.Items.Type == "" {
						openModels, err = self.modelUsed(prop.Items.Ref, openModels)
						prop.setGoType(goName(prop.Ref), err, name, cur.Id, true)
					} else {
						typeName, err = prop.Collection.Items.goPrimitiveType()
						prop.setGoType(typeName, err, name, cur.Id, true)
					}
				} else {
					typeName, err = prop.goPrimitiveType()
					prop.setGoType(typeName, err, name, cur.Id, false)
				}
			}
		}
	}
}

func (self *DataType) setGoType(typeName string, err error, name, curId string, isArray bool) {
	var prefix, extraMsg string
	if isArray {
		prefix = "[]"
		extraMsg = "array items "
	}
	if err != nil {
		log.Printf("Property %s %s of model %s is invalid: %v", name, extraMsg, curId, err)
		self.GoTypeInvalid = true
	}
	self.GoType = prefix + typeName
}

func (self *Ingester) modelUsed(name string, usedModels []*Model) (newUsed []*Model, err error) {
	newUsed = usedModels
	for _, model := range self.models {
		if model.Id == name {
			if !model.GoUses {
				newUsed = append(usedModels, model)
			}
			return
		}

	}
	err = fmt.Errorf("Model %s doesn't appear in known models.", name)
	return
}

func loadJSON(fullpath string, into interface{}) {
	data, err := os.Open(fullpath)
	if err != nil {
		log.Print("Trouble with", fullpath, ":", err)
		return
	}

	dec := json.NewDecoder(data)

	if err := dec.Decode(into); err == io.EOF {
		log.Fatal("Trouble with empty", fullpath, ":", err)
		return
	} else if err != nil {
		log.Print("Trouble parsing", fullpath, ":", err)
		return
	}
}

func (ingester *Ingester) ingestApi(fullpath string) {
	log.Print("Processing:", fullpath)
	ingester.swaggers = append(ingester.swaggers, &Swagger{})
	swagger := ingester.swaggers[len(ingester.swaggers)-1]
	loadJSON(fullpath, swagger)

	goify(swagger)
}
