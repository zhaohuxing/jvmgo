package handler

import (
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"runtime/debug"
	"strings"
)

const (
	TEMPLATE_DIR = "./views"
	UPLOAD_DIR   = "./uploads"
)

var templates = make(map[string]*template.Template)

//将模板信息预先加载到内存
func init() {
	// Get TEMPLATE_DIR's info
	fileInfoArr, err := ioutil.ReadDir(TEMPLATE_DIR)
	if err != nil {
		panic(err)
	}

	//define template's name and path
	var templateName, templatePath string

	for _, fileInfo := range fileInfoArr {
		templateName = fileInfo.Name()
		if ext := path.Ext(templateName); ext != ".html" {
			continue
		}

		templatePath = TEMPLATE_DIR + "/" + templateName
		log.Println("Loding template:", templatePath)
		t := template.Must(template.ParseFiles(templatePath))
		templates[templateName] = t
		log.Println(templateName)
	}

}

func PhotoHandler() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", safeHandler(listHandler))
	mux.HandleFunc("/upload", safeHandler(uploadHandler))
	mux.HandleFunc("/view", safeHandler(viewHandler))
	return mux
}

//这里函数也能作为参数传入
func safeHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if e, ok := recover().(error); ok {
				http.Error(w, e.Error(), http.StatusInternalServerError)
				log.Printf("WARN: panic in %v. - %v.", fn, e)
				log.Println(string(debug.Stack()))
			}
		}()
		fn(w, r)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	imageId := r.FormValue("id")
	imagePath := UPLOAD_DIR + "/" + imageId
	if exists := isExists(imagePath); !exists {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "image")
	http.ServeFile(w, r, imagePath)
}

func isExists(imagePath string) bool {
	_, err := os.Stat(imagePath)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		renderHtml(w, "upload", nil)
	}

	if r.Method == "POST" {
		f, h, err := r.FormFile("image")
		if err != nil {
			panic(err)
		}
		filename := h.Filename
		defer f.Close()

		t, err := ioutil.TempFile(UPLOAD_DIR, filename)
		if err != nil {
			panic(err)
		}
		defer t.Close()

		_, err = io.Copy(t, f)
		if err != nil {
			panic(err)
		}
		name := strings.Replace(t.Name(), "uploads/", "", 1)
		http.Redirect(w, r, "/view?id="+name, http.StatusFound)
	}
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	fileInfoArr, err := ioutil.ReadDir(UPLOAD_DIR)
	if err != nil {
		panic(err)
	}

	locals := make(map[string]interface{})
	images := []string{}
	for _, fileInfo := range fileInfoArr {
		images = append(images, fileInfo.Name())
	}
	locals["images"] = images
	renderHtml(w, "list", locals)
}

func renderHtml(w http.ResponseWriter, tmpName string, locals map[string]interface{}) {
	tmpName += ".html"
	err := templates[tmpName].Execute(w, locals)
	if err != nil {
		panic(err)
	}
}
