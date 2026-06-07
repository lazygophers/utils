//go:build (lang_fr || lang_all) && (country_africa || country_all || country_eastern_africa || country_ss)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSouthSudan.RegisterName(xlanguage.French, "Soudan du Sud")
	dataSouthSudan.RegisterOfficialName(xlanguage.French, "République du Soudan du Sud")
	dataSouthSudan.RegisterCapital(xlanguage.French, "Djouba")
}
