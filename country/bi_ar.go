//go:build (lang_ar || lang_all) && (country_africa || country_all || country_bi || country_eastern_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBurundi.RegisterName(xlanguage.Arabic, "بوروندي")
	dataBurundi.RegisterOfficialName(xlanguage.Arabic, "جمهورية بوروندي")
	dataBurundi.RegisterCapital(xlanguage.Arabic, "غيتيغا")
}
