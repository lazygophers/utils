//go:build (lang_ja || lang_all) && (country_all || country_ax || country_europe || country_northern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAlandIslands.RegisterName(xlanguage.Japanese, "オーランド諸島")
	dataAlandIslands.RegisterOfficialName(xlanguage.Japanese, "オーランド諸島")
	dataAlandIslands.RegisterCapital(xlanguage.Japanese, "マリエハムン")
}
