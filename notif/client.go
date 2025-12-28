package notif

import (
	"fmt"
	"log/slog"
	"net/smtp"

	"github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"
)

type Client struct {
	twilioClient *twilio.RestClient
	originNumber string

	smtpHost string
	smtpPort string
	auth     smtp.Auth
	from     string
}

func NewClient(accountSid, authToken, originNumber, smtpHost, smtpPort, smtpUser, smtpPassword, from string) *Client {
	// Find your Account SID and Auth Token at twilio.com/console
	// and set the environment variables. See http://twil.io/secure
	twilioClient := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})
	auth := smtp.PlainAuth("", smtpUser, smtpPassword, smtpHost)
	return &Client{twilioClient, originNumber, smtpHost, smtpPort, auth, from}
}

func (c *Client) call(destinationNumber, message string) error {
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

func (c *Client) text(destinationNumber, message string) error {
	params := &api.CreateMessageParams{}
	params.SetTo(destinationNumber)
	params.SetFrom(c.originNumber)
	params.SetBody(message)

	resp, err := c.twilioClient.Api.CreateMessage(params)
	if err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	} else {
		if resp.Sid != nil {
			fmt.Println(*resp.Sid)
		} else {
			fmt.Println(resp.Sid)
		}
	}

	return nil
}

func (c *Client) email(destinationEmail, message string) error {
	// Manually construct the email message with headers
	msg := []byte(fmt.Sprintf("To: %s\r\n"+
		"Subject: %s\r\n"+
		"\r\n"+
		"%s\r\n", destinationEmail, "MVC Notification", message))

	// Connect to the server and send the email
	slog.Info("Sending email", "host", c.smtpHost, "port", c.smtpPort, "from", c.from, "to", destinationEmail, "msg", string(msg))
	err := smtp.SendMail(c.smtpHost+":"+c.smtpPort, c.auth, c.from, []string{destinationEmail}, msg)
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}

func (c *Client) SendNotification(destinationNumber, destinationEmail, message string) error {
	if destinationNumber != "" {
		slog.Info("Calling", "number", destinationNumber)
		err := c.call(destinationNumber, message)
		if err != nil {
			return fmt.Errorf("failed to send notification: %w", err)
		}
	}
	if destinationEmail != "" {
		slog.Info("Sending email", "email", destinationEmail)
		err := c.email(destinationEmail, message)
		if err != nil {
			return fmt.Errorf("failed to send notification: %w", err)
		}
	}
	return nil
}
