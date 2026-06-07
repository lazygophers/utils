//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIndia.RegisterName(xlanguage.Arabic, "الهند")
	dataIndia.RegisterOfficialName(xlanguage.Arabic, "جمهورية الهند")
	dataIndia.RegisterCapital(xlanguage.Arabic, "نيودلهي")
}
