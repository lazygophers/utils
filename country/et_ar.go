//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataEthiopia.RegisterName(xlanguage.Arabic, "إثيوبيا")
	dataEthiopia.RegisterOfficialName(xlanguage.Arabic, "جمهورية إثيوبيا الديمقراطية الاتحادية")
	dataEthiopia.RegisterCapital(xlanguage.Arabic, "أديس أبابا")
}
