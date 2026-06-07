//go:build (lang_ja || lang_all) && (country_africa || country_all || country_gm || country_western_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGambia.RegisterName(xlanguage.Japanese, "ガンビア")
	dataGambia.RegisterOfficialName(xlanguage.Japanese, "ガンビア共和国")
	dataGambia.RegisterCapital(xlanguage.Japanese, "バンジュール")
}
