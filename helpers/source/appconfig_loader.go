package source

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/rommms07/idream-erp/config"
)

// gormConfig schema is used by the appConfigType that contains the struct info of our gormConfig
// defined in $ROOTDIR/config/app_config.json; If you want to add an extra fields to the appConfig.gormConfig
// you can update this schema to incldue the newly added field to the parsed config.
type gormConfig struct {
	DefaultStringSize uint64
	DisableDateTimePrecision, DontSupportRenameIndex, DontSupportRenameColumn, SkipInitVersion bool
}

// appVersion struct is the schema for the parsed version defined in the app_config.json if the version
// is not formatted properly `<major>.<minor>.<build>-<release>` the output will get truncated by the
// `loadConfig`.
type appVersion struct {
	Major   uint64
	Minor   uint64
	Build   uint64
	Release string
}

// appConfigType is the map to which the $ROODIR/config/app_config.json will be based upon on,
// any field in the app_config.json that does not corresponds to any of the fields of appConfigType
// will inevitably ignored by the `loadConfig`
type appConfigType struct {
	Version        string
	VersionInfo    *appVersion
	FbSdkVersion   string
	FbClientId     string
	FbClientSecret string
	FbRedirectUri  string
	Message        string
	GormConfig     *gormConfig
}

var (
	loadedConfig *appConfigType
)

// parseVersion parses the version defined in the app_config.json, since this function can be called anywhere
// in the local scope of this package, it can be used to parse any string that satisfies the defined format.
func parseVersion(v string) *appVersion {
	const (
		MAJOR   = "major"
		MINOR   = "minor"
		BUILD   = "build"
		RELEASE = "release"
	)

	vpatt := regexp.MustCompile(
		fmt.Sprintf(`(?P<%s>\d+?)[.](?P<%s>\d+?)[.](?P<%s>\d+?)\-(?P<%s>(alpha|beta|build|testing))`, MAJOR, MINOR, BUILD, RELEASE),
	)

	dmatch := vpatt.FindStringSubmatchIndex(v)

	major, _ := strconv.ParseUint(string(vpatt.ExpandString([]byte{}, "$"+MAJOR, v, dmatch)), 10, 64)
	minor, _ := strconv.ParseUint(string(vpatt.ExpandString([]byte{}, "$"+MINOR, v, dmatch)), 10, 64)
	build, _ := strconv.ParseUint(string(vpatt.ExpandString([]byte{}, "$"+BUILD, v, dmatch)), 10, 64)

	return &appVersion{major, minor, build, string(vpatt.ExpandString([]byte{}, "$"+RELEASE, v, dmatch))}
}

// loadConfig is the function that will be called by `AppConfig` to load the app_config.json file and parse its
// content to fit into the appConfigType struct. This can be called by any batch codes that modifies the
// app_config.json at runtime to rehydrate the `loadedConfig` struct.
func loadConfig() {
	loadedConfig = &appConfigType{}

	b, err := os.ReadFile(config.DEFAULT)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error loading app_config.json: %s", err.Error())
		os.Exit(1)
	}

	err = json.Unmarshal(b, &loadedConfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error unmarshaling app_config.json: %s", err.Error())
		os.Exit(1)
	}

	sdkverpatt := regexp.MustCompile(`^v\d{2,}[.]\d{1}$`)

	fbSdkVer := os.Getenv("FB_SDK_VERSION")
	fbClientId := os.Getenv("FB_CLIENT_ID")
	fbClientSecret := os.Getenv("FB_CLIENT_SECRET")
	fbRedirectUri := os.Getenv("FB_REDIRECT_URI")

	if !sdkverpatt.MatchString(fbSdkVer) {
		fmt.Fprintf(os.Stderr, "FB_SDK_VERSION did not satisfy the expected version regexp.")
		os.Exit(1)
	}

	loadedConfig.VersionInfo = parseVersion(loadedConfig.Version)	
	loadedConfig.FbClientId = fbClientId
	loadedConfig.FbClientSecret = fbClientSecret
	loadedConfig.FbSdkVersion = fbSdkVer
	loadedConfig.FbRedirectUri = fbRedirectUri
}

// AppConfig returns the `loadedConfig` struct locally defined in this scope.
func AppConfig() *appConfigType {
	if loadedConfig == nil {
		loadConfig()
	}

	return loadedConfig
}
