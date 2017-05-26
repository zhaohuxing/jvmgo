package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

//代表zip 或jar
type ZipEntry struct {
	absPath string
}

func newZipEntry(path string) *ZipEntry {
	//获取绝对路径，如果发生异常，则会赋值给err
	absPath, err := filepath.Abs(path)
	//nil代表无值
	if err != nil {
		panic(err) //类似抛出异常
	}
	return &ZipEntry{absPath}
}

//defer:延迟语句，当函数执行到最后时,这些defer语句会按照逆序执行
func (self *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	r, err := zip.OpenReader(self.absPath)
	if err != nil {
		return nil, nil, err
	}
	defer r.Close()
	for _, f := range r.File {
		if f.Name == className {
			//试图打开className文件
			rc, err := f.Open()
			if err != nil {
				return nil, nil, err
			}
			defer rc.Close()

			//可以打开文件，读取数据
			data, err := ioutil.ReadAll(rc)
			if err != nil {
				return nil, nil, err
			}
			return data, self, nil
		}
	}
	return nil, nil, errors.New("class not found: " + className)
}

func (self *ZipEntry) String() string {
	return self.absPath
}
