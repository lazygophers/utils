//go:build country_africa || country_all || country_sh || country_western_africa

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSaintHelena.RegisterName(xlanguage.English, "Saint Helena, Ascension and Tristan da Cunha")
	dataSaintHelena.RegisterOfficialName(xlanguage.English, "Saint Helena, Ascension and Tristan da Cunha")
	dataSaintHelena.RegisterCapital(xlanguage.English, "Jamestown")
}
