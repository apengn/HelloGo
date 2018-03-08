package main

import (
	"os/exec"
	"log"
	"fmt"
	"bytes"
)


func main() {
	cmd := exec.Command("dir", "-lah")

	var stdout, stderr bytes.Buffer

	cmd.Stderr = &stderr
	cmd.Stdout = &stdout
	err := cmd.Run()

	if err != nil {
		log.Fatalf("cmd.run() failed with %s\n", err)
	}

	outStr,errStr:=string(stdout.Bytes()),string(stderr.Bytes())


	fmt.Printf("out:\n%s\n err:\n%s\n",outStr,errStr)
}
