//go:build (lang_ja || lang_all) && (country_all || country_eastern_europe || country_europe || country_ro)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataRomania.RegisterName(xlanguage.Japanese, "ルーマニア")
	dataRomania.RegisterOfficialName(xlanguage.Japanese, "ルーマニア")
	dataRomania.RegisterCapital(xlanguage.Japanese, "ブカレスト")
}
