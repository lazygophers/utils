//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGeorgia.RegisterName(xlanguage.Arabic, "جورجيا")
	dataGeorgia.RegisterOfficialName(xlanguage.Arabic, "جورجيا")
	dataGeorgia.RegisterCapital(xlanguage.Arabic, "تبليسي")
}
