package classpath

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

//目录形式的类路径
type DirEntry struct {
	//存放目录的绝对路径:
	absDir string
}

/*把参数转换为绝对路径*/
func newDirEntry(path string) *DirEntry {
	//获取绝对路径
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err) // 若出现错误，调用panic()函数终止程序执行
	}
	return &DirEntry{absDir}
}

func (self *DirEntry) readClass(className string) ([]byte, Entry, error) {
	//将目录和class文件拼成一个完整的路径
	fileName := filepath.Join(self.absDir, className)
	//打印文件名
	fmt.Printf("fileName: %s", fileName)

	//通过ReadFile函数读取class文件内容
	data, err := ioutil.ReadFile(fileName)

	//返回内容
	return data, self, err
}

func (self *DirEntry) String() string {
	return self.absDir
}
