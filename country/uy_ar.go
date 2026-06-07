//go:build (lang_ar || lang_all) && (country_all || country_americas || country_south_america || country_uy)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUruguay.RegisterName(xlanguage.Arabic, "الأوروغواي")
	dataUruguay.RegisterOfficialName(xlanguage.Arabic, "جمهورية الأوروغواي الشرقية")
	dataUruguay.RegisterCapital(xlanguage.Arabic, "مونتيفيديو")
}
