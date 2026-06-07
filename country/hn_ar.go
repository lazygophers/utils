//go:build (lang_ar || lang_all) && (country_all || country_americas || country_central_america || country_hn)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataHonduras.RegisterName(xlanguage.Arabic, "هندوراس")
	dataHonduras.RegisterOfficialName(xlanguage.Arabic, "جمهورية هندوراس")
	dataHonduras.RegisterCapital(xlanguage.Arabic, "تيغوسيغالبا")
}
