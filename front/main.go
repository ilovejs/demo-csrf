package main

import (
	"fmt"
	"log"
	"net/http"

	"text/template"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {

		t, err := template.ParseFiles("./index.html")
		if err != nil {
			log.Fatal("template not found", err)
			return
		}

		err = t.Execute(w, nil)
		if err != nil {
			return
		}
	})

	fmt.Println("http://localhost:3000/")

	log.Fatal(http.ListenAndServe(":3000", nil))
}
