//go:build (lang_fr || lang_all) && (country_africa || country_all || country_northern_africa || country_tn)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTunisia.RegisterName(xlanguage.French, "Tunisie")
	dataTunisia.RegisterOfficialName(xlanguage.French, "République tunisienne")
	dataTunisia.RegisterCapital(xlanguage.French, "Tunis")
}
