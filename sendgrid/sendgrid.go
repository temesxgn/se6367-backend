package sendgrid

import (
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/temesxgn/se6367-backend/config"
	"sync"
)

var (
	api  *sendgrid.Client
	once sync.Once
	from *mail.Email
)

func init() {
	once.Do(func() {
		api = sendgrid.NewSendClient(config.GetSendGridSecret())
		from = mail.NewEmail("Fairbanks", config.GetNoReplyEmail())
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
func SendEventsTodays() error {
	//file, err := fileutils.LoadEventsTodayEmailTemplate()
	//t, err := template.New("confirmation-email").Parse(file)
	//if err != nil {
	//	return err
	//}
	//
	//buf := new(bytes.Buffer)
	//if err := t.Execute(buf, data); err != nil {
	//	errMsg := fmt.Sprintf("Error executing confirmation email request template for order %v Err: %v", order, err.Error())
	//	//slack.NotifyError(errMsg)
	//	return errors.New(errMsg)
	//}
	//
	//to := mail.NewEmail(data.UserInfo.Name, data.UserInfo.Email)
	//subject := "Thank You For Your Purchase"
	//msg := buildMessage(to, from, subject, "s", string(buf.Bytes()))
	//
	//return sendEmail(msg)
	panic("not implemented")
}
