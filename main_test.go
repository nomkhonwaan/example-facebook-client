package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	fb "github.com/nomkhonwaan/example-facebook-client/facebook"
)

func TestAuthenticate(t *testing.T) {
	expected := `<a href="http://mock-facebook-oauth-dialog-url?client_id=YOUR_FACEBOOK_APP_ID&amp;redirect_uri=http://localhost:8080/">Temporary Redirect</a>.`

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c := fb.NewMockClient(ctrl)
	c.EXPECT().
		Authenticate("YOUR_FACEBOOK_APP_ID", "YOUR_FACEBOOK_APP_SECRET", redirectURI, onLoggedInRedirectURI).
		Return(func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, "http://mock-facebook-oauth-dialog-url?client_id=YOUR_FACEBOOK_APP_ID&redirect_uri="+redirectURI, http.StatusTemporaryRedirect)
		})

	r := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	authenticate(c)(w, r)

	resp := w.Result()

	if resp.StatusCode != http.StatusTemporaryRedirect {
		t.Error("invalid redirect type")
	}
	if data, _ := ioutil.ReadAll(resp.Body); strings.Trim(string(data), "\n") != expected {
		t.Errorf("expected %q but got %q", expected, data)
	}
}

func TestMe(t *testing.T) {
	expected := `You have been logged-in as:
ID: mock-facebook-id
DisplayName: mock-facebook-displayname`

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c := fb.NewMockClient(ctrl)
	c.EXPECT().
		Me().
		Return(map[string]interface{}{
			"id":   "mock-facebook-id",
			"name": "mock-facebook-displayname",
		}, nil)

	s := httptest.NewServer(me(c))
	defer s.Close()

	resp, err := http.Get(s.URL)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if data, _ := ioutil.ReadAll(resp.Body); string(data) != expected {
		t.Errorf("expected %q but got %q", expected, data)
	}
}
