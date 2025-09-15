//go:build lang_en || lang_all

package validator

// Register English language configuration
func init() {
	RegisterLocaleConfig("en", &LocaleConfig{
		Language: "en",
		Region:   "US",
		Messages: map[string]string{
			"required": "{field} is required",
			"email":    "{field} must be a valid email address",
			"url":      "{field} must be a valid URL",
			"min":      "{field} minimum value is {param}",
			"max":      "{field} maximum value is {param}",
			"len":      "{field} length must be {param} characters",
		},
	})
}