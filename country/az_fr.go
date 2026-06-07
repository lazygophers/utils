//go:build (lang_fr || lang_all) && (country_all || country_asia || country_az || country_western_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAzerbaijan.RegisterName(xlanguage.French, "Azerbaïdjan")
	dataAzerbaijan.RegisterOfficialName(xlanguage.French, "République d'Azerbaïdjan")
	dataAzerbaijan.RegisterCapital(xlanguage.French, "Bakou")
}
