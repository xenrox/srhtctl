package config

import (
	"fmt"
	"os"
	"strings"

	"git.xenrox.net/~xenrox/srhtctl/helpers/errorhelper"
	"github.com/vaughan0/go-ini"
)

// ConfigPath is a path to a non default config.ini file
var ConfigPath string

var configFile *ini.File

// LoadConfig parses the config.ini file and returns it.
func LoadConfig() *ini.File {
	var configPath string
	if ConfigPath == "" {

		xdgConfigHome, err := os.UserConfigDir()
		if err != nil {
			errorhelper.ExitError(err)
		}
		configPath = fmt.Sprintf("%s/srhtctl/config.ini", xdgConfigHome)
	} else {
		configPath = ConfigPath
	}
	file, err := ini.LoadFile(configPath)
	if err != nil {
		errorhelper.ExitError(err)
	}
	return &file
}

// InitConfig gets called by cobra and only calls LoadConfig.
func InitConfig() {
	configFile = LoadConfig()
}

// GetConfigValue returns a value from the config.ini file as a string.
// It is possible to set a default value as a third argument.
func GetConfigValue(section string, key string, defaultValue ...string) string {
	if configFile == nil {
		InitConfig()
	}
	value, ok := configFile.Get(section, key)
	if !ok {
		if len(defaultValue) > 0 {
			value = defaultValue[0]
		} else {
			fmt.Fprintf(os.Stderr, "%s missing from section %s in config.ini.\n", key, section)
			os.Exit(1)
		}
	}
	return value
}

// GetURL returns a formatted url for a service
func GetURL(service string) string {
	url := GetConfigValue(service, "url", fmt.Sprintf("%s.sr.ht", service))
	if strings.HasPrefix(url, "https://") {
		return url
	}
	return fmt.Sprintf("https://%s", url)
}
