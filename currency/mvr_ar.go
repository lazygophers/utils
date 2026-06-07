//go:build (lang_ar || lang_all) && (country_all || country_asia || country_mv || country_southern_asia || currency_all || currency_mvr)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Mvr.RegisterName(xlanguage.Arabic, "روفيه مالديفي")
}
