package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/go-xorm/xorm"
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

func UserGetHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

func UserPostHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

var (
	Engine *xorm.Engine
)

func main() {
	router := httprouter.New()
	router.ServeFiles("/js/*filepath", http.Dir("public/js"))
	router.ServeFiles("/css/*filepath", http.Dir("public/css"))
	router.ServeFiles("/img/*filepath", http.Dir("public/img"))
	router.GET("/", Index)
	router.GET("/user", UserGetHandler)
	router.POST("/user", UserPostHandler)

	// var err error
	// Engine, err = xorm.NewEngine("mysql", "test:1234@/naiping?charset=utf8")
	// if err != nil {
	// 	log.Fatal("database err:", err)
	// }
	// Engine.ShowDebug = true
	// Engine.ShowErr = true
	// Engine.ShowWarn = true
	// Engine.ShowSQL = true
	// f, err := os.Create("sql.log")
	// if err != nil {
	// 	log.Fatal("sql log create failed", err)
	// }
	// defer f.Close()
	// Engine.Logger = xorm.NewSimpleLogger(f)

	// err = Engine.Sync2(new(User), new(UserItem), new(Item))
	// if err != nil {
	// 	log.Fatal("database sync failed", err)
	// }

	log.Fatal(http.ListenAndServe(":8080", router))
}
