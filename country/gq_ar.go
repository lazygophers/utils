//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataEquatorialGuinea.RegisterName(xlanguage.Arabic, "غينيا الاستوائية")
	dataEquatorialGuinea.RegisterOfficialName(xlanguage.Arabic, "جمهورية غينيا الاستوائية")
	dataEquatorialGuinea.RegisterCapital(xlanguage.Arabic, "مالابو")
}
