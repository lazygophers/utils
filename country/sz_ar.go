//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataEswatini.RegisterName(xlanguage.Arabic, "إسواتيني")
	dataEswatini.RegisterOfficialName(xlanguage.Arabic, "مملكة إسواتيني")
	dataEswatini.RegisterCapital(xlanguage.Arabic, "مباباني")
}
