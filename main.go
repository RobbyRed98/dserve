package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// AllFiles gets the content of a directory and all its subdirectories.
// Note: It utilizes thereby tail-recursive calls. (Yeah, not done well, but good enough.)
func AllFiles(directories []string) []string {
	if len(directories) == 0 {
		return []string{}
	}

	head := directories[0]
	fsItem, err := ioutil.ReadDir(head)

	if err != nil || len(fsItem) == 0 {
		return []string{}
	}

	var dirsToProcess []string
	if len(directories) != 0 {
		dirsToProcess = directories[1:]
	}

	var foundFiles []string
	for _, file := range fsItem {
		if file.IsDir() {
			dirsToProcess = append(dirsToProcess, filepath.Join(head, file.Name()))
		} else {
			foundFiles = append(foundFiles, filepath.Join(head, file.Name()))
		}
	}

	return append(foundFiles, AllFiles(dirsToProcess)...)
}

// GetAvailablity writes an HTML availablity message to an Response.
func GetAvailablity(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<p>Dserve is running. (GO)</p>")
}

// FileRequestHandleFunc creates a function which handles the request for the passed file.
func FileRequestHandleFunc(file string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fileContent, _ := ioutil.ReadFile(file)
		fmt.Fprintf(w, string(fileContent))
	}
}

func main() {
	host := flag.String("ip", "localhost", "The host address to run dserve on.")
	port := flag.Int("port", 8080, "The port to run dserve on.")
	path := flag.String("dir", os.Args[0], "The directories path to host.")
	flag.Parse()

	log.Println("Starting traffic light controller service...")
	runtimeDir := filepath.Dir(*path)
	router := mux.NewRouter()
	log.Println(runtimeDir)
	files := AllFiles([]string{runtimeDir})

	for _, file := range files {
		fmt.Println("Added handler for:", file)
		router.HandleFunc("/"+file, FileRequestHandleFunc(file))
	}
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", *host, *port), handlers.CORS()(router)))
}

