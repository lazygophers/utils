//go:build (lang_ja || lang_all) && (country_all || country_asia || country_la || country_south_eastern_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLaos.RegisterName(xlanguage.Japanese, "ラオス")
	dataLaos.RegisterOfficialName(xlanguage.Japanese, "ラオス人民民主共和国")
	dataLaos.RegisterCapital(xlanguage.Japanese, "ヴィエンチャン")
}
