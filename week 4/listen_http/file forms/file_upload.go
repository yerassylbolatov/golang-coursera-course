package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

var uploadFormTmpl = []byte(
	`<html>
		<body>
		<form action="/upload" method="post"
		enctype="multipart/form-data">
			Image: <input type="file" name="my_file">
			<input type="submit" name"Upload">
		</form>
		</body>
	</html>`)

func mainPageTmpl(w http.ResponseWriter, r *http.Request) {
	w.Write(uploadFormTmpl)
}

func uploadPage(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(5 * 1024 * 1025)
	file, handler, err := r.FormFile("my_file")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fmt.Fprintf(w, "handler.Filename %v\n", handler.Filename)
	fmt.Fprintf(w, "handler.Header %#v\n", handler.Header)

	hasher := md5.New()
	io.Copy(hasher, file)
	fmt.Fprintf(w, "md5 %x\n", hasher.Sum(nil))
}

func uploadRawData(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	p := &Params{}
	err = json.Unmarshal(body, &p)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Fprintf(w, "content-type %#v\n", r.Header.Get("Content-Type"))
	fmt.Fprintf(w, "params %#v\n", p)
}

//curl -v -X POST -H "Content-Type: application/json" -d '{"id": 1, "user": "yerassyl"}' http://localhost:8080/raw_body

type Params struct {
	ID   int
	User string
}

func main() {
	http.HandleFunc("/", mainPageTmpl)
	http.HandleFunc("/upload", uploadPage)
	http.HandleFunc("/raw_body", uploadRawData)
	fmt.Println("server starts at :8080")
	http.ListenAndServe(":8080", nil)
}
