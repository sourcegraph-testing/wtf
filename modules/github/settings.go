package github

import (
	"os"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const (
	defaultFocusable = true
	defaultTitle     = "GitHub"
)

// Settings defines the configuration properties for this module
type Settings struct {
	*cfg.Common

	apiKey                 string        `help:"Your GitHub API token."`
	baseURL                string        `help:"Your GitHub Enterprise API URL." optional:"true"`
	customQueries          []customQuery `help:"Custom queries allow you to filter pull requests and issues however you like. Give the query a title and a filter. Filters can be copied directly from GitHub’s UI." optional:"true"`
	enableStatus           bool          `help:"Display pull request mergeability status (‘dirty’, ‘clean’, ‘unstable’, ‘blocked’)." optional:"true"`
	repositories           []string      `help:"A list of github repositories." values:"Example: wtfutil/wtf"`
	showMyPullRequests     bool          `help:"Show my pull requests section" optional:"true"`
	showOpenReviewRequests bool          `help:"Show open review requests section" optional:"true"`
	showStats              bool          `help:"Show repository stats section" optional:"true"`
	uploadURL              string        `help:"Your GitHub Enterprise upload URL (often the same as API URL)." optional:"true"`
	username               string        `help:"Your GitHub username. Used to figure out which review requests you’ve been added to."`
}

type customQuery struct {
	title   string `help:"Display title for this query"`
	filter  string `help:"Github query filter"`
	perPage int    `help:"Number of issues to show"`
}

// NewSettingsFromYAML creates a new settings instance from a YAML config block
func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	settings := Settings{
		Common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),

		apiKey:                 ymlConfig.UString("apiKey", ymlConfig.UString("apikey", os.Getenv("WTF_GITHUB_TOKEN"))),
		baseURL:                ymlConfig.UString("baseURL", os.Getenv("WTF_GITHUB_BASE_URL")),
		enableStatus:           ymlConfig.UBool("enableStatus", false),
		showMyPullRequests:     ymlConfig.UBool("showMyPullRequests", true),
		showOpenReviewRequests: ymlConfig.UBool("showOpenReviewRequests", true),
		showStats:              ymlConfig.UBool("showStats", true),
		uploadURL:              ymlConfig.UString("uploadURL", os.Getenv("WTF_GITHUB_UPLOAD_URL")),
		username:               ymlConfig.UString("username"),
	}
	settings.repositories = cfg.ParseAsMapOrList(ymlConfig, "repositories")
	settings.customQueries = parseCustomQueries(ymlConfig)

	cfg.ModuleSecret(name, globalConfig, &settings.apiKey).
		Service(settings.baseURL).Load()

	return &settings
}

/* -------------------- Unexported Functions -------------------- */

func parseCustomQueries(ymlConfig *config.Config) []customQuery {
	result := []customQuery{}
	if customQueries, err := ymlConfig.Map("customQueries"); err == nil {
		for _, query := range customQueries {
			c := customQuery{}
			for key, value := range query.(map[string]any) {
				switch key {
				case "title":
					c.title = value.(string)
				case "filter":
					c.filter = value.(string)
				case "perPage":
					c.perPage = value.(int)
				}
			}

			if c.title != "" && c.filter != "" {
				result = append(result, c)
			}
		}
	}
	return result
}
