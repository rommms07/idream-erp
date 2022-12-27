package facebook_test

import (
	"bytes"
	"fmt"
	"net/url"
	"testing"

	"github.com/rommms07/idream-erp/core/auth/facebook"
	"github.com/rommms07/idream-erp/helpers/source"
	"github.com/stretchr/testify/assert"
)

const (
	MIN_SDK_VERSION = "v15.0"
)

func Test_appConfigShouldContainTheNecessaryPropsForFacebookLogin(t *testing.T) {
	assert.Equal(t, MIN_SDK_VERSION, source.AppConfig().FbSdkVersion, "Facebook SDK version should match the minimum expected SDK version.")
}

func Test_shouldCreateAnFbLoginUrl(t *testing.T) {
	q := url.Values{}

	q.Add("client_id", source.AppConfig().FbClientId)
	q.Add("redirect_uri", source.AppConfig().ServerAddr)

	xpcted_url := facebook.FACEBOOK_LOGIN_DIALOG + "?" + q.Encode()
	url := facebook.MakeFbLoginUrl(&facebook.FacebookLoginOptions{RedirectUri: source.AppConfig().ServerAddr})

	assert.Equal(t, xpcted_url, url, "make_fblogin_url did not returned the expected facebook login url.")
}

func Test_writeRpShouldExecuteAndWriteTheTemplateInAMem(t *testing.T) {
	buf := bytes.NewBuffer([]byte{})

	err := facebook.WriteRp(buf, nil)
	if err != nil {
		assert.Nil(t, err, fmt.Sprintf("There was an error executing the template. (error: %s)", err.Error()))
	}

	if buf.Len() == 0 {
		assert.Fail(t, "write_rp did not execute the expected template.")
	}
}

func Test_shouldCreateAnUrlForExhangingTheAuthCode(t *testing.T) {
	assert.Fail(t, "exchange_code_to_token is implemented but not being unit test.")
}
