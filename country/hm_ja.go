//go:build (lang_ja || lang_all) && (country_all || country_antarctic || country_hm)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataHeardAndMcDonaldIslands.RegisterName(xlanguage.Japanese, "ハード島とマクドナルド諸島")
	dataHeardAndMcDonaldIslands.RegisterOfficialName(xlanguage.Japanese, "ハード島とマクドナルド諸島")
}
