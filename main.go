package main

import ("fmt"
	"log"
	"net/http"
)
func helloHandler(w http.ResponseWriter, r *http.Request){
if r.URL.Path != "/hello" {
	http.Error(w, "404 not found", http.StatusNotFound)
	return
}
if r.Method != "GET"{
	http.Error(w, "method is not supported". http.StatusNotFound)
	return
}
}
func main (){
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)
	http.HandeFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)
	
	fmt.Printf("Starting Server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
