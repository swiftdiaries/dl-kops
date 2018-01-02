package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

//ExecuteThroughSSH takes in a script file path and executes it on a host machine and returns the output logs
func ExecuteThroughSSH(hostname string, hostip string, keyfile string, path string) []string {

	sshcmd := "ssh"
	var args []string
	args = []string{"-i", keyfile, hostname + "@" + hostip}
	var output []string
	inFile, _ := os.Open(path)
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		temparg := append(args, scanner.Text())
		cmd := exec.Command(sshcmd, temparg...)
		cmd.Stdin = strings.NewReader("")
		var out bytes.Buffer
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s", out.String())
		output = append(output, out.String())
	}
	return output
}

//PrintLines used to print line by line
func PrintLines(path string) []string {
	inFile, _ := os.Open(path)
	var output []string
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		output = append(output, scanner.Text())
	}
	return output
}
