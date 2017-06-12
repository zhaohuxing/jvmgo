package classpath

import (
	"os"
	"strings"
)

//存放路径分隔符
const pathListSeparator = string(os.PathListSeparator)

type Entry interface {
	//负责寻找和加载class
	readClass(className string) ([]byte, Entry, error) //Go函数或方法运行返回多个值

	//类似toString()
	String() string
}

func newEntry(path string) Entry {
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}
	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}
	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") ||
		strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {
		return newZipEntry(path)
	}
	return newDirEntry(path)
}

// DirEntry : 目录形式的类路径
//
