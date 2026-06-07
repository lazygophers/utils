//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLesotho.RegisterName(xlanguage.Arabic, "ليسوتو")
	dataLesotho.RegisterOfficialName(xlanguage.Arabic, "مملكة ليسوتو")
	dataLesotho.RegisterCapital(xlanguage.Arabic, "ماسيرو")
}
