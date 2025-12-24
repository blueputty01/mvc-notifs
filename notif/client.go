package notif

import (
	"fmt"

	"github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"
)

type Client struct {
	twilioClient *twilio.RestClient
	originNumber string
}

func NewClient(accountSid, authToken, originNumber string) *Client {
	// Find your Account SID and Auth Token at twilio.com/console
	// and set the environment variables. See http://twil.io/secure
	twilioClient := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})
	return &Client{twilioClient, originNumber}
}

func (c *Client) SendNotification(destinationNumber, message string) error {
	params := &api.CreateCallParams{}
	params.SetTo(destinationNumber)
	params.SetFrom(c.originNumber)

	params.SetTwiml(fmt.Sprintf("<Response><Say>%s</Say></Response>", message))

	resp, err := c.twilioClient.Api.CreateCall(params)
	if err != nil {
		return fmt.Errorf("failed to create call: %w", err)
	} else {
		if resp.Sid != nil {
			fmt.Println(*resp.Sid)
		} else {
			fmt.Println(resp.Sid)
		}
	}

	return nil
}
