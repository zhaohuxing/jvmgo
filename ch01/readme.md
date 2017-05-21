>源于只因遇见Go，被Go的魅力所折服，慢慢将发生下面的故事

## Ubuntu下Go的安装
Go的安装方式有好几种，可以根据自己的系统版本，下载对应的文件．分享下我的安装方式，笔者系统版本是linux(ubuntu)64位，往后都是基于Liunx(Ubuntu)进行分享．
- [下载Go的源码](https://golang.org/dl/)
- 解压源码--->自定义文件夹
- 设置环境变量
  - GOROOT:`export GOROOT=/home/ubuntu/go/`,Go源码文件夹所在地址
  - GOPATH: `export GOPATH=/home/ubuntu/workspace/`,Go代码存放位置
  - PATH:`export PATH=$PATH:$GOROOT/bin:$GOPATH/bin`,设置下系统环境变量

注意：说明下，根据本地脚本的类型在响应的配置文件中修改变量，比如：我的是zsh, 我就需要在`.zshrc`中修改

完成上述操作后，我们可以命令行中敲入`go`,验证是否安装成功,类似`java`命令．

## 小试牛刀
安装Go之后，我们可以在`workspace`中编写我们的Go代码了吗？可以，但是我们要管理我们程序结构，所以`workspace`中需要三个文件夹：
- bin: 编译之后可执行文件
- pkg: 应用包
- src:应用源码(我们编写的代码)

在src目录下创建`jvmgo`工程，再在`jvmgo`文件下创建子目录`ch01`,我工程目录图如下：
![workspace.jpg](http://upload-images.jianshu.io/upload_images/2031765-ac64df805fc63740.jpg?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

简单实现命令行设置，执行效果如下:

![result.jpg](http://upload-images.jianshu.io/upload_images/2031765-eb2906ba753efbba.jpg?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

命令解释:
- go build : 编译代码，默认生成与文件相同的可执行文件
- ch01 -version: ch01,编译后生成的可执行文件，-version, 预定义命令．
- ch01 -help: ch01,编译后生成的可执行文件，-help, 预定义命令．

实现思路:
- 编写命令行工具文件，定义命令行信息，如:terminal.go
- 利用`flag`模块预先定义命令
- 编写main函数测试执行

在/ch01/文件夹下创建terminal.go文件,定义命令行信息,并预定义命令:
```
//定义命令行信息，Termial结构体
type Terminal struct {
	helpFlag    bool     //　-help
	versionFlag bool     // -version
	cpOption    string   // -cp
	class       string   // 指定class
	args        []string //参数
}

//预先定义命令
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
```

在/ch01/文件下创建main.go文件，测试:
```
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

```

注:上述代码算是伪代码吧, 源文件还请移步[源码](https://github.com/zhaohuXing/jvmgo/tree/master/ch01)

#### 参考文献：
- 自己动手写Java虚拟机
- https://github.com/astaxie/build-web-application-with-golang
