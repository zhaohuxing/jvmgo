package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"path"
)

const (
	TEMPLATE_DIR = "./views"
)

var templates = make(map[string]*template.Template)

//使用init()函数将所有的模板加载到程序中
func init() {
	//读取view文件下的信息
	fileInfoArr, err := ioutil.ReadDir(TEMPLATE_DIR)
	if err != nil {
		panic(err)
	}

	var templateName, templatePath string
	for _, fileInfo := range fileInfoArr {
		templateName = fileInfo.Name()
		//根据文件名获取扩展类型为html文件,
		if ext := path.Ext(templateName); ext != ".html" {
			continue
		}

		templatePath = TEMPLATE_DIR + "/" + templateName
		log.Println("Loding template:", templatePath)
		t := template.Must(template.ParseFiles(templatePath))
		//考虑到只有两个html文件,将templateName做为索引
		templates[templateName] = t
	}
}

func main() {
	//NewServeMux()什么意思?
	mux := http.NewServeMux()
	mux.HandleFunc("/", safeHandler(listHandler))
	//	mux.HandleFunc("/view", safeHandler(viewHandler))     // TODO
	//	mux.HandleFunc("/upload", safeHandler(uploadHandler)) //TODO
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
