package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"

	"github.com/Jkarage/linuxservicestatus"
	"github.com/joho/godotenv"
)

var (
	service = flag.String("svc", "", "A service to look for it's status")
)

func init() {
	godotenv.Load()
}

func main() {
	flag.Parse()
	fmt.Println(*service)

	cmd := exec.Command("systemctl", "is-active", *service)
	err := cmd.Run()
	if err != nil {
		err = linuxservicestatus.SendSMS(*service)
		if err != nil {
			log.Fatal(err)
		}
		// err = linuxservicestatus.SendMail(*service)
		// if err != nil {
		// 	log.Fatal(err)
		// }
	}
}
