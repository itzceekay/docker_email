package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/mail"
	"strings"

	"github.com/gorilla/mux"
	"github.com/scorredoira/email"
)

var router *mux.Router

func regroute() {
	port := 8000
	router = mux.NewRouter()
	router.HandleFunc("/email/{email}", mainEmail).Methods("GET")
	address := fmt.Sprintf(":%d", port)
	log.Fatal(http.ListenAndServe(address, router))
}

func mainEmail(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	uemail := params["email"]

	m := email.NewMessage("subject", "message")
	m.From = mail.Address{Name: "docker_alert", Address: "docker_test_intern@peoplenetonline.com"}
	m.To = strings.Split(uemail, ";")
	err := email.Send(
		"aspmx.l.google.com:25",
		nil,
		m)

	if err != nil {
		log.Fatal(err)

	} else {
		w.WriteHeader(200)
		response := "Email sent"
		json.NewEncoder(w).Encode(response)
	}

}

func main() {
	go regroute()
	forever := make(chan bool)
	<-forever
}
