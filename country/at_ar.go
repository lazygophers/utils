//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAustria.RegisterName(xlanguage.Arabic, "النمسا")
	dataAustria.RegisterOfficialName(xlanguage.Arabic, "جمهورية النمسا")
	dataAustria.RegisterCapital(xlanguage.Arabic, "فيينا")
}
