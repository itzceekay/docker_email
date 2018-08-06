package docker_email

import (
	"log"
	"net/mail"
	"strings"

	"github.com/scorredoira/email"
)

func main() {
	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.

	m := email.NewMessage("subject", "message")
	m.From = mail.Address{Name: "docker_alert", Address: "docker_test_intern@peoplenetonline.com"}
	m.To = strings.Split("kirubanand_chandrasekaran@trimble.com", ";")
	err := email.Send(
		"aspmx.l.google.com:25",
		nil,
		m)

	if err != nil {
		log.Fatal(err)

	}

}
