//go:build (lang_es || lang_all) && (country_africa || country_all || country_northern_africa || country_tn)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTunisia.RegisterName(xlanguage.Spanish, "Túnez")
	dataTunisia.RegisterOfficialName(xlanguage.Spanish, "República Tunecina")
	dataTunisia.RegisterCapital(xlanguage.Spanish, "Túnez")
}
