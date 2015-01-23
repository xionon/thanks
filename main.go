package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

var recipient string
var message string

func init() {
	flag.StringVar(&recipient, "to", "", "The name of the project or person you are thanking")
	flag.StringVar(&message, "message", "", "The particular message you want to send them")
	flag.Parse()
}

func main() {
  route := "http://localhost:3000/messages"
  email := os.Getenv("THANKS_EMAIL")
  token := os.Getenv("THANKS_TOKEN")

	_, err := http.PostForm(route, url.Values{"message[recipient]": {recipient}, "message[body]": {message}, "format": {"json"}, "email": {email}, "token": {token}})

	if err != nil {
		fmt.Println("Uh-oh, we ran into an error sending your thank-you!")
		fmt.Println(err)
		return
	}

	fmt.Fprintf(os.Stdout, "Your thank-you to %s is on the way!\n", recipient)
}
