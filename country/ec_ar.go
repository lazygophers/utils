//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataEcuador.RegisterName(xlanguage.Arabic, "الإكوادور")
	dataEcuador.RegisterOfficialName(xlanguage.Arabic, "جمهورية الإكوادور")
	dataEcuador.RegisterCapital(xlanguage.Arabic, "كيتو")
}
