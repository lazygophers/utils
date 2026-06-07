//go:build (lang_ja || lang_all) && (country_all || country_americas || country_caribbean || country_gp)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGuadeloupe.RegisterName(xlanguage.Japanese, "グアドループ")
	dataGuadeloupe.RegisterOfficialName(xlanguage.Japanese, "グアドループ")
	dataGuadeloupe.RegisterCapital(xlanguage.Japanese, "バステール")
}
