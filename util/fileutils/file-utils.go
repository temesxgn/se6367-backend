package fileutils

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

// LoadFileAsString - loads the file as the given path and returns as string
func LoadFileAsString(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("File reading error", err)
		return "", nil
	}

	return string(data), nil
}

// LoadCreateOrderRequestTemplate - loads the create order request template file
func LoadCreateOrderRequestTemplate() (string, error) {
	path, _ := filepath.Abs("./slack/templates/create-order-request.txt")
	return LoadFileAsString(path)
}

// LoadCreateOrderConfirmationEmailTemplate - loads the create order confirmation email template file
func LoadCreateOrderConfirmationEmailTemplate() (string, error) {
	path, _ := filepath.Abs("./sendgrid/templates/confirmation-email.html")
	return LoadFileAsString(path)
}

// LoadCancelOrderEmailTemplate - loads the cancel order email template file
func LoadCancelOrderEmailTemplate() (string, error) {
	path, _ := filepath.Abs("./sendgrid/templates/cancel-order.html")
	return LoadFileAsString(path)
}
