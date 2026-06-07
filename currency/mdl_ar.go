//go:build (lang_ar || lang_all) && (country_all || country_eastern_europe || country_europe || country_md || currency_all || currency_mdl)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	MDL.RegisterName(xlanguage.Arabic, "ليو مولدوفي")
}
