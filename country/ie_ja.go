//go:build (lang_ja || lang_all) && (country_all || country_europe || country_ie || country_northern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIreland.RegisterName(xlanguage.Japanese, "アイルランド")
	dataIreland.RegisterOfficialName(xlanguage.Japanese, "アイルランド")
	dataIreland.RegisterCapital(xlanguage.Japanese, "ダブリン")
}
