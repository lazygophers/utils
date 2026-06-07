//go:build (lang_ar || lang_all) && (country_all || country_americas || country_gl || country_northern_america)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGreenland.RegisterName(xlanguage.Arabic, "غرينلاند")
	dataGreenland.RegisterOfficialName(xlanguage.Arabic, "غرينلاند")
	dataGreenland.RegisterCapital(xlanguage.Arabic, "نوك")
}
