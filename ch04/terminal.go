package main

import "flag"
import "fmt"
import "os"

//定义Termial结构体
type Terminal struct {
	helpFlag    bool
	versionFlag bool
	cpOption    string
	XjreOption  string // 添加XjreOption字段
	class       string
	args        []string
}

func parseTerminal() *Terminal {
	terminal := &Terminal{}
	flag.Usage = printUsage
	flag.BoolVar(&terminal.helpFlag, "help", false, "print help message")
	flag.BoolVar(&terminal.helpFlag, "?", false, "print help message")
	flag.BoolVar(&terminal.versionFlag, "version", false, "print version and exit")
	flag.StringVar(&terminal.cpOption, "classpath", "", "classpath")
	flag.StringVar(&terminal.cpOption, "cp", "", "classpath")
	flag.StringVar(&terminal.XjreOption, "Xjre", "", "path to jre")
	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {
		terminal.class = args[0]
		terminal.args = args[1:]
	}
	return terminal
}

func printUsage() {
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
}
