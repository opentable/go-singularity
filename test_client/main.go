package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os/exec"
	"regexp"
	"strings"
	"syscall"

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

func dockerUp(env []string, singDir string) (upCmd *exec.Cmd) {
	upCmd = exec.Command("docker-compose", "up")
	upCmd.Env = env
	upCmd.Dir = singDir
	err := upCmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	return
}

func dockerDown(cmd *exec.Cmd) (err error) {
	cmd.Process.Signal(syscall.SIGTERM)
	return cmd.Wait()
}

func main() {
	machineName := "default"
	singDir := "/Users/jlester/sous/Singularity"

	alreadyStarted := regexp.MustCompile("is already running")
	setPrefix := regexp.MustCompile("^SET ")

	log.SetFlags(log.Flags() | log.Lshortfile)
	_, stderr, err := dockerMachine("start", machineName)
	if err != nil && !alreadyStarted.MatchString(stderr) {
		log.Fatal("start", err)
	}
	envStr, _, err := dockerMachine("env", "--shell", "cmd", machineName) //other shells get extraneous quotes
	if err != nil {
		log.Fatal("env", err)
	}

	env := make([]string, 0, 4)
	for _, str := range strings.Split(envStr, "\n") {
		env = append(env, setPrefix.ReplaceAllString(str, ""))
	}

	upCmd := dockerUp(env, singDir)

	ip, _, err := dockerMachine("ip", machineName)
	if err != nil {
		log.Fatal("ip", err)
	}

	log.Print(ip)

	dockerDown(upCmd)
	//var obj dtos.DTO
	//singularity := client.New("http://" + ip + ":7099/singularity")

}

func output(title string, thing dtos.DTO, err error) {
	if err != nil {
		fmt.Println(title, "ERR:", err)
	} else {
		fmt.Println(title, thing.FormatJSON())
	}
}
