//go:build country_africa || country_all || country_ma || country_northern_africa

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMorocco.RegisterName(xlanguage.English, "Morocco")
	dataMorocco.RegisterOfficialName(xlanguage.English, "Kingdom of Morocco")
	dataMorocco.RegisterCapital(xlanguage.English, "Rabat")
}
