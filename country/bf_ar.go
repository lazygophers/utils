//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBurkinaFaso.RegisterName(xlanguage.Arabic, "بوركينا فاسو")
	dataBurkinaFaso.RegisterOfficialName(xlanguage.Arabic, "بوركينا فاسو")
	dataBurkinaFaso.RegisterCapital(xlanguage.Arabic, "واغادوغو")
}
