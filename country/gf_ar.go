//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataFrenchGuiana.RegisterName(xlanguage.Arabic, "غويانا الفرنسية")
	dataFrenchGuiana.RegisterOfficialName(xlanguage.Arabic, "غويانا الفرنسية")
	dataFrenchGuiana.RegisterCapital(xlanguage.Arabic, "كايين")
}
