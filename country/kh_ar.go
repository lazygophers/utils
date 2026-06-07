//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCambodia.RegisterName(xlanguage.Arabic, "كمبوديا")
	dataCambodia.RegisterOfficialName(xlanguage.Arabic, "مملكة كمبوديا")
	dataCambodia.RegisterCapital(xlanguage.Arabic, "بنوم بنه")
}
