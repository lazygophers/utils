//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSanMarino.RegisterName(xlanguage.Arabic, "سان مارينو")
	dataSanMarino.RegisterOfficialName(xlanguage.Arabic, "جمهورية سان مارينو")
	dataSanMarino.RegisterCapital(xlanguage.Arabic, "مدينة سان مارينو")
}
