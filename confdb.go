package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/rhinoman/couchdb-go"
)

type TestDocument struct {
	Title string
	Note  string
}

func prepareDb() {
	timeout := time.Duration(500 * time.Millisecond)
	couchPort, _ := strconv.Atoi(os.Getenv("COUCHDB_PORT"))
	conn, err := couchdb.NewConnection(os.Getenv("COUCHDB_HOST"), couchPort, timeout)
	auth := couchdb.BasicAuth{Username: os.Getenv("COUCHDB_USER"), Password: os.Getenv("COUCHDB_PASSWORD")}
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
}

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.Methods("GET").Path("/").HandlerFunc(Index)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))

}

func Index(w http.ResponseWriter, r *http.Request) {
	log.Println("Got request")
	prepareDb()
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
