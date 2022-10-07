package pagerduty

import (
	"os"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const (
	defaultFocusable = false
	defaultTitle     = "PagerDuty"
)

// Settings defines the configuration properties for this module
type Settings struct {
	*cfg.Common

	apiKey           string `help:"Your PagerDuty API key."`
	escalationFilter []any  `help:"An array of schedule names you want to filter the OnCalls on."`
	myName           string `help:"The name to highlight when on-call in PagerDuty."`
	scheduleIDs      []any  `help:"An array of schedule IDs you want to restrict the OnCalls query to."`
	showIncidents    bool   `help:"Whether or not to list incidents." optional:"true"`
	showOnCallEnd    bool   `help:"Whether or not to display the date the oncall schedule ends." optional:"true"`
	showSchedules    bool   `help:"Whether or not to show schedules." optional:"true"`
	teamIDs          []any  `help:"An array of team IDs to restrict the incidents query to" optional:"true"`
	userIDs          []any  `help:"An array of user IDs to restrict the incidents query to" optional:"true"`
}

// NewSettingsFromYAML creates a new settings instance from a YAML config block
func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	settings := Settings{
		Common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),

		apiKey:           ymlConfig.UString("apiKey", ymlConfig.UString("apikey", os.Getenv("WTF_PAGERDUTY_API_KEY"))),
		escalationFilter: ymlConfig.UList("escalationFilter"),
		myName:           ymlConfig.UString("myName"),
		scheduleIDs:      ymlConfig.UList("scheduleIDs", []any{}),
		showIncidents:    ymlConfig.UBool("showIncidents", true),
		showOnCallEnd:    ymlConfig.UBool("showOnCallEnd", false),
		showSchedules:    ymlConfig.UBool("showSchedules", true),
		teamIDs:          ymlConfig.UList("teamIDs", []any{}),
		userIDs:          ymlConfig.UList("userIDs", []any{}),
	}

	cfg.ModuleSecret(name, globalConfig, &settings.apiKey).Load()

	return &settings
}
