// This package setups a temporary http server, this is done so that we could
// get the authorization code coming from the redirect_uri.

package facebook

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rommms07/idream-erp/config"
	"github.com/rommms07/idream-erp/helpers/source"
)

type FacebookLoginOptions struct {
	LoginUrl     string
	ClientId     string
	ClientSecret string
	RedirectUri  string
	State        map[string]string
	Code         string
	Declined     bool
	Token        *FacebookAccessToken
	Pending      chan struct{}

	ErrorReason string
	Error       string
	ErrorDesc   string
}

type FacebookAccessToken struct {
	Access_token string
	Expires_in   uint64
	Token_type   string
}

var (
	FACEBOOK_LOGIN_DIALOG = fmt.Sprintf("https://www.facebook.com/%s/dialog/oauth", source.AppConfig().FbSdkVersion)
	FACEBOOK_GRAPH        = fmt.Sprintf("https://graph.facebook.com/%s", source.AppConfig().FbSdkVersion)

	pendingLoginReq = make(map[string]*FacebookLoginOptions)
)

func get_def_opts(q url.Values, opts *FacebookLoginOptions) (string, string, string) {
	redirect_uri, client_id, client_secret := "", "", ""

	if len(opts.RedirectUri) != 0 {
		redirect_uri = opts.RedirectUri
	} else {
		redirect_uri = source.AppConfig().FbRedirectUri
	}

	if len(opts.ClientId) != 0 {
		client_id = opts.ClientId
	} else {
		client_id = source.AppConfig().FbClientId
	}

	if len(opts.ClientSecret) != 0 {
		client_secret = opts.ClientSecret
	} else {
		client_secret = source.AppConfig().FbClientSecret
	}

	return redirect_uri, client_id, client_secret
}

// make_fblogin_url returns a facebook login url which can be use to generate an authorization code.
func make_fblogin_url(opts *FacebookLoginOptions) string {

	login, err := url.Parse(FACEBOOK_LOGIN_DIALOG)
	if err != nil {
		return ""
	}

	q := login.Query()
	redirect_uri, client_id, _ := get_def_opts(q, opts)

	q.Add("client_id", client_id)
	q.Add("redirect_uri", redirect_uri)

	b, err := json.Marshal(opts.State)
	if err != nil {
		return ""
	}

	if len(b) != 0 && string(b) != "null" {
		q.Add("state", string(b))
	}

	login.RawQuery = q.Encode()
	return login.String()
}

func write_rp(w io.Writer, data any) error {
	tmpl, err := template.ParseFiles(config.ROOTDIR + "/core/auth/facebook/static/redirect.html")
	if err != nil {
		return err
	}

	return tmpl.Execute(w, data)
}

// exchange_code_to_token is responsible for exchanging the authorization code that comes from Facebook
// to an access token.
func exchange_code_to_token(opts *FacebookLoginOptions) (token *FacebookAccessToken, err error) {
	token = &FacebookAccessToken{}
	exchanger, _ := url.Parse(fmt.Sprintf("%s/oauth/access_token", FACEBOOK_GRAPH))
	q := exchanger.Query()
	redirect_uri, client_id, client_secret := get_def_opts(q, opts)

	q.Add("client_id", client_id)
	q.Add("client_secret", client_secret)
	q.Add("redirect_uri", redirect_uri)
	q.Add("code", opts.Code)

	exchanger.RawQuery = q.Encode()

	for Nt := 0; Nt < 5; Nt++ {
		res, err := http.Get(exchanger.String())
		if err != nil {
			continue
		}

		defer res.Body.Close()

		buf := bytes.NewBuffer([]byte{})
		_, err = io.Copy(buf, res.Body)
		if err != nil {
			continue
		}

		err = json.Unmarshal(buf.Bytes(), token)
		if err != nil {
			continue
		}

		break
	}

	opts.Token = token
	return
}

// FbRedirectHandler is the request handler for invoking the Facebook login flow. Usually this
// kind of handler must be guarded by a rate limiting middleware to avoid someone abuse the
// this handler or possible take down the server by overflowing the `pendingLoginRp`
func FbRedirectHandler(c *gin.Context) {
	state := make(map[string]string)
	err := json.Unmarshal([]byte(c.Request.URL.Query().Get("state")), &state)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]any{
			"status_code": http.StatusBadRequest,
			"error":       err.Error(),
		})
		return
	}

	err = write_rp(c.Writer, pendingLoginReq[state["uuid"]])
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]any{
			"status_code": http.StatusBadRequest,
			"error":       err.Error(),
		})
		return
	}

	// declined request
	if c.Query("error_reason") != "" {
		pendingLoginReq[state["uuid"]].ErrorReason = c.Query("error_reason")
		pendingLoginReq[state["uuid"]].ErrorDesc = c.Query("error_description")
		pendingLoginReq[state["uuid"]].Error = c.Query("error")
		pendingLoginReq[state["uuid"]].Declined = true
	}

	// Store authorization code
	pendingLoginReq[state["uuid"]].Code = c.Query("code")

	// Signal `Login` that the login phase has now been fulfilled or not by sending a
	// struct{}{} to the semaphore defined in pendingLoginReq.
	pendingLoginReq[state["uuid"]].Pending <- struct{}{}
}

// Login is where we connect all of the things we have defined above. It is the
// function to which we call when we want to start the Facebook login flow and
// to get a short-lived user access token from Facebook.
func Login(opts *FacebookLoginOptions) (*FacebookAccessToken, error) {
	opts.Pending = make(chan struct{})
	opts.State = map[string]string{
		"uuid": uuid.NewString(),
	}

	pendingLoginReq[opts.State["uuid"]] = opts

	if len(opts.LoginUrl) == 0 {
		opts.LoginUrl = make_fblogin_url(opts)
	}

	select {
	case <-pendingLoginReq[opts.State["uuid"]].Pending:
		if opts.Declined {
			b, _ := json.Marshal(map[string]string{
				"error":             opts.Error,
				"error_reason":      opts.ErrorReason,
				"error_description": opts.ErrorDesc,
			})

			return nil, errors.New(string(b))
		}
	case <-time.After(time.Minute * 15):
		return nil, errors.New("error: facebook login timeout")
	}

	// Exchange the received authorization code for a new access token.
	token, err := exchange_code_to_token(opts)
	if err != nil {
		return nil, err
	}

	return token, nil
}
