//go:build country_all || country_americas || country_caribbean || country_vc

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSaintVincentAndGrenadines.RegisterName(xlanguage.English, "Saint Vincent and the Grenadines")
	dataSaintVincentAndGrenadines.RegisterOfficialName(xlanguage.English, "Saint Vincent and the Grenadines")
	dataSaintVincentAndGrenadines.RegisterCapital(xlanguage.English, "Kingstown")
}
