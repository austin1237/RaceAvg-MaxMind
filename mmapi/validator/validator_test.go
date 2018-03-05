package validator

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateLocationQueryNoIp(t *testing.T) {
	mockErr := errors.New("A ipv4 address was missing from the param ip in the query string")
	err := ValidateLocationQuery("")
	assert.Equal(t, mockErr, err)
}

func TestValidateLocationQueryInvalidIp(t *testing.T) {
	mockErr := errors.New("Invalid ipv4 address sent in query string")
	err := ValidateLocationQuery("Invalid Ip")
	assert.Equal(t, mockErr, err)
}

func TestValidateLocationQueryValid(t *testing.T) {
	err := ValidateLocationQuery("174.29.5.118")
	assert.Nil(t, err)
}
