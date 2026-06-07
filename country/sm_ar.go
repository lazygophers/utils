//go:build (lang_ar || lang_all) && (country_all || country_europe || country_sm || country_southern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSanMarino.RegisterName(xlanguage.Arabic, "سان مارينو")
	dataSanMarino.RegisterOfficialName(xlanguage.Arabic, "جمهورية سان مارينو")
	dataSanMarino.RegisterCapital(xlanguage.Arabic, "مدينة سان مارينو")
}
