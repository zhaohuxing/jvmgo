package main

import "flag"
import "fmt"
import "os"

//定义Termial结构体
type Terminal struct {
	helpFlag    bool     //　-help
	versionFlag bool     // -version
	cpOption    string   // -cp
	class       string   // 指定class
	args        []string //参数
}

func parseTerminal() *Terminal {
	//定义一个terminal结构体, 因c,c++,go存在指针,所以yong"&Terminal{}",将地址给terminal
	terminal := &Terminal{}

	/*定义命令行参数*/
	flag.Usage = printUsage
	flag.BoolVar(&terminal.helpFlag, "help", false, "print help message")
	flag.BoolVar(&terminal.helpFlag, "?", false, "print help message")
	flag.BoolVar(&terminal.versionFlag, "version", false, "print version and exit")
	flag.StringVar(&terminal.cpOption, "classpath", "", "classpath")
	flag.StringVar(&terminal.cpOption, "cp", "", "classpath")
	//命令行参数解析
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
