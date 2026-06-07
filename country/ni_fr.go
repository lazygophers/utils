//go:build (lang_fr || lang_all) && (country_all || country_americas || country_central_america || country_ni)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNicaragua.RegisterName(xlanguage.French, "Nicaragua")
	dataNicaragua.RegisterOfficialName(xlanguage.French, "République du Nicaragua")
	dataNicaragua.RegisterCapital(xlanguage.French, "Managua")
}
