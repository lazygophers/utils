//go:build (lang_fr || lang_all) && (country_africa || country_all || country_eh || country_northern_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataWesternSahara.RegisterName(xlanguage.French, "Sahara occidental")
	dataWesternSahara.RegisterOfficialName(xlanguage.French, "République arabe sahraouie démocratique")
	dataWesternSahara.RegisterCapital(xlanguage.French, "El Aaiún")
}
