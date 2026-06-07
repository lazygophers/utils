//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSaudiArabia.RegisterName(xlanguage.French, "Arabie saoudite")
	dataSaudiArabia.RegisterOfficialName(xlanguage.French, "Royaume d'Arabie saoudite")
	dataSaudiArabia.RegisterCapital(xlanguage.French, "Riyad")
}
