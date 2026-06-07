//go:build (lang_ja || lang_all) && (country_all || country_melanesia || country_oceania || country_sb)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSolomonIslands.RegisterName(xlanguage.Japanese, "ソロモン諸島")
	dataSolomonIslands.RegisterOfficialName(xlanguage.Japanese, "ソロモン諸島")
	dataSolomonIslands.RegisterCapital(xlanguage.Japanese, "ホニアラ")
}
