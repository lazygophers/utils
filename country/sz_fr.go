//go:build (lang_fr || lang_all) && (country_africa || country_all || country_southern_africa || country_sz)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataEswatini.RegisterName(xlanguage.French, "Eswatini")
	dataEswatini.RegisterOfficialName(xlanguage.French, "Royaume d'Eswatini")
	dataEswatini.RegisterCapital(xlanguage.French, "Mbabane")
}
