// main.go file
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// func IndexHandler(w http.ResponseWriter,r *http.Request, params.httprouter.Params)

func IndexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
	
}

func main() {
	log.Println("read config")
	fmt.Println("lol")
	router := httprouter.New()
	router.GET("/", IndexHandler)
	

	log.Fatal(http.ListenAndServe(":8080", router))

}
