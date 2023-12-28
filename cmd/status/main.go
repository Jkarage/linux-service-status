package main

import (
	"flag"
	"log"
	"os/exec"

	linuxservicestatus "github.com/jkarage/linux-service-status"
	"github.com/joho/godotenv"
)

var (
	service = flag.String("service", "", "A service to look for it's status")
)

func init() {
	godotenv.Load()
}

func main() {
	flag.Parse()

	cmd := exec.Command("systemctl", "is-active", *service)
	err := cmd.Run()
	if err != nil {
		err = linuxservicestatus.SendSMS(*service)
		if err != nil {
			log.Fatal(err)
		}
		err = linuxservicestatus.SendMail(*service)
		if err != nil {
			log.Fatal(err)
		}
	}
}
