package main

import (
	"fmt"
	"log"
	"net/http"

	fb "github.com/nomkhonwaan/example-facebook-client/facebook"
)

const (
	appID     = "YOUR_FACEBOOK_APP_ID"
	appSecret = "YOUR_FACEBOOK_APP_SECRET"
)

var (
	redirectURI           = "http://localhost:8080/"
	onLoggedInRedirectURI = "http://localhost:8080/me"
)

func main() {
	c := fb.New()

	http.HandleFunc("/", authenticate(c))
	http.HandleFunc("/me", me(c))

	log.Println("server has been listening on :8080")
	http.ListenAndServe(":8080", nil)
}

func authenticate(c fb.Client) http.HandlerFunc {
	return c.Authenticate(appID, appSecret, redirectURI, onLoggedInRedirectURI)
}

func me(c fb.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := c.Me()
		if err != nil {
			fmt.Fprintf(w, "An error has occurred: %s", err.Error())
			return
		}
		fmt.Fprintf(w, `You have been logged-in as:
ID: %s
DisplayName: %s`, data["id"], data["name"])
	}
}
