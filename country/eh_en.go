//go:build country_africa || country_all || country_eh || country_northern_africa

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataWesternSahara.RegisterName(xlanguage.English, "Western Sahara")
	dataWesternSahara.RegisterOfficialName(xlanguage.English, "Sahrawi Arab Democratic Republic")
	dataWesternSahara.RegisterCapital(xlanguage.English, "El Aaiun")
}
