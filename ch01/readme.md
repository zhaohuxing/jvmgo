## 重温(170721)

> 学习 “自己动手编写 JVM” ，本文主要涉及 Go 的 flag 包的使用。


## 实现目标
实现一个简单的命令行工具。栗子：
```
$ java -version
$ Java 1.9
```
## 实现原理
利用 Go 的处理命令行的 package: flag, 来实现。具体查看下文代码实现以及flag 的官方文档。
## 代码实现
在`gopath`的目录下的`src`文件中创建`jvm`工程，并创建 `cmd.go`和`main.go`文件。
- `cmd.go`用来预设置命令行
- `main.go` 用来简单的逻辑执行，也是程序的入口

附上我的工程目录:
![ch01.png](http://upload-images.jianshu.io/upload_images/2031765-0b8d2c714502c801.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

实现步骤：
- 定义 Cmd 结构体
- 预定义命令
- 在 main 中编写简单的逻辑

#### 定义 Cmd 结构体
```
type Cmd struct {
	versionFlag bool
	helpFlag    bool
	cpOption    string
	class       string
	args        []string
}
```
#### 预定义命令
```
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
```
#### 在 main 中编写简单的逻辑
```

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
```

代码编写完后，在`ch01`目录下敲入命令行：
```
$ go build -o cmd
```
这时在当前目录便生成了可执行文件`cmd`，然后进行测试，敲入命令行：
```
$ ./cmd -version
或
$ ./cmd -help
```
附上我的效果图：

![ch01_result_01.png](http://upload-images.jianshu.io/upload_images/2031765-19930a3f8c504062.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)


## Go 加餐
本文涉及到flag中的函数：
-  [func Parse()](https://golang.org/pkg/flag/#Parse) ：用于解析输入的命令行
- [func BoolVar(p *bool, name string, value bool, usage string)](https://golang.org/pkg/flag/#BoolVar)
- [func StringVar(p *string, name string, value string, usage string)](https://golang.org/pkg/flag/#StringVar)

详情查看文档吧。

##参考文献：
- 自己动手写Java虚拟机
- https://github.com/astaxie/build-web-application-with-golang
- 源码：https://github.com/zhaohuXing/jvmgo/tree/master/ch01

精彩文章，持续更新，请关注微信公众号：
![帅哥美女扫一扫](http://upload-images.jianshu.io/upload_images/2031765-2c4654abe66cd4c8.jpg?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)
