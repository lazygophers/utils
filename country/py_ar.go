//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataParaguay.RegisterName(xlanguage.Arabic, "باراغواي")
	dataParaguay.RegisterOfficialName(xlanguage.Arabic, "جمهورية باراغواي")
	dataParaguay.RegisterCapital(xlanguage.Arabic, "أسونسيون")
}
