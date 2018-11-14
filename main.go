package main

import (
	"encoding/json"
	"fmt"
	"io"
	//"io/ioutil"
	"log"
	"net/http"
	"os"
)

var Info *log.Logger = log.New(os.Stdout,
	"INFO: ",
	log.Ldate|log.Ltime)

// The data model to be sent
type Response struct {
	Message string
	Code    int
}

// hello world, the web server
func Home(w http.ResponseWriter, req *http.Request) {

	if req.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch req.Method {
	case "GET":
		data := []Response{Response{"Hello GET", 200}}
		b, _ := json.Marshal(data)

		Info.Println("Sending response", string(b))
		io.WriteString(w, string(b))
	case "POST":
		fmt.Fprintf(os.Stdout, "Post from website! req.PostFrom = %v\n", req.PostForm)
		Info.Println("Post from website from", req.PostForm)
		data := []Response{Response{"Hello POST", 200}}
		b, _ := json.Marshal(data)

		Info.Println("Sending response", string(b))
		io.WriteString(w, string(b))
	default:
		Info.Println("Sorry, only GET and POST methods are supported.")
		io.WriteString(w, "Sorry, only GET and POST methods are supported.")
	}

}

func main() {
	http.HandleFunc("/", Home)

	Info.Println("Listening at 5555")

	log.Fatalf("APP: %d %s ", log.Ldate|log.Ltime, http.ListenAndServe(":5555", nil))
}
