package main

import "fmt"
import "strings"
import "jvmgo/ch03/classpath"

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
	cp := classpath.Parse(terminal.XjreOption, terminal.cpOption)
	fmt.Printf("classpath:%v class:%v args:%v\n",
		cp, terminal.class, terminal.args)
	className := strings.Replace(terminal.class, ".", "/", -1)
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		fmt.Printf("Could not find or load main class %s\n", terminal.class)
	}
	fmt.Printf("class data:%v\n", classData)
}
