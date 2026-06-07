//go:build (lang_ar || lang_all) && (country_all || country_oceania || country_polynesia || country_to || currency_all || currency_top)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Top.RegisterName(xlanguage.Arabic, "بانغا تونغية")
}
