package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"time"

	//"io"
	"net/http"
	//"os"
	"github.com/jsatof/qbook/middleware"
)

func index(w http.ResponseWriter, r *http.Request) {
	user, err := middleware.GetUserWithMaxID()
	if err != nil {
		emessage := fmt.Sprintf("Could not fetch the user with the max id\n%s\n", err)
		log.Fatalf(emessage)
		io.WriteString(w, emessage)
	}
	id := user.ID
	id++
	err = middleware.InsertUser(middleware.NewUser(id))
	if err != nil {
		emessage := fmt.Sprintf("Could not insert user with id=%d\n%s\n", id, err)
		log.Fatalf(emessage)
		io.WriteString(w, emessage)
	}

}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	server := http.Server{
		Addr:         ":8000",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	err := middleware.InitDB()
	if err != nil {
		log.Fatalf("Could not connect to the database: %s\n", err)
	}
	defer middleware.Close()
	log.Println("Database connection good")

	log.Println("Starting webserver")
	err = server.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		log.Println("Closing server")
	} else if err != nil {
		log.Fatalf("ERROR: in server\n%s\n", err)
	}
}
