//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCostaRica.RegisterName(xlanguage.Arabic, "كوستاريكا")
	dataCostaRica.RegisterOfficialName(xlanguage.Arabic, "جمهورية كوستاريكا")
	dataCostaRica.RegisterCapital(xlanguage.Arabic, "سان خوسيه")
}
