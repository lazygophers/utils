//go:build (lang_fr || lang_all) && (country_all || country_oceania || country_polynesia || country_ws)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSamoa.RegisterName(xlanguage.French, "Samoa")
	dataSamoa.RegisterOfficialName(xlanguage.French, "État indépendant des Samoa")
	dataSamoa.RegisterCapital(xlanguage.French, "Apia")
}
