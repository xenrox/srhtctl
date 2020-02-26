package config

import (
	"fmt"
	"log"
	"os"
	"strings"

	"git.xenrox.net/~xenrox/srhtctl/helpers/errorhelper"
	"github.com/vaughan0/go-ini"
)

var configFile ini.File

// LoadConfig parses the config.ini file and returns it.
func LoadConfig() ini.File {
	xdgConfigHome, err := os.UserConfigDir()
	if err != nil {
		errorhelper.ExitError(err)
	}
	configPath := fmt.Sprintf("%s/srhtctl/config.ini", xdgConfigHome)
	file, err := ini.LoadFile(configPath)
	if err != nil {
		errorhelper.ExitError(err)
	}
	return file
}

// InitConfig gets called by cobra and only calls LoadConfig.
func InitConfig() {
	configFile = LoadConfig()
}

// GetConfigValue returns a value from the config.ini file as a string.
// It is possible to set a default value as a third argument.
func GetConfigValue(section string, key string, defaultValue ...string) string {
	value, ok := configFile.Get(section, key)
	if !ok {
		if len(defaultValue) > 0 {
			value = defaultValue[0]
		} else {
			log.Fatalf("%s missing from section %s in config.ini.\n", key, section)
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
