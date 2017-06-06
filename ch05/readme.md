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
