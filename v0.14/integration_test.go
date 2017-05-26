package singularity_test

import (
	"log"
	"os"
	"testing"

	"github.com/opentable/go-singularity"
	"github.com/opentable/swaggering"
	"github.com/opentable/test_with_docker"
)

var ip string
var client *singularity.Client

func TestMain(m *testing.M) {
	os.Exit(wrapCompose(m))
}

func wrapCompose(m *testing.M) int {
	test_with_docker := test_with_docker.NewAgent(60.0, "default")
	log.SetFlags(log.Flags() | log.Lshortfile)

	singDir := os.Getenv("SINGULARITY_DIR")

	if len(singDir) == 0 {
		log.Print("Environment variable SINGULARITY_DIR needs to be set to the path to a checkout of https://github.com/HubSpot/Singularity")
		log.Print("  You'll need to `docker pull` in that directory, as well.")
		os.Exit(1)
	}

	started, err := test_with_docker.ComposeServices(singDir, map[string]uint{"Singularity": 7099})
	if started != nil {
		log.Print("deferring shutdown of services")
		defer test_with_docker.Shutdown(started)

	}
	if err != nil {
		log.Panic(err)
	}

	log.Print(test_with_docker.IP())

	ip, _ := test_with_docker.IP()
	client = singularity.NewClient("http://" + ip.String() + ":7099/singularity")

	return m.Run()
}

func TestGetState(t *testing.T) {
	obj, err := client.GetState(true, true)
	output("GetState", obj, err)
	if obj.GeneratedAt == 0 {
		t.Error("Zero value for obj.GeneratedAt - at best suspicious.")
	}

	if obj.AuthDatastoreHealthy != true {
		t.Error("AuthDatastoreHealty not 'true'")
	}
}

func TestGetPendingDeploys(t *testing.T) {
	obj, err := client.GetPendingDeploys()
	output("GetPendingDeploys", &obj, err)
}

func TestGetRacks(t *testing.T) {
	obj, err := client.GetRacks("")
	output("GetRacks", &obj, err)

	if !(len(obj) > 0) {
		t.Error("Rack list is still empty (should have at least one element)")
	}
}

func TestGetRequests(t *testing.T) {
	obj, err := client.GetRequests()
	output("GetRequests", &obj, err)
}

func output(title string, thing swaggering.DTO, err error) {
	if err != nil {
		log.Panic(title, " ERR: ", err)
	} else {
		log.Print(title, "\n", thing.FormatJSON())
	}
}
