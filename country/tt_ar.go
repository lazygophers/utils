//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTrinidadAndTobago.RegisterName(xlanguage.Arabic, "ترينيداد وتوباغو")
	dataTrinidadAndTobago.RegisterOfficialName(xlanguage.Arabic, "جمهورية ترينيداد وتوباغو")
	dataTrinidadAndTobago.RegisterCapital(xlanguage.Arabic, "بورت أوف سبين")
}
