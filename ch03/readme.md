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

## ===================部分文档

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


