package main

import (
	"log"
	"os"

	"github.com/opentable/sous-singularity/swaggering"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Usage:", os.Args[0], "<path to Singularity api-doc dir> <path to client root dir>")
	}
	serviceSource := os.Args[1]
	renderTarget := os.Args[2]

	context := swaggering.NewContext()

	swaggering.ProcessService(serviceSource, context)
	swaggering.ResolveService(context)
	swaggering.RenderService(renderTarget, context)
}
