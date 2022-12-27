package facebook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/rommms07/idream-erp/helpers/loader"
)

func (token *FacebookAccessToken) GetLongLivedToken() (res_token *FacebookAccessToken, err error) {
	res_token = &FacebookAccessToken{}
	fbGraphUrl, _ := url.Parse(fmt.Sprintf("%s/oauth/access_token", FACEBOOK_GRAPH))
	config := loader.AppConfig()
	q := fbGraphUrl.Query()

	q.Add("grant_type", "fb_exchange_token")
	q.Add("client_id", config.FbClientId)
	q.Add("client_secret", config.FbClientSecret)
	q.Add("fb_exchange_token", token.Access_token)

	fbGraphUrl.RawQuery = q.Encode()

	for Nt := 5; Nt > 0; Nt-- {
		res, err := http.Get(fbGraphUrl.String())
		if err != nil {
			continue
		}

		defer res.Body.Close()

		buf := bytes.NewBuffer([]byte{})

		_, err = buf.ReadFrom(res.Body)
		if err != nil {
			continue
		}

		err = json.Unmarshal(buf.Bytes(), res_token)
		if err != nil {
			continue
		}

		break
	}

	if err != nil {
		return nil, err
	}

	return
}
