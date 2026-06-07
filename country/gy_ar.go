//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGuyana.RegisterName(xlanguage.Arabic, "غيانا")
	dataGuyana.RegisterOfficialName(xlanguage.Arabic, "جمهورية غيانا التعاونية")
	dataGuyana.RegisterCapital(xlanguage.Arabic, "جورجتاون")
}
