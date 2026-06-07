//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSuriname.RegisterName(xlanguage.Arabic, "سورينام")
	dataSuriname.RegisterOfficialName(xlanguage.Arabic, "جمهورية سورينام")
	dataSuriname.RegisterCapital(xlanguage.Arabic, "باراماريبو")
}
