package main

import (
	"fmt"
	"jvmgo/ch06/classpath"
	"jvmgo/ch06/rtda/heap"
	"strings"
)

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
	classLoader := heap.NewClassLoader(cp)
	className := strings.Replace(terminal.class, ".", "/", -1)
	mainClass := classLoader.LoadClass(className)
	mainMethod := mainClass.GetMainMethod()
	if mainMethod != nil {
		interpret(mainMethod)
	} else {
		fmt.Printf("Main method not found in class %s\n", terminal.class)
	}
}
