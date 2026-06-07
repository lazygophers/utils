//go:build (lang_ja || lang_all) && (country_all || country_europe || country_gr || country_southern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGreece.RegisterName(xlanguage.Japanese, "ギリシャ")
	dataGreece.RegisterOfficialName(xlanguage.Japanese, "ギリシャ共和国")
	dataGreece.RegisterCapital(xlanguage.Japanese, "アテネ")
}
