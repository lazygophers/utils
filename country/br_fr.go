//go:build (lang_fr || lang_all) && (country_all || country_americas || country_br || country_south_america)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBrazil.RegisterName(xlanguage.French, "Brésil")
	dataBrazil.RegisterOfficialName(xlanguage.French, "République fédérative du Brésil")
	dataBrazil.RegisterCapital(xlanguage.French, "Brasilia")
}
