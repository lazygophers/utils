//go:build (lang_ar || lang_all) && (country_africa || country_all || country_southern_africa || country_za)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSouthAfrica.RegisterName(xlanguage.Arabic, "جنوب أفريقيا")
	dataSouthAfrica.RegisterOfficialName(xlanguage.Arabic, "جمهورية جنوب أفريقيا")
	dataSouthAfrica.RegisterCapital(xlanguage.Arabic, "بريتوريا")
}
