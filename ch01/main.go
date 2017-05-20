package main

import "fmt"

func main() {
	terminal := parseTerminal()
	if terminal.versionFlag {
		fmt.Println("version 0.0.1")
	} else if terminal.helpFlag || terminal.class == "" {
		printUsage()
	} else {
		startJVM(terminal)
	}
}

func startJVM(terminal *Terminal) {
	fmt.Printf("classpath:%s class:%s args:%v\n", terminal.cpOption, terminal.class, terminal.args)
}
