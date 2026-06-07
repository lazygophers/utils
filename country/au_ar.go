//go:build (lang_ar || lang_all) && (country_all || country_au || country_australia_and_new_zealand || country_oceania)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAustralia.RegisterName(xlanguage.Arabic, "أستراليا")
	dataAustralia.RegisterOfficialName(xlanguage.Arabic, "كومنولث أستراليا")
	dataAustralia.RegisterCapital(xlanguage.Arabic, "كانبرا")
}
