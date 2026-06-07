//go:build (lang_fr || lang_all) && (country_all || country_americas || country_py || country_south_america)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataParaguay.RegisterName(xlanguage.French, "Paraguay")
	dataParaguay.RegisterOfficialName(xlanguage.French, "République du Paraguay")
	dataParaguay.RegisterCapital(xlanguage.French, "Asuncion")
}
