//go:build (lang_ja || lang_all) && (country_africa || country_all || country_ma || country_northern_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMorocco.RegisterName(xlanguage.Japanese, "モロッコ")
	dataMorocco.RegisterOfficialName(xlanguage.Japanese, "モロッコ王国")
	dataMorocco.RegisterCapital(xlanguage.Japanese, "ラバト")
}
