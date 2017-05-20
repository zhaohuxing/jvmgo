package classpath

import (
	"errors"
	"strings"
)

//由dir, zip组成
type CompositeEntry []Entry

func newCompositeEntry(pathList string) CompositeEntry {
	//启动类路径
	//扩展类路径
	//用户类路径
	compositeEntry := []Entry{}
	for _, path := range strings.Split(pathList, pathListSeparator) {
		entry := newEntry(path)
		compositeEntry = append(compositeEntry, entry)
	}
	return compositeEntry
}

func (self CompositeEntry) readClass(className string) ([]byte, Entry, error) {
	for _, entry := range self {
		data, from, err := entry.readClass(className)
		if err == nil {
			return data, from, nil
		}
	}
	return nil, nil, errors.New("class not found:" + className)
}

//make 用于内建类型(map, slice, channel)的内存分配
//new 用于各种类型的内存分配
func (self CompositeEntry) String() string {
	//分配self大小的string的内存给strs
	strs := make([]string, len(self))

	for i, entry := range self {
		strs[i] = entry.String()
	}
	return strings.Join(strs, pathListSeparator)
}
