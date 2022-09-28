package main

import (
	"fmt"
	"log"
	"net/http"
)

// handler : it is a function which contains chunk of code which supposed to make every request made by the route
// Route is just a path in the URL
//But what is this http.Handler?
// It is a interface anything that implements this interface is a handler
// type Handler interface{
//        ServeHTTP(ResponseWriter,* request)
//}
//The SERVER MULTIPLEXER:
//  Whenever we create a request from the client to the server it goes to multiplexer then it matches the specific route to the handlers
// mux:=http.NewServeMux()
type a int

func (b a) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("<h1>hello Darling</h1>"))
}
func main() {
	var b a
	fmt.Println("hello the port is running yaar in 8080")
	http.Handle("/", b)
	//http.HandleFunc("/durga", func(w http.ResponseWriter, r *http.Request) {
	//	w.WriteHeader(200)
	//	w.Write([]byte("<h1>Bye Darling</h1>"))
	//})
	//

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln(err)
	}

}
