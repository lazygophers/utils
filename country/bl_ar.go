//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSaintBarthelemy.RegisterName(xlanguage.Arabic, "سان بارتيلمي")
	dataSaintBarthelemy.RegisterOfficialName(xlanguage.Arabic, "جماعة سان بارتيلمي")
	dataSaintBarthelemy.RegisterCapital(xlanguage.Arabic, "غوستافيا")
}
