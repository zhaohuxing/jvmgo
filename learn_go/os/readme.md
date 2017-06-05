## os包---基本使用(170605) 
- main.go:读取文件`test.txt`,并将数据写到`output.txt`中
- test.txt: 随意的文字 
- main:可执行文件

Go的读写文件想当的简单，了解过Java的同学，很定知道Java读写类一堆一堆的，而且给出庞大的库，个人感觉单纯处理磁盘文件的读写的话，Go是个不错的选择，当然我们可以从Java IO库中学学设计模式嘛，写出漂亮的Go版的代码。

## Go版读写文件
用Go系统自带包实现文件的读写，思路如下:
- 使用`os.Open(filename)`打开文件
- 使用`bufio.NewReader(file)`创建一个`byteReader`
- 使用for循环读取byte，借用`io.EOF`标记完成读取
- 使用`Reader.ReadLine()`来读取一行的bytes,这里注意一行的数据不能太长
- 到此为止，完成读取操作
- 使用`os.Create(filename)`打开文件,如果文件不存在，就创建，如果存在就情况文件里面的数据
- 使用`file.WriteString(value)`将数据写入文件
- 考虑到`os.Create(file)`的，若文件存在需要清空原有的数据，所以，在将读取的数据存储到一个string中，然后一次性写入到输出文件中

详情见源码吧！
