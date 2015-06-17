package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	file, err := os.Open("public/html/index.html")
	defer file.Close()
	if err != nil {
		fmt.Fprint(w, "Error: %v\n", err)
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Fprint(w, "Error: %v\n", err)
	}
	fmt.Fprint(w, string(data))
}

func main() {
	router := httprouter.New()
	router.ServeFiles("/js/*filepath", http.Dir("public/js"))
	router.ServeFiles("/css/*filepath", http.Dir("public/css"))
	router.GET("/", Index)

	log.Fatal(http.ListenAndServe(":8080", router))
}
