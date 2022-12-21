// This package setups a temporary http server, this is done so that we could
// get the authorization code coming from the redirect_uri.

package facebook

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/rommms07/idream-erp/helpers/source"
	"golang.org/x/net/context"
)

var (
	FACEBOOK_LOGIN_DIALOG = fmt.Sprintf("https://www.facebook.com/%s/dialog/oauth", source.AppConfig().FbSdkVersion)
	FACEBOOK_GRAPH        = fmt.Sprintf("https://graph.facebook.com/%s", source.AppConfig().FbSdkVersion)
)

func init() {
	url, _ := gen_fblogin_url("")
	println(url)
}

func gen_fblogin_url(reqState string) (string, error) {
	login, err := url.Parse(FACEBOOK_LOGIN_DIALOG)
	if err != nil {
		return "", err
	}

	q := login.Query()

	q.Add("client_id", source.AppConfig().FbClientId)
	q.Add("redirect_uri", source.AppConfig().FbRedirectUri)
	q.Add("response_type", "code")

	if len(reqState) != 0 {
		q.Add("state", reqState)
	}

	login.RawQuery = q.Encode()

	println(login.String())

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*15)
	defer cancel()

	max_attempts := 5

	for max_attempts > 0 {
		req, _ := http.NewRequestWithContext(ctx, http.MethodGet, login.String(), nil)
		if err != nil {
			max_attempts--
			continue
		}

		res, _ := http.DefaultClient.Do(req)
		if err != nil {
			max_attempts--
			continue
		}

		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			max_attempts--
		} else {
			println(res.Header.Get("location"))
		}
	}

	return "", nil
}
