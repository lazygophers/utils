//go:build country_all || country_americas || country_caribbean || country_do

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataDominicanRepublic.RegisterName(xlanguage.Spanish, "República Dominicana")
	dataDominicanRepublic.RegisterOfficialName(xlanguage.Spanish, "República Dominicana")
	dataDominicanRepublic.RegisterCapital(xlanguage.Spanish, "Santo Domingo")
}
