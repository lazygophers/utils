//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGreenland.RegisterName(xlanguage.Arabic, "غرينلاند")
	dataGreenland.RegisterOfficialName(xlanguage.Arabic, "غرينلاند")
	dataGreenland.RegisterCapital(xlanguage.Arabic, "نوك")
}
