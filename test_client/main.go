package main

import (
	"fmt"

	"github.com/opentable/sous-singularity/client"
)

func main() {
	singularity := client.New()
	state, err := singularity.GetState()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(state.FormatText())
	}
}
