//go:build (lang_fr || lang_all) && (country_africa || country_all || country_sh || country_western_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSaintHelena.RegisterName(xlanguage.French, "Sainte-Hélène, Ascension et Tristan da Cunha")
	dataSaintHelena.RegisterOfficialName(xlanguage.French, "Sainte-Hélène, Ascension et Tristan da Cunha")
	dataSaintHelena.RegisterCapital(xlanguage.French, "Jamestown")
}
