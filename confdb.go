package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rhinoman/couchdb-go"
)

type TestDocument struct {
	Title string
	Note  string
}

func main() {
	timeout := time.Duration(500 * time.Millisecond)
	conn, err := couchdb.NewConnection("127.0.0.1", 5984, timeout)
	auth := couchdb.BasicAuth{Username: "user", Password: "password"}
	conn.CreateDB("mydatabase", &auth)
	db := conn.SelectDB("mydatabase", &auth)
	theDoc := TestDocument{
		Title: "My Document",
		Note:  "This is a note",
	}

	theId := "MyID" //use whatever method you like to generate a uuid
	//The third argument here would be a revision, if you were updating an existing document

	readDoc := TestDocument{}

	rev, err := db.Read(theId, &readDoc, nil)
	fmt.Println(readDoc.Note)
	rev, err = db.Save(theDoc, theId, rev)
	fmt.Println(rev)
	fmt.Println(err)

	router := mux.NewRouter().StrictSlash(true)
	router.Methods("GET").Path("/").HandlerFunc(Index)
	log.Fatal(http.ListenAndServe(":8080", router))

}

func Index(w http.ResponseWriter, r *http.Request) {
	log.Println("Got request")
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
