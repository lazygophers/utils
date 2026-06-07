//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataChina.RegisterName(xlanguage.Arabic, "الصين")
	dataChina.RegisterOfficialName(xlanguage.Arabic, "جمهورية الصين الشعبية")
	dataChina.RegisterCapital(xlanguage.Arabic, "بكين")
}
