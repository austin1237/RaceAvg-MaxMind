package main

import (
	"net/http"
	"testing"

	"github.com/appleboy/gofight"
	"github.com/stretchr/testify/assert"
)

func TestHealthEndpoint(t *testing.T) {
	r := gofight.New()

	r.GET("/v1/health").
		// turn on the debug mode.
		SetDebug(true).
		Run(CreateApi(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusOK, r.Code)
		})
}

func TestLocationEndpoint(t *testing.T) {
	r := gofight.New()

	r.GET("/v1/location?ip=\"174.29.5.118\"").
		// turn on the debug mode.
		SetDebug(true).
		Run(CreateApi(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusOK, r.Code)
		})
}
