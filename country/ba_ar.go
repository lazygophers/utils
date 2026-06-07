//go:build (lang_ar || lang_all) && (country_all || country_ba || country_europe || country_southern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBosniaAndHerzegovina.RegisterName(xlanguage.Arabic, "البوسنة والهرسك")
	dataBosniaAndHerzegovina.RegisterOfficialName(xlanguage.Arabic, "البوسنة والهرسك")
	dataBosniaAndHerzegovina.RegisterCapital(xlanguage.Arabic, "سراييفو")
}
