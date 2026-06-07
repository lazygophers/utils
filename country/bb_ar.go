//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBarbados.RegisterName(xlanguage.Arabic, "باربادوس")
	dataBarbados.RegisterOfficialName(xlanguage.Arabic, "باربادوس")
	dataBarbados.RegisterCapital(xlanguage.Arabic, "بريدجتاون")
}
