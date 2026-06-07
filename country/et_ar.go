//go:build (lang_ar || lang_all) && (country_africa || country_all || country_eastern_africa || country_et)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataEthiopia.RegisterName(xlanguage.Arabic, "إثيوبيا")
	dataEthiopia.RegisterOfficialName(xlanguage.Arabic, "جمهورية إثيوبيا الديمقراطية الاتحادية")
	dataEthiopia.RegisterCapital(xlanguage.Arabic, "أديس أبابا")
}
