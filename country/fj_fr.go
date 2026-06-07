//go:build (lang_fr || lang_all) && (country_all || country_fj || country_melanesia || country_oceania)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataFiji.RegisterName(xlanguage.French, "Fidji")
	dataFiji.RegisterOfficialName(xlanguage.French, "République des Fidji")
	dataFiji.RegisterCapital(xlanguage.French, "Suva")
}
