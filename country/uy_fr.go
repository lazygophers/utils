//go:build (lang_fr || lang_all) && (country_all || country_americas || country_south_america || country_uy)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUruguay.RegisterName(xlanguage.French, "Uruguay")
	dataUruguay.RegisterOfficialName(xlanguage.French, "République orientale de l'Uruguay")
	dataUruguay.RegisterCapital(xlanguage.French, "Montevideo")
}
