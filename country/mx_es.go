//go:build country_all || country_americas || country_central_america || country_mx

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMexico.RegisterName(xlanguage.Spanish, "México")
	dataMexico.RegisterOfficialName(xlanguage.Spanish, "Estados Unidos Mexicanos")
	dataMexico.RegisterCapital(xlanguage.Spanish, "Ciudad de México")
}
