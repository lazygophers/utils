//go:build (lang_es || lang_all) && (country_all || country_oceania || country_polynesia || country_ws)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSamoa.RegisterName(xlanguage.Spanish, "Samoa")
	dataSamoa.RegisterOfficialName(xlanguage.Spanish, "Estado Independiente de Samoa")
	dataSamoa.RegisterCapital(xlanguage.Spanish, "Apia")
}
