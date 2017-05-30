## 说明
- 代码完毕

##-------------------文档说明

> 后知后觉，慢慢遇到自己喜欢的Go。

## Java运行时数据区
Java虚拟机在执行Java程序的过程中会把它管理的内存化分为若干个不同的数据区域。按照其生命周期来分类，大概分为两种:
- 线程共享区域 ：随着Java虚拟机进程的创建而创建，销毁而销毁。
  - 方法区 ：存放JVM加载类信息，常量，静态变量等(较复杂)
  - 堆 ：几乎所以对象的内存分配区域
- 线程私有区域 ：随着线程的创建而创建，销毁而销毁。
  - 虚拟机栈 ：Java方法执行的内存模型
  - 程序计数器(pc) ：当前线程所执行的字节码的行号指示器

![run-time-data-area.png](http://upload-images.jianshu.io/upload_images/2031765-81bdeed4ba4cbcee.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

> 主要实现线程私有中的数据结构---栈。

## 实现思路
- 编写thread.go，在里面定义一个Thread结构体，包括pc,stack，并编写进栈出栈的方法
```
type Thread struct {
	pc    int `存放字节码执行的行号`
	stack *Stack
}
```
- 编写jvm_stack.go，在里面定义一个Stack结构体，包括size, maxSize和栈顶指针，并编写各种进栈出栈的方法
```
type Stack struct {
	maxSize uint
	size    uint
	_top    *Frame `栈帧`
}
```
- 编写frame.go，在里面定义一个Frame结构体，包括frame指针，局部变量表，操作数栈，并编写Get方法
```
type Frame struct {
	lower        *Frame        `next frame`
	localVars    LocalVars     `局部变量表`
	operandStack *OperandStack `操作数栈`
}
```
- 局部变量表中的变量除了long,double都用一个slot来表示，所以编写slot.go,定义一个Slot结构体，并编写local_vars.go，定义slice的Slot来表示局部变量表，为了存取方便并定义不同类型的变量的方法
```
//solt.go
type Slot struct {
	num int32
	ref *Object
}
//local_vars.go
type LocalVars []Slot
```
- 编写operand_stack.go中定义OperandStack结构体，包括若干Slot，size，为了存取方便并定义不同类型的变量的方法
```
type OperandStack struct {
	size  uint
	slots []Slot
}
```
- 在编写过程中遇到Object结构体，所以编写object.go，创建Object结构体，目前是空的。

注：这部分就像Go的特点一样，简单，所以就不一一叙述了，源码地址[传送门](https://github.com/zhaohuXing/jvmgo/tree/master/ch04)

## 工程目录
工程目录，为了方便清晰的编码，并附上执行后效果。

![ch04.png](http://upload-images.jianshu.io/upload_images/2031765-44bc59f3996781db.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

- gopath：代码存放区(workspace)
- gopath/src：存放源码的区域
- jvmgo：项目名
- ch04：子文件
- ch04/rtda/：这次主要编写的代码

搭建运行时数据区，在`ch04`目录下编码，需要新建`rtda`(run-time-data-area)文件夹, `实现思路`中创建的文件都是在`rtda`目录下完成的。如上文件编码完毕后，需要更改main.go中的代码，具体请查看源码。所以编码结束后，需要在`ch04/`目录下执行`go build -o run-time-data-area`，便在该目录下生成了`run-time-data-area`可执行文件，之后执行`run-time-data-area test`测试就ok了。

![04_result_01.png](http://upload-images.jianshu.io/upload_images/2031765-0d5377d56dce8b36.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

## 参考文献
- 自己动手写Java虚拟机
- [https://github.com/zhaohuXing/jvmgo](https://github.com/zhaohuXing/jvmgo)


精彩文章，持续更新，请关注微信公众号：
![帅哥美女扫一扫](http://upload-images.jianshu.io/upload_images/2031765-2c4654abe66cd4c8.jpg?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)
