//go:build (lang_fr || lang_all) && (country_africa || country_all || country_ma || country_northern_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMorocco.RegisterName(xlanguage.French, "Maroc")
	dataMorocco.RegisterOfficialName(xlanguage.French, "Royaume du Maroc")
	dataMorocco.RegisterCapital(xlanguage.French, "Rabat")
}
