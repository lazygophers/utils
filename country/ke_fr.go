//go:build (lang_fr || lang_all) && (country_africa || country_all || country_eastern_africa || country_ke)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataKenya.RegisterName(xlanguage.French, "Kenya")
	dataKenya.RegisterOfficialName(xlanguage.French, "République du Kenya")
	dataKenya.RegisterCapital(xlanguage.French, "Nairobi")
}
