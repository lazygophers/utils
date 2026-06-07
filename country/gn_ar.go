//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGuinea.RegisterName(xlanguage.Arabic, "غينيا")
	dataGuinea.RegisterOfficialName(xlanguage.Arabic, "جمهورية غينيا")
	dataGuinea.RegisterCapital(xlanguage.Arabic, "كوناكري")
}
