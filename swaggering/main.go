package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"text/template"
)

type (
	Swagger struct {
		BasePath, ResourcePath string
		Apis                   []Api
		Models                 map[string]Model
	}

	Api struct {
		Path, Description string
		Operations        []Operation
	}

	Operation struct {
		Method, Nickname, Deprecated string
		Parameters                   []Parameter
	}

	Parameter struct {
		ParamType, Name         string
		Required, AllowMultiple bool
	}

	Model struct {
		Id, Description, Discriminator string
		GoName                         string
		Required, SubTypes             []string
		Properties                     map[string]Property
	}

	Property struct {
		GoType       string
		Type, Format string
		Ref          string `json:"$ref"`
		Enum         []string
		Items        Item
		UniqueItems  bool
	}

	Item struct {
		Type, Format string
		Ref          string `json:"$ref"`
	}
)

type Renderer struct {
	modelTmpl *template.Template
}

func main() {
	swaggerFiles := []string{
		"api_deploys.json",
	}

	if len(os.Args) < 2 {
		log.Fatal("Usage:", os.Args[0], "<path to Singularity api-doc dir>")
	}
	path := os.Args[1]
	renderer := &Renderer{}
	renderer.setup()

	for _, file := range swaggerFiles {
		renderer.process(filepath.Join(path, file))
	}
}

func (renderer *Renderer) setup() {
	renderer.modelTmpl = template.Must(template.New("model").Parse(`
{{.GoName}}{
	{{range $name, $prop := .Properties}}{{$name}}: {{$prop.GoType}}
	{{end}}
}`))
	/*
		package dtos

		import "io"

		type Deploy struct {
			//Resources [com.hubspot.mesos.Resources]
			//ContainerInfo [SingularityContainerInfo]
			//HealthcheckProtocol [HealthcheckProtocol]
			//LoadBalancerGroups [Set]
			//ExecutorData [ExecutorData]
			//CustomExecutorResources [Resources]

			RequestId string
			Metadata  map[string]string
			Labels    map[string]string
			Version   string
			Id        string

			Uris            []string
			ServiceBasePath string

			Command   string
			Arguments []string
			Env       map[string]string

			Timestamp uint64

			CustomExecutorId     string
			CustomExecutorUser   string
			CustomExecutorSource string
			CustomExecutorCmd    string

			HealthcheckTimeoutSeconds             uint64
			HealthcheckMaxRetries                 int
			HealthcheckMaxTotalTimeoutSeconds     uint64
			HealthcheckUri                        string
			SkipHealthchecksOnDeploy              bool
			ConsiderHealthyAfterRunningForSeconds uint64
			DeployHealthTimeoutSeconds            uint64
			HealthcheckIntervalSeconds            uint64

			LoadBalancerOptions map[string]interface{}
		}

		type Deploys []Deploy

		func (deps *Deploys) Get(client APIClient) (err error) {
			return client.APIGet("deploys", deps)
		}

		func (deps *Deploys) Populate(jsonReader io.ReadCloser) (err error) {
			err = ReadPopulate(jsonReader, deps)
			return
		}

		func (deps *Deploys) FormatText() string {
			return FormatText(deps)
		}

		func (deps *Deploys) FormatJSON() string {
			return FormatJSON(deps)
		}
		`)
	*/
}

func (renderer *Renderer) process(fullpath string) {
	data, err := os.Open(fullpath)
	if err != nil {
		log.Print("Trouble with", fullpath, ":", err)
		return
	}

	dec := json.NewDecoder(data)
	var swagger *Swagger = &Swagger{}

	if err := dec.Decode(&swagger); err == io.EOF {
		return
	} else if err != nil {
		log.Print("Trouble parsing", fullpath, ":", err)
		return
	}

	goify(swagger)
	renderer.render(swagger)
}

func goify(swagger *Swagger) {

	for name, model := range swagger.Models {
		model.GoName = goName(model.Id)
		for name, property := range model.Properties {
			property.GoType = goType(&property)
			model.Properties[name] = property
		}
		swagger.Models[name] = model
	}
}

func goName(name string) string {
	listRE := regexp.MustCompile(`^List\[([^,]*)]`)
	mapRE := regexp.MustCompile(`^Map\[([^,]*),([^,]*)]`)
	singRE := regexp.MustCompile("^Singularity")
	maps := mapRE.FindStringSubmatch(name)
	if maps == nil {
		lists := listRE.FindStringSubmatch(name)
		if lists == nil {
			return singRE.ReplaceAllString(name, "")
		}
		return fmt.Sprintf("[]%s", goName(lists[1]))
	}
	return fmt.Sprintf("map[%s]%s", goName(maps[1]), goName(maps[2]))
}

func goType(property *Property) string {
	switch property.Type {
	case "boolean":
		return "bool"
	case "array":
		if property.Items.Type == "" {
			return fmt.Sprintf("[]%s", property.Items.Type)
		} else {
			return fmt.Sprintf("[]%s", goName(property.Items.Ref))
		}
	case "":
		return goName(property.Ref)
	default:
		return property.Type
	}
}

func (self *Renderer) render(swagger *Swagger) {
	for _, model := range swagger.Models {
		self.modelTmpl.Execute(os.Stdout, model)
	}
}
