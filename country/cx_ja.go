//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataChristmasIsland.RegisterName(xlanguage.Japanese, "クリスマス島")
	dataChristmasIsland.RegisterOfficialName(xlanguage.Japanese, "クリスマス島")
	dataChristmasIsland.RegisterCapital(xlanguage.Japanese, "フライングフィッシュ・コーブ")
}
