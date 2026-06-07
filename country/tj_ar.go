//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTajikistan.RegisterName(xlanguage.Arabic, "طاجيكستان")
	dataTajikistan.RegisterOfficialName(xlanguage.Arabic, "جمهورية طاجيكستان")
	dataTajikistan.RegisterCapital(xlanguage.Arabic, "دوشنبه")
}
