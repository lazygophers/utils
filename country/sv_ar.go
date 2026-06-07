//go:build (lang_ar || lang_all) && (country_all || country_americas || country_central_america || country_sv)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataElSalvador.RegisterName(xlanguage.Arabic, "السلفادور")
	dataElSalvador.RegisterOfficialName(xlanguage.Arabic, "جمهورية السلفادور")
	dataElSalvador.RegisterCapital(xlanguage.Arabic, "سان سلفادور")
}
