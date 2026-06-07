//go:build (lang_fr || lang_all) && (country_all || country_asia || country_la || country_south_eastern_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLaos.RegisterName(xlanguage.French, "Laos")
	dataLaos.RegisterOfficialName(xlanguage.French, "République démocratique populaire lao")
	dataLaos.RegisterCapital(xlanguage.French, "Vientiane")
}
