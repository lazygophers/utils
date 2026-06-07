//go:build (lang_ja || lang_all) && (country_all || country_ck || country_oceania || country_polynesia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCookIslands.RegisterName(xlanguage.Japanese, "クック諸島")
	dataCookIslands.RegisterOfficialName(xlanguage.Japanese, "クック諸島")
	dataCookIslands.RegisterCapital(xlanguage.Japanese, "アバルア")
}
