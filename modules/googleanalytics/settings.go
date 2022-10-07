package googleanalytics

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const (
	defaultFocusable = false
	defaultTitle     = "Google Analytics"
)

type Settings struct {
	*cfg.Common

	months         int
	secretFile     string `help:"Your Google client secret JSON file." values:"A string representing a file path to the JSON secret file."`
	viewIds        map[string]any
	enableRealtime bool
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {

	settings := Settings{
		Common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),

		months:         ymlConfig.UInt("months"),
		secretFile:     ymlConfig.UString("secretFile"),
		viewIds:        ymlConfig.UMap("viewIds"),
		enableRealtime: ymlConfig.UBool("enableRealtime", false),
	}

	settings.SetDocumentationPath("google/analytics")

	return &settings
}
