//go:build (lang_ja || lang_all) && (country_africa || country_all || country_tg || country_western_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTogo.RegisterName(xlanguage.Japanese, "トーゴ")
	dataTogo.RegisterOfficialName(xlanguage.Japanese, "トーゴ共和国")
	dataTogo.RegisterCapital(xlanguage.Japanese, "ロメ")
}
