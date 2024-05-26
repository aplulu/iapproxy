package config

import (
	"fmt"
	"regexp"

	"github.com/kelseyhightower/envconfig"
)

type config struct {
	Listen                          string   `envconfig:"listen" default:""`
	Port                            string   `envconfig:"port" default:"8080"`
	GoogleImpersonateServiceAccount string   `envconfig:"google_impersonate_service_account" default:""`
	URLPatterns                     []string `envconfig:"url_patterns" default:""`
	IAPClientID                     string   `envconfig:"iap_client_id" default:""`
	ProxyVerbose                    bool     `envconfig:"proxy_verbose" default:"false"`
}

var conf config
var urlPatterns []regexp.Regexp

// LoadConf loads the configuration from the environment variables.
func LoadConf() error {
	if err := envconfig.Process("", &conf); err != nil {
		return fmt.Errorf("config.LoadConf: failed to load conf: %w", err)
	}

	for _, pattern := range conf.URLPatterns {
		re, err := regexp.Compile(pattern)
		if err != nil {
			return fmt.Errorf("config.LoadConf: failed to compile URL pattern: %s: %w", pattern, err)
		}
		urlPatterns = append(urlPatterns, *re)
	}

	return nil
}

func Listen() string {
	return conf.Listen
}

func Port() string {
	return conf.Port
}

func GoogleImpersonateServiceAccount() string {
	return conf.GoogleImpersonateServiceAccount
}

func URLPatterns() []regexp.Regexp {
	return urlPatterns
}

func IAPClientID() string {
	return conf.IAPClientID
}

func ProxyVerbose() bool {
	return conf.ProxyVerbose
}
