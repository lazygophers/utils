//go:build (lang_ar || lang_all) && (country_all || country_asia || country_om || country_western_asia || currency_all || currency_omr)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	OMR.RegisterName(xlanguage.Arabic, "ريال عماني")
}
