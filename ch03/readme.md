##  文件说明:
- class_reader.go中定义了ClassReader结构体,可以读取(u1, u2, u4, u8)类型。
- class_file.go中定义了ClassFile结构体,并存在一系列的函数和方法
- member_info.go中定义了MemeberInfo统一表示字段和方法
- constant_pool.go中定义了slice的ConstantInfo
- constant_info.go中定义ConstantInfo(接口)结构体,并声明了14中常量的tag值,方便创建对应的常量
	- cp_utf8_info.go定义了ConstantUtf8Info实现了ConstantInfo接口
	- cp_numberic_info.go定义了ConstantIntegerInfo, ConstantFloatInfo, ConstantLongInfo, ConstantDoubleInfo都实现了ConstantInfo接口
	- cp_class_info.go定义了ConstantClassInfo实现了ConstantInfo接口
	- cp_string_info.go定义了ConstantStringInfo实现了ConstantInfo接口
	- cp_member_ref_info.go定义了ConstantMemberrefInfo, 而ConstantFieldrefInfo, ConstantMethodrefInfo, ConstantInterfaceMethodrefInfo “继承”了ConstantMemberrefInfo
	- cp_name_and_type_info.go定义了ConstantNameAndTypeInfo实现了ConstantInfo接口
	- cp_invoke_dynamic_info.go定义了ConstantMethodHandleInfo,ConstantMethodTypeInfo, ConstantInvokeDynamicInfo都实现了ConstantInfo 
- attribute_info.go中定义了AttributeInfo接口:
	- attr_markers.go中定义了MarkerAttribute并实现了AttributeInfo,DeprecatedAttribute和SyntheticAttribute“继承”了MarkerAttribute, 结构体中内容为空（摆设）

	- attr_source_file.go中定义了SourceFileAttribute结构体并实现了AttributeInfo, 该属性只出现在ClassFile结构中
	- attr_constant_value.go中定义了ConstantValueAttribute结构体并实现了AttributeInfo.该属性只出现在Field_info结构中
	- attr_code.go中定义了CodeAttribute结构体并实现了AttributeInfo，该属性只出现在method_info中
	- attr_exceptions.go中定义了ExceptionsAttribute结构体并实现了AttributeInfo
	- attr_line_number_table.go中定义了LineNuberTableAttribute结构体并实现了AttributeInfo（不是运行期所必须的）
	- attr_local_variable_table.go中定义了LocalVariableTableAttribute结构体实现了AttributeInfo(不是运行期所必须的)
	- attr_unparsed.go中定义了未解析的结构体

## ========================== 用Go编写JVM之解析Class文件

，被Go的魅力所折服，慢慢将发生下面的故事

## 类文件结构
一个class文件由那些部分构成的呢？
>  Class文件是一组以8位字节为基础单位的二进制流，各个数据项目严格按照顺序紧凑地排列在Class文件之中，中间没有添加任何分隔符。

(这使得整个class文件中存储的内容几乎全部是程序运行的必要数据，没有空隙存在。当遇到需要占用8位字节以上空间数据项时，则会按照高位在前的方式分割成若干个8位字节进行存储)

注意: 任意一个class文件，对应着唯一一个类或接口的定义信息，但反过来，类或接口并不一定都得定义在文件里。（可以通过类加载器生成）。主要说明一点就是:它不是以磁盘文件的形式存在。

Class文件结构中只要两种数据类型:无符号数和表
- 无符号数:基本的数据类型，以u1, u2, u3, u4, u8来分别代表1个字节，2个字节，4个字节和8个字节的无符号数，可以用来描述:数字、索引引用、数量值或按照UTF-8编码构成字符串值
- 表: 由多个无符号数或者其他表作为数据项构成的复合数据类型

整个class文件本质上就是一张表,所有的表都习惯性地以"_info"结尾。

## Class文件格式
|  类型 | 名称  | 数量  |
| :-----: | :-----: | :-----: |
|  u4  |  magic  |  1  | 
|  u2  | minor_version | 1 |
|  u2  | major_version  | 1 |
|  u2  |  constant_pool_count  |  1  |
|  cp_info  | constant_pool  |  constant_pool_count - 1  |
|  u2  |  access_flags  |  1  |
|  u2  |  this_class  |  1  |
|  u2  |  super_class  |  1  |
|  u2  |  interfaces_count  |  1  |
|  u2  |  interfaces  |  interfaces  |
|  u2  |  fields_count  |  1  |
|  field_info  |  fields  |  fields_count  |
|  u2  |  methods_count  |  1  |
|  methods_info  |  methods  |  methods_count  |
|  u2  |  attributes_count  |  1  |
|  attribute_info  |  attributes  |  attributes_count  |

这个表格更加体现出各个数据项目严格按照顺序紧凑地排列在Class文件之中。

## 简略描述下class中各个数据项的含义
- magic：判断是否是class文件，位于文件开头，固定字符值:`0xCACFBABE`
- version ：minor_version次版本号，major_version主版本号，版本号由`major_version.minor_version`组成
- constant_pool_count ：常量池的大小，这里注意下标从1开始
- constant_pool ：常量池内容(截止JDK1.7存在14个常量结构)
- access_flags  ：类或接口层次的访问信息

见名知义，剩下就不简单的描述了。constant_pool中索引值0是为了标识该类中没有任何常量池中引用所设置的，所以下标从1开始。

## 实现思路
上一篇《用Go编写JVM之搜索class文件》完成搜索class文件及读取class文件的二进制流，在此基础上将二进制流解析成class结构中的各个数据项。
- 通过classpath.ReadClass(className)，获取class文件的二进制数据流
- 通过classfile.Parse(classData)，将二进制数据流解析成各个数据项的数据

![flow.png](http://upload-images.jianshu.io/upload_images/2031765-e9026a51fc99b222.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

核心代码:
```

type ClassFile struct {
	magic        uint32
	minorVersion uint16
	majorVersion uint16
	constantPool ConstantPool
	accessFlags  uint16
	thisClass    uint16
	superClass   uint16
	interfaces   []uint16
	fields       []*MemberInfo
	methods      []*MemberInfo
	attributes   []AttributeInfo
}
//classfile.go中核心方法
func (self *ClassFile) read(reader *ClassReader) {
	//读取魔数
	self.readAndCheckMagic(reader)
	//读取版本号
	self.readAndCheckVersion(reader)
	//读取常量池
	self.constantPool = readConstantPool(reader)
	//读取accessFlags
	self.accessFlags = reader.readUint16()
	self.thisClass = reader.readUint16()
	self.superClass = reader.readUint16()
	self.interfaces = reader.readUint16s()
	self.fields = readMembers(reader, self.constantPool)
	self.methods = readMembers(reader, self.constantPool)
	self.attributes = readAttributes(reader, self.constantPool)
}

```
 因class文件内部是按照严格顺序摆放，所以依次读取就OK,这里创建了一个`ClassReader`来包装二进制流成各种格式的数据。这里除了常量池和属性表比较复杂外，其余的都是按序读取对应格式的数据。

## 编码过程
编码过程从三方面介绍，首先是工程目录，让读者可以清晰看见。其次是运行效果，读者可根据实现的解析class文件的效果如何，选择是否继续往下观看。源码请移步[传送门](https://github.com/zhaohuXing/jvmgo/tree/master/ch03)

- 工程目录
- 运行效果
- 代码实现

#### 工程目录
- gopath：代码存放区(workspace)
- gopath/src：存放源码的区域
- jvmgo：项目名
- ch03：子文件
- ch03/classfile/：这次主要编写的代码


![ch03.png](http://upload-images.jianshu.io/upload_images/2031765-6621640e981188c8.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

#### 运行效果
在`ch03`目录下编译`go build -o parseClass`后所在目录就生成了`parseClass`可执行文件

验证通过启动类或扩展类加载器加载`java.lang.Object`，-Xjre是预定义的命令行参数，"/home/sprint/java/utils/jdk1.8.0_91/jre"指定jre路径，效果图如下：

![03_result_01.png](http://upload-images.jianshu.io/upload_images/2031765-8d0d731f13a6c2fe.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)


验证通过应用程序类加载器加载`ch03/Test` class(自己编写的类)，如果不指定路径，默认当前路径。-cp是预定义的命令行参数，""代表当前目录，效果图如下：

![03_result_02.png](http://upload-images.jianshu.io/upload_images/2031765-adc9db6d9e83dff3.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

#### 代码实现
由于代码较多，只贴些部分代码，详情请看源码[传送门](https://github.com/zhaohuXing/jvmgo/tree/master/ch03)
- 定义一个ClassReader结构体，并编写解决二进制向u1, u2, u4, u8转换的一系列方法，在方法内部实现数据下标的移动(源码：classfile/class_reader.go)

```
/*
	定义ClassReader结构体，里面存放bytes
	定义各种bytes向uint8, uint16, uint32, uint64, []uint16转换的方法
*/
type ClassReader struct {
	data []byte
}

//u1 无符号1字节
func (self *ClassReader) readUint8() uint8 {
	//读取u1类型数据
	val := self.data[0]
	self.data = self.data[1:] //读取后指针后移1byte
	return val

}
//省略其他类型
```
- 定义一个ClassFile结构体来表示Class文件结构(源码：classfile/class_file.go)
- Class文件结构中一些可以用`ClassReader`中的方法可以搞定：如magic(源码：classfile/class_file.go)
- 定义一个MemberInfo结构体处理字段和方法读取(源码：classfile/member_info.go)
- 定义一个ConstantPool来处理常量池，有关各个文件之间关系图如下:

![constant_pool.png](http://upload-images.jianshu.io/upload_images/2031765-a0ee6a64418037fe.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

- 定义一个AttributeInfo来处理属性表,关系图如下:

![attribute_info.png](http://upload-images.jianshu.io/upload_images/2031765-c8093650decbf803.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

完成这一系列操作，那么可以解析class二进制流啦。

#### 测试
重写`main.go`中的`startJVM()`方法，并加载class（/ch03/main.go）。完成这一系列操作后，在`/ch03/`编译`go build -o parseClass`，验证就ok了！

## 参考文献
- 自己动手写Java虚拟机
- https://github.com/zhaohuXing/jvmgo
