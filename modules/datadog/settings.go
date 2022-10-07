package datadog

import (
	"os"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const (
	defaultFocusable = true
	defaultTitle     = "DataDog"
)

type Settings struct {
	*cfg.Common

	apiKey         string `help:"Your Datadog API key."`
	applicationKey string `help:"Your Datadog Application key."`
	tags           []any  `help:"Array of tags you want to query monitors by."`
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {

	settings := Settings{
		Common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),

		apiKey:         ymlConfig.UString("apiKey", ymlConfig.UString("apikey", os.Getenv("WTF_DATADOG_API_KEY"))),
		applicationKey: ymlConfig.UString("applicationKey", os.Getenv("WTF_DATADOG_APPLICATION_KEY")),
		tags:           ymlConfig.UList("monitors.tags"),
	}

	cfg.ModuleSecret(name+"-api", globalConfig, &settings.apiKey).Load()
	cfg.ModuleSecret(name+"-app", globalConfig, &settings.applicationKey).Load()

	return &settings
}
