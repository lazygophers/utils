//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTogo.RegisterName(xlanguage.Arabic, "توغو")
	dataTogo.RegisterOfficialName(xlanguage.Arabic, "الجمهورية التوغولية")
	dataTogo.RegisterCapital(xlanguage.Arabic, "لومي")
}
