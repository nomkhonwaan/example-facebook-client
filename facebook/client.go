package facebook

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"text/template"
)

const (
	OAuthDialogURL         = "https://www.facebook.com/v2.10/dialog/oauth?client_id={{.AppID}}&redirect_uri={{.RedirectURI}}"
	ExchangeAccessTokenURL = "https://graph.facebook.com/v2.10/oauth/access_token?client_id={{.AppID}}&redirect_uri={{.RedirectURI}}&client_secret={{.AppSecret}}&code={{.Code}}"
	GraphAPIsURL           = "https://graph.facebook.com"
)

// Client is an interface for doing things with Facebook Graph APIs
type Client interface {
	// Authenticate used to get an access token from Facebook
	Authenticate(appID, appSecret, redirectURI, onLoggedInRedirectURI string) http.HandlerFunc
	// Me returns basic user details
	Me() (map[string]interface{}, error)
}

type ClientImpl struct {
	appID, appSecret string
	accessToken      string
}

func New() *ClientImpl {
	return &ClientImpl{}
}

func (c *ClientImpl) Authenticate(appID, appSecret, redirectURI, onLoggedInRedirectURI string) http.HandlerFunc {
	var buf bytes.Buffer

	return func(w http.ResponseWriter, r *http.Request) {
		if code := r.URL.Query().Get("code"); code != "" {
			data := struct {
				AccessToken string `json:"access_token"`
			}{}
			buf.Reset()
			template.Must(template.New("exchange_access_token").Parse(ExchangeAccessTokenURL)).
				Execute(&buf, map[string]string{
					"AppID":       appID,
					"AppSecret":   appSecret,
					"RedirectURI": redirectURI,
					"Code":        code,
				})
			resp, err := http.Get(buf.String())
			if err != nil {
				log.Printf("an error occurred: %s\n", err.Error())
				return
			}
			defer resp.Body.Close()

			if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
				log.Printf("an error occurred: %s\n", err.Error())
				return
			}
			c.accessToken = data.AccessToken
			http.Redirect(w, r, onLoggedInRedirectURI, http.StatusTemporaryRedirect)
		} else {
			buf.Reset()
			template.Must(template.New("authenticate_url").Parse(OAuthDialogURL)).
				Execute(&buf, map[string]string{
					"AppID":       appID,
					"RedirectURI": redirectURI,
				})
			http.Redirect(w, r, buf.String(), http.StatusTemporaryRedirect)
		}
	}
}

// Me returns basic user profiles
func (c *ClientImpl) Me() (map[string]interface{}, error) {
	var (
		data = make(map[string]interface{})
	)
	resp, err := http.Get(GraphAPIsURL + "/me?access_token=" + c.accessToken)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}
