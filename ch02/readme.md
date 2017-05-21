> 源于只因遇见Go，被Go的魅力所折服，慢慢将发生下面的故事

## Java类加载机制
Java类加载器的作用，将class文件加载到内存。从Java开发人员的角度来说，类加载器可分为三种：
- 启动类加载器(Bootstrap ClassLoader)：加载`<JAVA_HOME>/jre/lib`下的jar
- 扩展类加载器(Extension ClassLoader)：加载`<JAVA_HOME>/jre/lib/ext/`下的jar
- 应用程序类加载器(Application ClassLoader)：加载用户自己编写的class

大多数的加载器使用`双亲委派模型`，双亲委派模型的工作过程是： 如果一个类加载器收到了类加载器的请求，它首先不会自己去尝试加载这类，而是把这个请求委派给父类加载器去完成，每一个层次的类加载器都是如此，因此所有的加载请求最终都应该传送到顶层的启动类加载器中，只有当父类加载器反馈自己无法完成这个加载请求(它的搜索范围中没有找到所需的类)时，子加载器才会尝试自己去加载。


![类加载流程图.jpg](http://upload-images.jianshu.io/upload_images/2031765-db838e4cec7e6643.jpg?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)


## 实现思路
- 设置命令行参数，指定路径
- 实现类路径：启动类加载器，扩展类加载器，应用程序类加载器
- 将类路径抽象出来，这三个作为其子类

## 编码过程
编码过程从三方面介绍，首先是工程目录，让读者可以清晰看见。其次是运行效果，读者可根据实现的搜索class文件的效果如何，选择是否继续往下观看。源码请移步[传送门](https://github.com/zhaohuXing/jvmgo/tree/master/ch02)
- 工程目录
- 运行效果
- 代码实现

#### 工程目录
- gopath：代码存放区(workspace)
- gopath/src：存放源码的区域
- jvmgo：项目名
- ch02：子文件
- ch02/classpath/：这次主要编写的代码


![project02.png](http://upload-images.jianshu.io/upload_images/2031765-9dc3a16fb1b8d4d6.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)




#### 运行效果
在`ch02/`下编译: `go build -o readClass` 后所在当前目录下便生成可执行文件`readClass`

验证通过启动类或扩展类加载器加载`java.lang.Object`，`-Xjre`是预定义的命令行参数，"/home/sprint/java/utils/jdk1.8.0_91/jre"指定jre路径，效果图如下：
![02_result_01.png](http://upload-images.jianshu.io/upload_images/2031765-6734054eeaea2f16.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

验证通过应用程序类加载器加载`ch02/Test`class(自己编写的类)，如果不指定路径，默认当前路径。`-cp`是预定义的命令行参数，""代表当前目录，效果图如下：

![02_result_02.png](http://upload-images.jianshu.io/upload_images/2031765-d367276856250b0d.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

#### 代码实现
- 在`terminal.go`中预定义`Xjre`命令行参数

```
//定义Termial结构体
type Terminal struct {
	//...省略字段，
	XjreOption  string // 添加XjreOption字段
}

func parseTerminal() *Terminal {
	terminal := &Terminal{}
	//省略部分代码...
	flag.StringVar(&terminal.XjreOption, "Xjre", "", "path to jre")
	flag.Parse()
	//省略部分代码...
	return terminal
}
   ```
- 编写类路径及其子类
总类路径包含:启动类加载器路径，扩展类加载器路径，应用程序类加载器路径。在`ch02/classpath/`下编写`classpath.go`，定义类路径如下:

  ```
type Classpath struct {
	//启动类加载器
	bootClasspath Entry
	//扩展类加载器
	extClasspath Entry
	//应用程序加载器
	userClasspath Entry
}
  ```
每种类加载器都有相同的方法，这时定义一个接口开表示类路径项。在`ch02/classpath/`下编写`entry.go`文件定义`Entry`接口:

  ```
  type Entry interface {
	//负责寻找和加载class
	readClass(className string) ([]byte, Entry, error) //Go函数或方法运行返回多个值

	//类似toString()
	String() string
}
  ```
 `DirEntry`代表:应用程序类加载器，用于加载文件目录中的class。`ZipEntry`代表:应用启动类和扩展类加载器，用于加载jar/zip文件。`CompositeEntry`和`WildcardEntry`代表:Entry的集合。在`ch02/classpath`下创建`entry_dir.go`, `entry_zip.go`,`entry_composite.go`,`entry_wildcard.go`文件实现`Entry`接口。Go的实现接口的方式与Java有所不同，想体验的话，动手实现下，体验下Go的魅力吧！由于考虑篇幅，所以接口具体实现不贴代码了，具体源码请移步[传送门](https://github.com/zhaohuXing/jvmgo/tree/master/ch02)
  
  
- 测试执行
编写完上述代码后，更改`ch02/main.go`中的startJVM()函数来进行测试。


  ```
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

```
一切准备就绪后，在`ch02`文件下敲入`go build -o readClass`，便生成了`readClass`可执行文件。之后按照"运行效果图"中命令行进行执行测试吧！

>小结与声明：一边进一步理解Java虚拟机，一边学习Go。知识主要来源于gitbook(下面有地址)和张秀宏老师的《自己动手编写Java虚拟机》，我的源代码与书上的源码稍微有些不同。

## 参考文献
- 自己动手写Java虚拟机
- https://github.com/astaxie/build-web-application-with-golang
- https://github.com/zhaohuXing/jvmgo
