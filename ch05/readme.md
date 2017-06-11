## 编码日志:
- 20170531: 存储指令，读取指令，栈指令

## ---------------------------------------
编码过程及目录详解：
- instructions/base/instructions.go中定义指令接口，并定义了"抽象指令":NoOperandsInstruction, Index8Instruction, Index16Instruction, BranchInstruction
- instructions/base/bytecode_reader.go中定一个BytecodeReader结构体，并实现了各种读取的方法
- instructions/constants/nop.go中定义了一个毫无操作的一个结构体和方法
- instructions/constants/const.go中定义了一系列的const指令，并实现实现指令接口
- instructions/constants/ipush.go中定义BIPUSH,SIPUSH方法, 并实现指令接口
- instructions/loads/aload.go中定义各种引用类型的结构体，并实现了指令接口
- instructions/loads/dload.go中定义各种double类型的结构体，并实现了指令接口
- instructions/loads/fload.go中定义各种float类型的结构体，并实现了指令接口
- instructions/loads/iload.go中定义各种int类型的结构体，并实现了指令接口
- instructions/loads/lload.go中定义各种long类型的结构体, 并实现了指令接口
- instructions/stores/astore.go中定义各种引用类型的结构体，并实现了指令接口
- instructions/stores/dstore.go中定义各种double类型的结构体，并实现了指令接口
- instructions/stores/fstore.go中定义各种float类型的结构体，并实现了指令接口
- instructions/stores/istore.go中定义各种int类型的结构体，并实现了指令接口
- instructions/stores/lstore.go中定义各种long类型的结构体，并实现了指令接口
- rtda/operand_stack.go中添加PushSlot()和PopSlot()方法
- instructions/stack/pop.go中定义POP,POP2结构体，分别用于弹出一个slot的变量，两个slot的变量,并实现了指令接口
- instructions/stack/dup.go中定义dupx系列，用于复制栈顶的变量,并实现了指令接口
- instructions/stack/swap.go中定义SWAP结构体,用于栈顶变量的交换, 并实现了指令接口
- instructions/math/下的全部文件，对应着各种基本类型的加减乘除，且，或，异等运算
- instructions/coversions/下全部文件，对应各种类型的转换,并实现了指令接口
- instructions/compairsons/下全部文件，对应着各种类型的比较，并实现了指令接口
- instructions/control/下全部文件，用于控制，并实现了指令接口
- instructions/extended/下全部文件，扩展指令
- instructions/factory.go，用于生产指令的工厂
- interpreter.go，解释器的实现
##----------------------------------------------------
更新内容:
- instructions/文件下所有的指令
	- base: 公共指令
	- comparisons: 比较指令
	- constants: 常量指令
	- control: 控制指令
	- conversions: 转换指令
	- extends: 扩展指令
	- loads: 加载指令
	- math: 数学指令
	- stack: 操作数栈指令
	- stores: 加载指令
	- factory.go: 生产指令的工厂

## ================= 文档 ======================
> 说在前面，代码为主，说明为辅。算是Go编写JVM指令集和解释器的代码说明文档吧！

## 指令集
>Java虚拟机的指令由一个字节长度的、代表着特定操作含义的**操作码(opcode)**以及跟随其后的零至多个代表此操作所需的**操作数(Operands)**所构成。


用Go将Java虚拟机指令抽象成接口，这样解码和执行写在对应的方法上，这样使用起来很灵活。伪代码如下：
```
type Instruction interface {
	//从字节码中提取操作数
	FetchOperands(reader *BytecodeReader)
	//执行指令逻辑
	Execute(frame *rtda.Frame)
}
```
因为用一个字节表示**操作码**，所以指令个数少于256个，截止JDK7现有200多条指令。每条指令都有助记符来帮助记忆指令的作用，举个栗子，指令iload_0表示:将int型0从局部变量表加载到操作数栈。Go代码实现:
```
type ILOAD_0 struct {
	base.NoOperandsInstruction //Java中的"继承"
}
// 因base.NoOprendsInstruction中实现了FetchOperands(reader)接口，所以不需要重复实现了
func (self *ILOAD_0) Execute(frame *rtda.Frame) {
	_iload(frame, 0)
}
```
其他指令实现格式类似，由于代码量繁多，就不一一描述啦，想具体了解，请参考[源码](https://github.com/zhaohuXing/jvmgo/tree/master/ch05)

## 解释器
顾名思义，这里说的解释器，当然是解释字节码的。采用《Java虚拟机规范》中一段伪代码来形象的说明下：
```
do {
	自动计算PC寄存器以及从PC寄存器的位置取出操作码;
	if (存在操作数) 取出操作数;
	执行操作码所定义的操作
} while (处理下一次循环);

```
从这段伪代码上看来，解释器需要指令集来完成如上循环中的一系列操作。
- pc寄存器： 线程私有，当前线程执行的字节码的行号指示器
- 操作码：操作码其实就是指令序列号，用来告诉JVM要执行哪一条指令
- 操作数：操作码所定义的操作，需要的参数

用Go表示解释器的工作:
```
func loop(thread *rtda.Thread, bytecode []byte) {
	frame := thread.PopFrame()
	reader := &base.BytecodeReader{}
	for {
		pc := frame.NextPC()
		thread.SetPC(pc)

		//decode
		reader.Reset(bytecode, pc)
		opcode := reader.ReadUint8()
		inst := instructions.NewInstruction(opcode)
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.PC())
		//execute
		fmt.Printf("pc:%2d inst:%T %v\n", pc, inst, inst)
		inst.Execute(frame)
	}
}

```

## 编码过程
编码过程从三方面介绍，首先是工程目录，让读者可以清晰看见。其次是运行效果，读者可根据实现的解释器的效果如何，选择是否继续往下观看。源码请移步[传送门](https://github.com/zhaohuXing/jvmgo/tree/master/ch05)

- 工程目录
- 运行效果
- 代码实现

#### 工程目录
- gopath：代码存放区(workspace)
- gopath/src：存放源码的区域
- jvmgo：项目名
- ch05：子文件
- ch05/instructions/, ch05/interpreter.go：这次主要编写的代码


![ch05.png](http://upload-images.jianshu.io/upload_images/2031765-ef7a2796eb26c0b8.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

#### 运行效果
在`ch05`下运行`go build -o parse`，便生成了一个可执行文件`parse`


验证通过应用程序类加载器加载`ch05/GaussTest` class(自己编写的类，从1到100的累加和)，如果不指定路径，默认当前路径。-cp是预定义的命令行参数，""代表当前目录，效果图如下：

![05_result_01.png](http://upload-images.jianshu.io/upload_images/2031765-edf6213f7bdc6128.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)


![05_result_02.png](http://upload-images.jianshu.io/upload_images/2031765-448b2ab4a08e991c.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

#### 代码实现
主要分为两部分，指令集和解释器:
- 指令集 ：代码量大而简单，主要是各种指令的编写
- 解释器： 仅能执行一个Java方法，相对来说代码量要少的多

关于指令集代码的编写参考随手记录的[编码文档](https://github.com/zhaohuXing/jvmgo/blob/master/ch05/readme.md)

这个解释器有点简陋，只能解释一个方法，还是特么是main方法，哈哈！

解释器编码思路:
- 在指令集差不多实现完的情况下
- 修改入口程序main.go中`startJVM`方法，在获取完并解析完class后，获取main方法
- 编写`interpret`方法对main方法进行解释执行

这个不手敲一遍，理解的效果不太好，所以附上源码地址[传送门](https://github.com/zhaohuXing/jvmgo/tree/master/ch05)

## 参考文献
- 《自己动手写Java虚拟机》
- 《Java虚拟机规范》(jdk7)
- [https://github.com/zhaohuXing/jvmgo](https://github.com/zhaohuXing/jvmgo)

精彩文章，持续更新，请关注微信公众号：
![帅哥美女扫一扫](http://upload-images.jianshu.io/upload_images/2031765-2c4654abe66cd4c8.jpg?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)
