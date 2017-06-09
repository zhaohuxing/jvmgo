package handler

import (
	"log"
	"net/http"
	"runtime/debug"
)

func ListHandler(w http.ResponseWriter, r *http.Request) {
	fileInfoAttr, err := ioutil.ReadDir("../uploads") //TODO
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

func renderHtml(w http.ResponseWriter, tmpl string, locals map[string]interface{}) {
	err := templates[tmpl].Execute(w, locals)
	if err != nil {
		panic(err)
	}
}
func SafeHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if e, ok := recover().(error); ok {
				http.Error(w, e.Error(), http.StatusInternalServerError)
				log.Printf("WARN: panic in %v. - %v", fn, e)
				log.Println(string(debug.Stack()))
			}
		}()
		fn(w, r)
	}
}
