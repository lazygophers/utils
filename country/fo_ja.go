//go:build (lang_ja || lang_all) && (country_all || country_europe || country_fo || country_northern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataFaroeIslands.RegisterName(xlanguage.Japanese, "フェロー諸島")
	dataFaroeIslands.RegisterOfficialName(xlanguage.Japanese, "フェロー諸島")
	dataFaroeIslands.RegisterCapital(xlanguage.Japanese, "トースハウン")
}
