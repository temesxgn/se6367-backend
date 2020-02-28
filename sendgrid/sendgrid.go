package sendgrid

import (
	"bytes"
	"errors"
	"fmt"
	"sync"
	"text/template"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/temesxgn/hermes/config"
	"github.com/temesxgn/hermes/resolver/models"
	"github.com/temesxgn/hermes/slack"
	"github.com/temesxgn/hermes/util/fileutils"
)

var (
	api  *sendgrid.Client
	once sync.Once
	from *mail.Email
)

func init() {
	once.Do(func() {
		api = sendgrid.NewSendClient(config.GetSendGridSecret())
		from = mail.NewEmail("MegaSweets", config.GetNoReplyEmail())
	})
}

func sendEmail(message *mail.SGMailV3) error {
	response, err := api.Send(message)
	if err != nil {
		return err
	}

	fmt.Println(response.StatusCode)
	fmt.Println(response.Body)
	fmt.Println(response.Headers)

	return nil
}

func buildMessage(to, from *mail.Email, subject, plainText, html string) *mail.SGMailV3 {
	return mail.NewSingleEmail(from, subject, to, plainText, html)
}

// SendCreateOrderConfirmationEmail -
func SendCreateOrderConfirmationEmail(order string, data models.CheckoutParams) error {
	file, err := fileutils.LoadCreateOrderConfirmationEmailTemplate()
	t, err := template.New("confirmation-email").Parse(file)
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	if err := t.Execute(buf, data); err != nil {
		errMsg := fmt.Sprintf("Error executing confirmation email request template for order %v Err: %v", order, err.Error())
		slack.NotifyError(errMsg)
		return errors.New(errMsg)
	}

	to := mail.NewEmail(data.UserInfo.Name, data.UserInfo.Email)
	subject := "Thank You For Your Purchase"
	msg := buildMessage(to, from, subject, "s", string(buf.Bytes()))

	return sendEmail(msg)
}

// SendCancelOrderConfirmationEmail -
func SendCancelOrderConfirmationEmail(order, name, email string) error {
	file, err := fileutils.LoadCancelOrderEmailTemplate()
	t, err := template.New("cancel-order-email").Parse(file)
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	if err := t.Execute(buf, order); err != nil {
		errMsg := fmt.Sprintf("Error executing cancel order email request template for order %v Err: %v", order, err.Error())
		slack.NotifyError(errMsg)
		return errors.New(errMsg)
	}

	to := mail.NewEmail(name, email)
	subject := "Sorry To See You Leave"
	msg := buildMessage(to, from, subject, "s", string(buf.Bytes()))

	return sendEmail(msg)
}
