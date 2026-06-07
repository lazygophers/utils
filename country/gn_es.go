//go:build (lang_es || lang_all) && (country_africa || country_all || country_gn || country_western_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGuinea.RegisterName(xlanguage.Spanish, "Guinea")
	dataGuinea.RegisterOfficialName(xlanguage.Spanish, "República de Guinea")
	dataGuinea.RegisterCapital(xlanguage.Spanish, "Conakri")
}
