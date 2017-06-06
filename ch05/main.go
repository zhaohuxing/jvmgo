package main

import (
	"fmt"
	"jvmgo/ch05/classfile"
	"jvmgo/ch05/classpath"
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
	className := strings.Replace(terminal.class, ".", "/", -1)
	cf := loadClass(className, cp)
	mainMethod := getMainMethod(cf)
	if mainMethod != nil {
		interpret(mainMethod)
	} else {
		fmt.Printf("Main method not found in class %s\n", terminal.class)
	}
}

func loadClass(className string, cp *classpath.Classpath) *classfile.ClassFile {
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		panic(err)
	}

	cf, err := classfile.Parse(classData)
	if err != nil {
		panic(err)
	}
	return cf
}

func getMainMethod(cf *classfile.ClassFile) *classfile.MemberInfo {
	for _, m := range cf.Methods() {
		if m.Name() == "main" && m.Descriptor() == "([Ljava/lang/String;)V" {
			return m
		}
	}
	return nil
}
