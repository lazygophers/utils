//go:build (lang_fr || lang_all) && (country_all || country_americas || country_central_america || country_hn)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataHonduras.RegisterName(xlanguage.French, "Honduras")
	dataHonduras.RegisterOfficialName(xlanguage.French, "République du Honduras")
	dataHonduras.RegisterCapital(xlanguage.French, "Tegucigalpa")
}
