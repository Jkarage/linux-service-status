package linuxservicestatus

import (
	"fmt"
	"time"

	"github.com/Jkarage/beemafrica"
)

func SendSMS(service string) error {
	message := fmt.Sprintf("Service %s is not currently in it's active state, Current time: %v", service, time.Now())
	client := beemafrica.NewSMS()
	_, err := client.SendSMS(message, []string{"255713507067", "255757294146"}, "")
	if err != nil {
		return err
	}

	return nil
}
