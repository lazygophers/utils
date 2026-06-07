//go:build (lang_ar || lang_all) && (country_all || country_asia || country_central_asia || country_tj || currency_all || currency_tjs)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	TJS.RegisterName(xlanguage.Arabic, "سوموني طاجيكي")
}
