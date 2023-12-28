package linuxservicestatus

import (
	"fmt"
	"time"

	"github.com/Jkarage/beemafrica"
)

func SendSMS(service string) error {
	message := fmt.Sprintf("Service %s is not currently in its active status, Current time: %v", service, time.Now())
	client := beemafrica.NewSMS()
	resp, err := client.SendSMS(message, []string{"255713507067", "255615266021"}, "")
	if err != nil {
		return err
	}

	fmt.Println(resp)
	return nil
}
