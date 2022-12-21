package facebook_test

import (
	"testing"

	_ "github.com/rommms07/idream-erp/core/auth/facebook"
	"github.com/rommms07/idream-erp/helpers/source"
	"github.com/stretchr/testify/assert"
)

const (
	MIN_SDK_VERSION = "v15.0"
)

func Test_appConfigShouldContainTheNecessaryPropsForFacebookLogin(t *testing.T) {
	assert.Equal(t, MIN_SDK_VERSION, source.AppConfig().FbSdkVersion, "Facebook SDK version should match the minimum expected SDK version.")
}

func Test_shouldStartTempHttpServerForAuthCode(t *testing.T) {
	assert.Fail(t, "Not implemented")
}
