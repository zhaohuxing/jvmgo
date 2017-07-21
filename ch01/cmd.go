package main

import (
	"flag"
	"fmt"
)

type Cmd struct {
	versionFlag bool
	helpFlag    bool
	cpOption    string
	class       string
	args        []string
}

func parseCmd() *Cmd {
	cmd := &Cmd{}

	// set command line variable
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version message")
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")

	//parse command line
	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}
	return cmd
}

func cmdOfUsage() {
	message := `
	The commands are:
		version 	print version message
		help 		show command line message
		cp 			set java classpath
	Examples:
		$ cmd -version
		$ java 1.9
	`
	fmt.Println(message)
}
