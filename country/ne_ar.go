//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNiger.RegisterName(xlanguage.Arabic, "النيجر")
	dataNiger.RegisterOfficialName(xlanguage.Arabic, "جمهورية النيجر")
	dataNiger.RegisterCapital(xlanguage.Arabic, "نيامي")
}
