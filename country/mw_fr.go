//go:build (lang_fr || lang_all) && (country_africa || country_all || country_eastern_africa || country_mw)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMalawi.RegisterName(xlanguage.French, "Malawi")
	dataMalawi.RegisterOfficialName(xlanguage.French, "République du Malawi")
	dataMalawi.RegisterCapital(xlanguage.French, "Lilongwe")
}
