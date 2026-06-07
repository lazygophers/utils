//go:build (lang_fr || lang_all) && (country_africa || country_all || country_lr || country_western_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLiberia.RegisterName(xlanguage.French, "Liberia")
	dataLiberia.RegisterOfficialName(xlanguage.French, "République du Liberia")
	dataLiberia.RegisterCapital(xlanguage.French, "Monrovia")
}
