//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTurkmenistan.RegisterName(xlanguage.French, "Turkménistan")
	dataTurkmenistan.RegisterOfficialName(xlanguage.French, "Turkménistan")
	dataTurkmenistan.RegisterCapital(xlanguage.French, "Achgabat")
}
