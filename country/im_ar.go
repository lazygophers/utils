//go:build (lang_ar || lang_all) && (country_all || country_europe || country_im || country_northern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIsleOfMan.RegisterName(xlanguage.Arabic, "جزيرة مان")
	dataIsleOfMan.RegisterOfficialName(xlanguage.Arabic, "جزيرة مان")
	dataIsleOfMan.RegisterCapital(xlanguage.Arabic, "دوغلاس")
}
