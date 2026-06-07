//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNigeria.RegisterName(xlanguage.Arabic, "نيجيريا")
	dataNigeria.RegisterOfficialName(xlanguage.Arabic, "جمهورية نيجيريا الاتحادية")
	dataNigeria.RegisterCapital(xlanguage.Arabic, "أبوجا")
}
