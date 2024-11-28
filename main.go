package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("listening on localhost:8080")

	err := http.ListenAndServe("localhost:8080", http.HandlerFunc(handler))
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func handler(res http.ResponseWriter, req *http.Request) {
	fmt.Println("root handler")

	_, err := res.Write([]byte(`{"status":"ok", "code":200}`))
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	res.WriteHeader(http.StatusOK)
}
