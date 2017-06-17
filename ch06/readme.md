

# 第六章 类和对象
- 方法区
  - 各种类型信息
  - 运行时常量池
- 类加载器
- 对象、实例变量和类变量
- 类和字段符号引用解析
- 类和对象相关指令
- 测试运行

> 编码日志，记录着编码过程

## 准备工作
- 复制`ch05`并改名为`ch06`,清理`ch05`中的测试文件
- 创建`/ch06/rtda/heap/`文件夹
- 将`/ch06/rtda/`下`Object.go`移动到`/ch06/rtda/heap/`下，`Object.go`里面的内容见源码
- 修改`/ch06/rtda/`下`slot.go`, `local_vars.go`,`operand_stack.go`文件，添加import, 并将`*Object`更改成`*heap.Object`

完成上述操作后，代码中依赖struct `Class`, `Slots`,标记下

## 方法区

-各种类型信息

