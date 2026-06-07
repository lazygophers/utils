//go:build country_all || country_americas || country_central_america || country_pa

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPanama.RegisterName(xlanguage.English, "Panama")
	dataPanama.RegisterOfficialName(xlanguage.English, "Republic of Panama")
	dataPanama.RegisterCapital(xlanguage.English, "Panama City")
}
