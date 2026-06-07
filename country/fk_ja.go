//go:build (lang_ja || lang_all) && (country_all || country_americas || country_fk || country_south_america)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataFalklandIslands.RegisterName(xlanguage.Japanese, "フォークランド諸島")
	dataFalklandIslands.RegisterOfficialName(xlanguage.Japanese, "フォークランド諸島")
	dataFalklandIslands.RegisterCapital(xlanguage.Japanese, "スタンレー")
}
