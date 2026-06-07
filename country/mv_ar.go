//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMaldives.RegisterName(xlanguage.Arabic, "المالديف")
	dataMaldives.RegisterOfficialName(xlanguage.Arabic, "جمهورية المالديف")
	dataMaldives.RegisterCapital(xlanguage.Arabic, "ماليه")
}
