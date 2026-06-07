//go:build (lang_ja || lang_all) && (country_all || country_europe || country_je || country_northern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataJersey.RegisterName(xlanguage.Japanese, "ジャージー")
	dataJersey.RegisterOfficialName(xlanguage.Japanese, "ジャージー")
	dataJersey.RegisterCapital(xlanguage.Japanese, "セント・ヘリア")
}
