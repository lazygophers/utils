//go:build (lang_fr || lang_all) && (country_all || country_americas || country_central_america || country_sv)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataElSalvador.RegisterName(xlanguage.French, "Salvador")
	dataElSalvador.RegisterOfficialName(xlanguage.French, "République du Salvador")
	dataElSalvador.RegisterCapital(xlanguage.French, "San Salvador")
}
