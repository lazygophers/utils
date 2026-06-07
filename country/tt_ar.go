//go:build (lang_ar || lang_all) && (country_all || country_americas || country_caribbean || country_tt)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTrinidadAndTobago.RegisterName(xlanguage.Arabic, "ترينيداد وتوباغو")
	dataTrinidadAndTobago.RegisterOfficialName(xlanguage.Arabic, "جمهورية ترينيداد وتوباغو")
	dataTrinidadAndTobago.RegisterCapital(xlanguage.Arabic, "بورت أوف سبين")
}
