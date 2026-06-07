//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUzbekistan.RegisterName(xlanguage.Arabic, "أوزبكستان")
	dataUzbekistan.RegisterOfficialName(xlanguage.Arabic, "جمهورية أوزبكستان")
	dataUzbekistan.RegisterCapital(xlanguage.Arabic, "طشقند")
}
