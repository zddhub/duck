package duck

import (
	"log"
	"net/http"
)

func Static(directory string) Handler {
	return func(res http.ResponseWriter, req *http.Request, log *log.Logger) {
		if req.Method != "GET" && req.Method != "HEAD" {
			return
		}

		dir := http.Dir(directory)

		filePath := req.URL.Path
		if req.URL.Path != "" && filePath[len(filePath)-1] == '/' {
			filePath += "index.html"
		}

		f, err := dir.Open(filePath)
		if err != nil {
			log.Println("Static", err)
			return
		}
		defer f.Close()

		fi, err := f.Stat()
		if err != nil || fi.IsDir() {
			return
		}

		log.Println("[Static] Serving", directory, req.URL.Path)

		http.ServeContent(res, req, directory+filePath, fi.ModTime(), f)
	}
}
