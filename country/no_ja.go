//go:build (lang_ja || lang_all) && (country_all || country_europe || country_no || country_northern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNorway.RegisterName(xlanguage.Japanese, "ノルウェー")
	dataNorway.RegisterOfficialName(xlanguage.Japanese, "ノルウェー王国")
	dataNorway.RegisterCapital(xlanguage.Japanese, "オスロ")
}
