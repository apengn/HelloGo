package main

import (
	"net/http"
	"log"
	"io"
	"bytes"
)

const maxUploadSize = 2 * 1024 // 2 mb
const uploadPath = "./tmp"

func main() {

	http.HandleFunc("/getfile/", func(writer http.ResponseWriter, request *http.Request) {
		//v := request.FormValue("values")
		//f, err := os.Create("dome.txt")
		//if err != nil {
		//	writer.WriteHeader(http.StatusInternalServerError)
		//	fmt.Fprint(writer)
		//	return
		//}
		//ioutil.WriteFile(f.Name(), []byte(v), os.ModePerm)
		FileHeader := []byte("xxxxxxxxxxxxxxxxxxxxxx")
		reader := bytes.NewReader(FileHeader)
		//FileContentType := http.DetectContentType(FileHeader)
		writer.Header().Set("Content-Disposition", "attachment; filename="+"xxxxxxx.txt")
		writer.Header().Set("Content-Type", "application/octet-stream;"+"charset=UTF-8")
		io.Copy(writer, reader) //'Copy' the file to the client
		//defer f.Close()
		//go func() {
		//	os.Remove(f.Name())
		//}()
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
