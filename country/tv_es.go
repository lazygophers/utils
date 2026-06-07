//go:build (lang_es || lang_all) && (country_all || country_oceania || country_polynesia || country_tv)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTuvalu.RegisterName(xlanguage.Spanish, "Tuvalu")
	dataTuvalu.RegisterOfficialName(xlanguage.Spanish, "Tuvalu")
	dataTuvalu.RegisterCapital(xlanguage.Spanish, "Funafuti")
}
