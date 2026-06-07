//go:build (lang_fr || lang_all) && (country_all || country_asia || country_western_asia || country_ye)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataYemen.RegisterName(xlanguage.French, "Yémen")
	dataYemen.RegisterOfficialName(xlanguage.French, "République du Yémen")
	dataYemen.RegisterCapital(xlanguage.French, "Sanaa")
}
