//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMalaysia.RegisterName(xlanguage.Arabic, "ماليزيا")
	dataMalaysia.RegisterOfficialName(xlanguage.Arabic, "ماليزيا")
	dataMalaysia.RegisterCapital(xlanguage.Arabic, "كوالالمبور")
}
