//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIran.RegisterName(xlanguage.Arabic, "إيران")
	dataIran.RegisterOfficialName(xlanguage.Arabic, "جمهورية إيران الإسلامية")
	dataIran.RegisterCapital(xlanguage.Arabic, "طهران")
}
