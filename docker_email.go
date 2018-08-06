package main

import (
	"fmt"
	"log"
	"net/http"
	"net/mail"
	"strings"

	"github.com/gorilla/mux"
	"github.com/peoplenet/go-util/config"
	_ "github.com/scorredoira/email"
)

var router *mux.Router

func regroute() {
	port := config.GetInt("8000")
	router = mux.NewRouter()
	router.HandleFunc("/{email}", mainEmail).Methods("GET")
	address := fmt.Sprintf(":%d", port)
	log.Fatal(http.ListenAndServe(address, router))
}

func mainEmail(w http.ResponseWriter, r *http.Request) {
	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	params := mux.Vars(r)
	email := params["email"]

	m := email.NewMessage("subject", "message")
	m.From = mail.Address{Name: "docker_alert", Address: "docker_test_intern@peoplenetonline.com"}
	m.To = strings.Split(email, ";")
	err := email.Send(
		"aspmx.l.google.com:25",
		nil,
		m)

	if err != nil {
		log.Fatal(err)

	}

}

func main() {
	go regroute()
}
