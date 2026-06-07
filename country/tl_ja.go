//go:build (lang_ja || lang_all) && (country_all || country_asia || country_south_eastern_asia || country_tl)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTimorLeste.RegisterName(xlanguage.Japanese, "東ティモール")
	dataTimorLeste.RegisterOfficialName(xlanguage.Japanese, "東ティモール民主共和国")
	dataTimorLeste.RegisterCapital(xlanguage.Japanese, "ディリ")
}
