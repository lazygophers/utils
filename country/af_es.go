//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAfghanistan.RegisterName(xlanguage.Spanish, "Afganistán")
	dataAfghanistan.RegisterOfficialName(xlanguage.Spanish, "Emirato Islámico de Afganistán")
	dataAfghanistan.RegisterCapital(xlanguage.Spanish, "Kabul")
}
