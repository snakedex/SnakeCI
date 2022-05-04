package main

import (
	"bytes"
	"log"
	"os/exec"
)

func docker_build(tag string, dockerfile string) error {
	cmd := exec.Command("docker", "build", "-t", tag, "-f", dockerfile)
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return err
	} else {
		return nil
	}
}

func main() {
	err := docker_build("example:test", "Dockerfile-tst")
	log.Fatal(err)
}
