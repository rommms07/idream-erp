package config

import (
	"fmt"
	"os"
)

var (
	ROOTDIR = os.Getenv("ROOTDIR")

	// Default path for the app_config.json
	DEFAULT = fmt.Sprintf("%s/config/app_config.json", ROOTDIR)
)
