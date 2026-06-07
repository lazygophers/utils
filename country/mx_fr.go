//go:build (lang_fr || lang_all) && (country_all || country_americas || country_central_america || country_mx)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMexico.RegisterName(xlanguage.French, "Mexique")
	dataMexico.RegisterOfficialName(xlanguage.French, "États-Unis mexicains")
	dataMexico.RegisterCapital(xlanguage.French, "Mexico")
}
