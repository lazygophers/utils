//go:build (lang_ar || lang_all) && (country_africa || country_all || country_eastern_africa || country_re)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataReunion.RegisterName(xlanguage.Arabic, "لا ريونيون")
	dataReunion.RegisterOfficialName(xlanguage.Arabic, "لا ريونيون")
	dataReunion.RegisterCapital(xlanguage.Arabic, "سان دوني")
}
