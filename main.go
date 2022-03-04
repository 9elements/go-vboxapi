package main

import (
	"log"
	"os"

	"github.com/9elements/go-vboxapi/vboxapi"
)

func main() {
	url := "http://127.0.0.1:18083"
	if len(os.Args) >= 2 {
		url = os.Args[1]
	}

	client := vboxapi.New("", "", url, false, "")
	if err := client.Logon(); err != nil {
		log.Fatalf("Unable to log on to vboxweb: %v\n", err)
	}
}
