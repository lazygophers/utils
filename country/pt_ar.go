//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPortugal.RegisterName(xlanguage.Arabic, "البرتغال")
	dataPortugal.RegisterOfficialName(xlanguage.Arabic, "الجمهورية البرتغالية")
	dataPortugal.RegisterCapital(xlanguage.Arabic, "لشبونة")
}
