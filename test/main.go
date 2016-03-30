package main

import (
	"io"
	"io/ioutil"
	"log"
	"net"
	"os/exec"
	"regexp"
	"strings"
	"syscall"
	"time"

	"github.com/opentable/sous-singularity/client"
	"github.com/opentable/sous-singularity/client/dtos"
)

func dockerMachine(args ...string) (stdoutStr, stderrStr string, err error) {
	dmCmd := exec.Command("docker-machine", args...)
	stdout, err := dmCmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	stderr, err := dmCmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}

	err = dmCmd.Start()
	if err != nil {
		log.Fatal(err)
	}

	stdoutStr = pipeString(stdout)
	stderrStr = pipeString(stderr)

	err = dmCmd.Wait()

	if err != nil {
		log.Print(dmCmd.Args, err)
		log.Print("\n", stdoutStr)
		log.Print("\n", stderrStr)
	}

	return
}

func pipeString(pipe io.Reader) string {
	bytes, err := ioutil.ReadAll(pipe)
	if err != nil {
		log.Fatal(err)
	}
	return string(bytes)
}

func dockerUp(machineName, singDir, ip string) (upCmd *exec.Cmd) {
	setPrefix := regexp.MustCompile("^SET ")
	envStr, _, err := dockerMachine("env", "--shell", "cmd", machineName) //other shells get extraneous quotes
	if err != nil {
		log.Panic("env", err)
	}

	env := make([]string, 0, 4)
	for _, str := range strings.Split(envStr, "\n") {
		env = append(env, setPrefix.ReplaceAllString(str, ""))
	}

	upCmd = exec.Command("docker-compose", "up")
	upCmd.Env = env
	upCmd.Dir = singDir
	err = upCmd.Start()
	if err != nil {
		log.Panic(err)
	}
	for times := 15; times > 0; times-- {
		if singularityRunning(ip) {
			log.Print("Singularity up and accepting connections")
			return
		}
		time.Sleep(1 * time.Second)
	}
	panic("Singularity not available!")
}

func dockerDown(cmd *exec.Cmd) (err error) {
	cmd.Process.Signal(syscall.SIGTERM)
	return cmd.Wait()
}

func singularityRunning(ip string) bool {
	for times := 10; times > 0; times-- {
		conn, err := net.Dial("tcp", ip+":7099")

		if err != nil {
			if nerr, ok := err.(*net.OpError); ok && nerr.Temporary() {
				time.Sleep(1 * time.Second)
				continue
			}
			return false
		}

		conn.Close()
		return true
	}
	return false
}

func main() {
	machineName := "default"
	singDir := "/Users/jlester/sous/Singularity"

	alreadyStarted := regexp.MustCompile("is already running")

	log.SetFlags(log.Flags() | log.Lshortfile)
	_, stderr, err := dockerMachine("start", machineName)
	if err != nil && !alreadyStarted.MatchString(stderr) {
		log.Panic("start", err)
	}

	ip, _, err := dockerMachine("ip", machineName)
	if err != nil {
		log.Panic("ip", err)
	}
	ip = strings.Trim(ip, " \n\t")

	log.Print(ip)

	if !singularityRunning(ip) {
		log.Print("Singularity needs to be started - tip: running `docker-compose up` for singularity will speed tests up.")

		upCmd := dockerUp(machineName, singDir, ip)
		defer dockerDown(upCmd)
	} else {
		log.Print("Singularity already up and running")
	}

	var obj dtos.DTO
	singularity := client.New("http://" + ip + ":7099/singularity")

	obj, err = singularity.GetState(true, true)
	output("GetState", obj, err)

}

func output(title string, thing dtos.DTO, err error) {
	if err != nil {
		log.Panic(title, " ERR: ", err)
	} else {
		log.Print(title, "\n", thing.FormatJSON())
	}
}
