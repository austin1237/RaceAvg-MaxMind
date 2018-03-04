package validator

import (
	"errors"

	"github.com/asaskevich/govalidator"
)

func ValidateLocationQuery(clientIP string) error {
	if clientIP == "" {
		err := errors.New("A ipv4 address was missing from the param ip in the query string")
		return err
	}

	if !govalidator.IsIPv4(clientIP) {
		err := errors.New("Invalid ipv4 address sent in query string")
		return err
	}
	return nil
}
