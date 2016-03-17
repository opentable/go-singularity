package main

import (
	"fmt"

	"github.com/opentable/sous-singularity/client"
	"github.com/opentable/sous-singularity/client/dtos"
)

func main() {
	var err error
	var obj dtos.DTO
	singularity := client.New("http://192.168.99.100:7099/singularity")

	obj, err = singularity.GetState()
	output("State:", obj, err)
	obj, err = singularity.GetDeploys()
	output("Deploys:", obj, err)
	obj, err = singularity.GetRequests()
	output("Requests:", obj, err)
}

func output(title string, thing dtos.DTO, err error) {
	if err != nil {
		fmt.Println(title, "ERR:", err)
	} else {
		fmt.Println(title, thing.FormatJSON())
	}
}
