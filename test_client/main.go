package main

import (
	"fmt"

	"github.com/opentable/sous-singularity/client"
)

func main() {
	singularity := client.New("http://192.168.99.100:7099/singularity")
	state, err := singularity.GetState()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(state.FormatText())
	}
}
