//go:build (lang_ar || lang_all) && (country_all || country_cy || country_europe || country_western_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCyprus.RegisterName(xlanguage.Arabic, "قبرص")
	dataCyprus.RegisterOfficialName(xlanguage.Arabic, "جمهورية قبرص")
	dataCyprus.RegisterCapital(xlanguage.Arabic, "نيقوسيا")
}
