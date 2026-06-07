//go:build (lang_ar || lang_all) && (country_all || country_melanesia || country_oceania || country_vu)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataVanuatu.RegisterName(xlanguage.Arabic, "فانواتو")
	dataVanuatu.RegisterOfficialName(xlanguage.Arabic, "جمهورية فانواتو")
	dataVanuatu.RegisterCapital(xlanguage.Arabic, "بورت فيلا")
}
