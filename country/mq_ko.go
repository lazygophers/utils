//go:build (lang_ko || lang_all) && (country_all || country_americas || country_caribbean || country_mq)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMartinique.RegisterName(xlanguage.Korean, "마르티니크")
	dataMartinique.RegisterOfficialName(xlanguage.Korean, "마르티니크")
	dataMartinique.RegisterCapital(xlanguage.Korean, "포르드프랑스")
}
