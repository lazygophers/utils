//go:build (lang_es || lang_all) && (country_africa || country_all || country_eh || country_northern_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataWesternSahara.RegisterName(xlanguage.Spanish, "Sahara Occidental")
	dataWesternSahara.RegisterOfficialName(xlanguage.Spanish, "República Árabe Saharaui Democrática")
	dataWesternSahara.RegisterCapital(xlanguage.Spanish, "El Aaiún")
}
