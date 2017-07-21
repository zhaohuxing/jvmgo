package main

import "fmt"

func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Println("java 1.9")
	} else if cmd.class != "" {
		startJVM(cmd)
	} else {
		cmdOfUsage()
	}
}

func startJVM(cmd *Cmd) {
	fmt.Println("startJVM")
}
