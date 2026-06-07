//go:build (lang_ja || lang_all) && (country_all || country_americas || country_caribbean || country_mq)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMartinique.RegisterName(xlanguage.Japanese, "マルティニーク")
	dataMartinique.RegisterOfficialName(xlanguage.Japanese, "マルティニーク")
	dataMartinique.RegisterCapital(xlanguage.Japanese, "フォール＝ド＝フランス")
}
