//go:build (lang_es || lang_all) && (country_all || country_antarctic || country_hm)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataHeardAndMcDonaldIslands.RegisterName(xlanguage.Spanish, "Islas Heard y McDonald")
	dataHeardAndMcDonaldIslands.RegisterOfficialName(xlanguage.Spanish, "Territorio de las Islas Heard y McDonald")
}
