//go:build country_africa || country_all || country_ci || country_western_africa

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIvoryCoast.RegisterName(xlanguage.English, "Ivory Coast")
	dataIvoryCoast.RegisterOfficialName(xlanguage.English, "Republic of Cote d'Ivoire")
	dataIvoryCoast.RegisterCapital(xlanguage.English, "Yamoussoukro")
}
