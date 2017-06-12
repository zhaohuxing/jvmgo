package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

/*其实wildcardEntry 也是一个CompositeEntry*/
func newWildcardEntry(path string) CompositeEntry {
	baseDir := path[:len(path)-1] // remove *
	compositeEntry := []Entry{}
	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && path != baseDir {
			return filepath.SkipDir
		}

		if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") {
			jarEntry := newZipEntry(path)
			compositeEntry = append(compositeEntry, jarEntry)
		}
		return nil
		//3428655361
	}
	//walkFn也是一参数,即函数作为参数
	filepath.Walk(baseDir, walkFn)
	return compositeEntry
}
