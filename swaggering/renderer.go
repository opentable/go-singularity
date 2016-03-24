package swaggering

import (
	"bytes"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

const (
	defaultModelTmpl = `
{{- /**/ -}}
package dtos

import "io"

type {{.GoName}} struct {
{{range $name, $prop := .Properties}}
{{- if $prop.GoTypeInvalid}}//{{end}}	{{$prop.GoName}} {{$prop.GoType}}
{{end}}}

func (self *{{.GoName}}) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *{{.GoName}}) FormatText() string {
	return FormatText(self)
}

func (self *{{.GoName}}) FormatJSON() string {
	return FormatJSON(self)
}
{{/**/ -}}
`
	defaultApiTmpl = `
{{- /**/ -}}
package client

{{range .Operations}}
func (client *client.Client){{.Nickname}}(
	{{- range .Parameters}}{{.Name}} {{.GoType}}, {{end -}}
) (response *Something) {
	pathParamMap := map[string]interface{}{
		{{range .Parameters -}}
		{{if eq "path" .ParamType -}}
		  "{{.Name}}": {{.Name}},
	  {{- end }}
		{{- end }}
	}
	queryParamMap := map[string]interface{}{
		{{range .Parameters -}}
		{{if eq "query" .ParamType -}}
		  "{{.Name}}": {{.Name}},
	  {{- end }}
		{{- end }}
	}
	client.Request({{.Method}}, pathParamMap, queryParamMap, body)
}
{{end}}
{{- /**/ -}}
	`
)

type Renderer struct {
	targetDir          string
	modelTmpl, apiTmpl *template.Template
}

func NewRenderer(tgt string) (renderer *Renderer) {
	renderer = &Renderer{targetDir: tgt}
	renderer.apiTmpl = template.Must(template.New("api").Parse(defaultApiTmpl))
	renderer.modelTmpl = template.Must(template.New("model").Parse(defaultModelTmpl))

	return
}

func (self *Renderer) RenderService(ingester *Ingester) {
	defer log.SetFlags(log.Flags())
	log.SetFlags(log.Flags() | log.Lshortfile)
	for _, model := range ingester.models {
		if model.GoUses {
			path := filepath.Join("dtos", snakeCase(model.GoName))
			self.renderModel(path, model)
		}
	}

	for _, api := range ingester.apis {
		self.renderApi(apiPath(api.Path), api)
	}
}

func apiPath(urlPath string) string {
	slashes := regexp.MustCompile("/+")
	unders := regexp.MustCompile("_+")
	urlPath = slashes.ReplaceAllString(urlPath, "_")
	urlPath = unders.ReplaceAllString(urlPath, "_")
	return urlPath
}

func snakeCase(symbol string) string {
	write := make([]byte, 0, len(symbol)*2)
	read := []byte(symbol)
	re := regexp.MustCompile("[A-Z]+")

	for {
		found := re.FindIndex(read)

		if found == nil {
			write = append(write, bytes.ToLower(read)...)
			return string(write)
		}

		write = append(write, read[0:found[0]]...)
		if found[0] != 0 {
			write = append(write, byte("_"[0]))
		}
		write = append(write, bytes.ToLower(read[found[0]:found[1]])...)
		read = read[found[1]:]
	}
}

func (self *Renderer) renderModel(path string, model *Model) {
	renderCode(self.targetDir, path, self.modelTmpl, model)
}

func (self *Renderer) renderApi(path string, api *Api) {
	renderCode(self.targetDir, path, self.apiTmpl, api)
}

func renderCode(dir, path string, tmpl *template.Template, context interface{}) {
	fullpath := filepath.Join(dir, path) + ".go"
	log.Print("Rendering to ", fullpath)
	targetFile, err := os.Create(fullpath)
	if err != nil {
		log.Fatal("Problem rendering ", tmpl.Name, " to ", targetFile, ": ", err)
	}
	err = tmpl.Execute(targetFile, context)
	if err != nil {
		log.Fatal("Problem rendering ", tmpl.Name, " to ", targetFile, ": ", err)
	}
}
