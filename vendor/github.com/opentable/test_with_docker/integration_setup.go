package test_with_docker

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os/exec"
	"regexp"
	"strings"
	"syscall"
	"time"
)

func dockerMachine(args ...string) (stdoutStr, stderrStr string, err error) {
	dmCmd := exec.Command("docker-machine", args...)
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	dmCmd.Stdout = &stdout
	dmCmd.Stderr = &stderr

	err = dmCmd.Start()
	if err != nil {
		log.Fatal(err)
	}

	err = dmCmd.Wait()
	stdoutStr = stdout.String()
	stderrStr = stderr.String()

	//	if err != nil {
	log.Print(dmCmd.Args, err)
	log.Print("\n", stdoutStr)
	log.Print("\n", stderrStr)
	//	}

	return
}

func pipeString(pipe io.Reader) string {
	bytes, err := ioutil.ReadAll(pipe)
	if err != nil {
		log.Fatal(err)
	}
	return string(bytes)
}

type serviceMap map[string]uint

func dockerComposeUp(machineName, dir, ip string, services serviceMap) (upCmd *exec.Cmd) {
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
	upCmd.Dir = dir
	err = upCmd.Start()
	if err != nil {
		log.Panic(err)
	}

	if servicesRunning(30.0, ip, services) {
		return
	}
	panic(fmt.Sprintf("Services were not available!"))
}

func dockerDown(cmd *exec.Cmd) (err error) {
	cmd.Process.Signal(syscall.SIGTERM)
	return cmd.Wait()
}

func servicesRunning(limit float32, ip string, services map[string]uint) bool {
	goodCh := make(chan string)
	goodCount := 0

	for name, port := range services {
		go func() {
			if serviceRunning(ip, port) {
				goodCh <- name
			}
		}()
	}

	for goodCount < len(services) {
		select {
		case good := <-goodCh:
			log.Printf("%s up and running", good)
			goodCount += 1
		case <-time.After(time.Duration(limit * float32(time.Second))):
			log.Printf("Attempt to contact remaining service expired")
			return false
		}
	}
	return true
}

func serviceRunning(ip string, port uint) bool {
	for times := 10; times > 0; times-- {
		conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", ip, port))

		if err != nil {
			log.Print(err)
			if nerr, ok := err.(*net.OpError); ok && nerr.Temporary() {
				time.Sleep(time.Duration(0.5 * float32(time.Second)))
				continue
			}
			return false
		}

		conn.Close()
		return true
	}
	return false
}

var ip string

func ComposeServices(machineName, dir string, servicePorts serviceMap) (ip string, err error) {
	alreadyStarted := regexp.MustCompile("is already running")
	_, stderr, err := dockerMachine("start", machineName)
	if err != nil && !alreadyStarted.MatchString(stderr) {
		log.Panic("start", err)
	}

	ip, _, err = dockerMachine("ip", machineName)
	if err != nil {
		return
	}
	ip = strings.Trim(ip, " \n\t")

	log.Print(ip)

	if !servicesRunning(3.0, ip, servicePorts) {
		log.Printf("Services need to be started - tip: running `docker-compose up` in %s will speed tests up.", dir)

		upCmd := dockerComposeUp(machineName, dir, ip, servicePorts)
		defer dockerDown(upCmd)
	} else {
		log.Printf("All services already up and running")
	}

	return
}
