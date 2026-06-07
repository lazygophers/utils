//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataFrance.RegisterName(xlanguage.Arabic, "فرنسا")
	dataFrance.RegisterOfficialName(xlanguage.Arabic, "الجمهورية الفرنسية")
	dataFrance.RegisterCapital(xlanguage.Arabic, "باريس")
}
