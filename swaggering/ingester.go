package swaggering

import (
	"encoding/json"
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

func ProcessService(dir string, ingester *Context) {
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

func (self *Context) ingestApi(fullpath string) {
	log.Print("Processing:", fullpath)
	self.swaggers = append(self.swaggers, &Swagger{})
	swagger := self.swaggers[len(self.swaggers)-1]
	loadJSON(fullpath, swagger)
}
