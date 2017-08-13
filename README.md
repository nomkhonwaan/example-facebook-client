# Example Facebook Client

This is a simple Facebook Log-in and Graph APIs query by using Go `net/http` package.

## Prerequisite

You Facebook Application ID and secret. For quick try with this project just go to [https://developers.facebook.com]() and fillout basic information.

To play this example on you localhost you need to setup `App Domains` to `localhost` and `Site URL` to `http://localhost:8080` in Settings > Basic

## Usage

If you already got Facebook's App ID and secret you must update it in `main.go` on line 13 - 14.

```go
const (
    appID     = "YOUR_FACEBOOK_APP_ID"
    appSecret = "YOUR_FACEBOOK_APP_SECRET"
)
```

Then run this example by using the following commands.

```bash
$ go get -u github.com/nomkhonwaan/example-facebook-client.git
$ cd $GOPATH/src/github.com/nomkhonwaan/example-facebook-client
$ go run main.go
```

The example will start an HTTP server with listening on port `:8080`