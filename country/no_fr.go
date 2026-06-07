//go:build (lang_fr || lang_all) && (country_all || country_europe || country_no || country_northern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNorway.RegisterName(xlanguage.French, "Norvège")
	dataNorway.RegisterOfficialName(xlanguage.French, "Royaume de Norvège")
	dataNorway.RegisterCapital(xlanguage.French, "Oslo")
}
