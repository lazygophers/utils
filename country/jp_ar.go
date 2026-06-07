//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataJapan.RegisterName(xlanguage.Arabic, "اليابان")
	dataJapan.RegisterOfficialName(xlanguage.Arabic, "اليابان")
	dataJapan.RegisterCapital(xlanguage.Arabic, "طوكيو")
}
