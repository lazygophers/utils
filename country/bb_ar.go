//go:build (lang_ar || lang_all) && (country_all || country_americas || country_bb || country_caribbean)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBarbados.RegisterName(xlanguage.Arabic, "باربادوس")
	dataBarbados.RegisterOfficialName(xlanguage.Arabic, "باربادوس")
	dataBarbados.RegisterCapital(xlanguage.Arabic, "بريدجتاون")
}
