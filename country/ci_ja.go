//go:build (lang_ja || lang_all) && (country_africa || country_all || country_ci || country_western_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIvoryCoast.RegisterName(xlanguage.Japanese, "コートジボワール")
	dataIvoryCoast.RegisterOfficialName(xlanguage.Japanese, "コートジボワール共和国")
	dataIvoryCoast.RegisterCapital(xlanguage.Japanese, "ヤムスクロ")
}
