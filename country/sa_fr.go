//go:build (lang_fr || lang_all) && (country_all || country_asia || country_sa || country_western_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSaudiArabia.RegisterName(xlanguage.French, "Arabie saoudite")
	dataSaudiArabia.RegisterOfficialName(xlanguage.French, "Royaume d'Arabie saoudite")
	dataSaudiArabia.RegisterCapital(xlanguage.French, "Riyad")
}
