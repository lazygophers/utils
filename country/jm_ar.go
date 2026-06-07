//go:build (lang_ar || lang_all) && (country_all || country_americas || country_caribbean || country_jm)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataJamaica.RegisterName(xlanguage.Arabic, "جامايكا")
	dataJamaica.RegisterOfficialName(xlanguage.Arabic, "جامايكا")
	dataJamaica.RegisterCapital(xlanguage.Arabic, "كينغستون")
}
