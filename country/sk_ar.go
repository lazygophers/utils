//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSlovakia.RegisterName(xlanguage.Arabic, "سلوفاكيا")
	dataSlovakia.RegisterOfficialName(xlanguage.Arabic, "الجمهورية السلوفاكية")
	dataSlovakia.RegisterCapital(xlanguage.Arabic, "براتيسلافا")
}
