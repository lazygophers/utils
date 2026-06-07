//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataRussia.RegisterName(xlanguage.Arabic, "روسيا")
	dataRussia.RegisterOfficialName(xlanguage.Arabic, "الاتحاد الروسي")
	dataRussia.RegisterCapital(xlanguage.Arabic, "موسكو")
}
