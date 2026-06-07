//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGrenada.RegisterName(xlanguage.Arabic, "غرينادا")
	dataGrenada.RegisterOfficialName(xlanguage.Arabic, "غرينادا")
	dataGrenada.RegisterCapital(xlanguage.Arabic, "سانت جورج")
}
