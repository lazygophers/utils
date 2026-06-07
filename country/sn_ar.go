//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSenegal.RegisterName(xlanguage.Arabic, "السنغال")
	dataSenegal.RegisterOfficialName(xlanguage.Arabic, "جمهورية السنغال")
	dataSenegal.RegisterCapital(xlanguage.Arabic, "داكار")
}
