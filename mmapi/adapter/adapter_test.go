package adapter

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClientIPToDbValid(t *testing.T) {
	dbIP, err := ClientIPToDb("174.36.207.186")
	assert.Nil(t, err)
	assert.Equal(t, 2921648058, dbIP)
}

func TestClientIPToDbNoIP(t *testing.T) {
	invalidIP := "not a valid ip"
	mockErr := errors.New("Invalid IP address " + invalidIP)
	_, err := ClientIPToDb(invalidIP)
	assert.Equal(t, mockErr, err)
}
