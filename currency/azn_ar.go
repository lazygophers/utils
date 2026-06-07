//go:build (lang_ar || lang_all) && (country_all || country_asia || country_az || country_western_asia || currency_all || currency_azn)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Azn.RegisterName(xlanguage.Arabic, "مانات أذربيجاني")
}
