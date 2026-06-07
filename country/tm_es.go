//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTurkmenistan.RegisterName(xlanguage.Spanish, "Turkmenistán")
	dataTurkmenistan.RegisterOfficialName(xlanguage.Spanish, "Turkmenistán")
	dataTurkmenistan.RegisterCapital(xlanguage.Spanish, "Asjabad")
}
