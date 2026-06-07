//go:build country_all || country_americas || country_caribbean || country_do

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataDominicanRepublic.RegisterName(xlanguage.English, "Dominican Republic")
	dataDominicanRepublic.RegisterOfficialName(xlanguage.English, "Dominican Republic")
	dataDominicanRepublic.RegisterCapital(xlanguage.English, "Santo Domingo")
}
