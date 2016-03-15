package integration_test

import (
	"net/http"
	"testing"

	"github.com/go-errors/errors"
	fosite "github.com/ory-am/fosite"
	"github.com/stretchr/testify/assert"
)

func tokenInfoHandler(t *testing.T, oauth2 fosite.OAuth2Provider, session interface{}) func(rw http.ResponseWriter, req *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		token := req.Header.Get("Authorization")
		t.Logf("Got bearer token in header: %s", token)
		if token == "" {
			rw.WriteHeader(http.StatusNotAcceptable)
		}
		rw.WriteHeader(http.StatusNoContent)
	}
}

func authEndpointHandler(t *testing.T, oauth2 fosite.OAuth2Provider, session interface{}) func(rw http.ResponseWriter, req *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		ctx := fosite.NewContext()

		ar, err := oauth2.NewAuthorizeRequest(ctx, req)
		if err != nil {
			t.Logf("Access request failed because %s.", err.Error())
			t.Logf("Request: %s.", ar)
			t.Logf("Stack: %s.", err.(*errors.Error).ErrorStack())
			oauth2.WriteAuthorizeError(rw, ar, err)
			return
		}

		// Normally, this would be the place where you would check if the user is logged in and gives his consent.
		// For this test, let's assume that the user exists, is logged in, and gives his consent...

		response, err := oauth2.NewAuthorizeResponse(ctx, req, ar, session)
		if err != nil {
			t.Logf("Access request failed because %s.", err.Error())
			t.Logf("Request: %s.", ar)
			t.Logf("Stack: %s.", err.(*errors.Error).ErrorStack())
			oauth2.WriteAuthorizeError(rw, ar, err)
			return
		}

		oauth2.WriteAuthorizeResponse(rw, ar, response)
	}
}

func authCallbackHandler(t *testing.T) func(rw http.ResponseWriter, req *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		q := req.URL.Query()
		if q.Get("code") == "" && q.Get("error") == "" {
			assert.NotEmpty(t, q.Get("code"))
			assert.NotEmpty(t, q.Get("error"))
		}

		if q.Get("code") != "" {
			rw.Write([]byte("code: ok"))
		}
		if q.Get("error") != "" {
			rw.WriteHeader(http.StatusNotAcceptable)
			rw.Write([]byte("error: " + q.Get("error")))
		}

	}
}

func tokenEndpointHandler(t *testing.T, oauth2 fosite.OAuth2Provider) func(rw http.ResponseWriter, req *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		ctx := fosite.NewContext()
		var mySessionData struct {
			Foo string
		}

		accessRequest, err := oauth2.NewAccessRequest(ctx, req, &mySessionData)
		if err != nil {
			t.Logf("Access request failed because %s.", err.Error())
			t.Logf("Request: %s.", accessRequest)
			t.Logf("Stack: %s.", err.(*errors.Error).ErrorStack())
			oauth2.WriteAccessError(rw, accessRequest, err)
			return
		}

		response, err := oauth2.NewAccessResponse(ctx, req, accessRequest)
		if err != nil {
			t.Logf("Access request failed because %s.", err.Error())
			t.Logf("Request: %s.", accessRequest)
			t.Logf("Stack: %s.", err.(*errors.Error).ErrorStack())
			oauth2.WriteAccessError(rw, accessRequest, err)
			return
		}

		oauth2.WriteAccessResponse(rw, accessRequest, response)
	}
}
